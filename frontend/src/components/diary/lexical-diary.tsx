import React, { useEffect, useRef, useState } from "react"
import { LexicalComposer } from "@lexical/react/LexicalComposer"
import { ContentEditable } from "@lexical/react/LexicalContentEditable"
import { LexicalErrorBoundary } from "@lexical/react/LexicalErrorBoundary"
import { HistoryPlugin } from "@lexical/react/LexicalHistoryPlugin"
import { OnChangePlugin } from "@lexical/react/LexicalOnChangePlugin"
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin"
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext"
import { HeadingNode, $createHeadingNode } from "@lexical/rich-text"
import { $setBlocksType } from "@lexical/selection"
import {
    $applyNodeReplacement,
    $createParagraphNode,
    $getSelection,
    $insertNodes,
    $isRangeSelection,
    DecoratorNode,
    EditorConfig,
    FORMAT_TEXT_COMMAND,
    LexicalEditor,
    NodeKey,
    SerializedLexicalNode,
    Spread
} from "lexical"
import { postFormData } from "../../api-methods"
import { getSession } from "../../logic/session-manager"
import { resolveServerUrl } from "../../utility"
import {
    DIARY_RICH_TEXT_DOCUMENT_VERSION,
    DiaryRichTextBlock,
    DiaryRichTextDocument,
    DiaryRichTextHeadingLevel,
    DiaryRichTextInlineNode,
    DiaryRichTextMark,
    DiaryRichTextTextNode
} from "./rich-text-document"
import styles from "./lexical-diary.module.css"

type UploadedFile = {
    id: number
    fileName: string
    url: string
}

export type DiaryInlineImage = {
    fileId: number
    fileName: string
    url: string
}

type SerializedDiaryImageNode = Spread<{
    alt: string
    fileId: number
    src: string
}, SerializedLexicalNode>

const pictureSource = (url: string) => {
    const token = getSession().token
    const separator = url.includes("?") ? "&" : "?"
    return resolveServerUrl(token ? `${url}${separator}token=${encodeURIComponent(token)}` : url)
}

export class DiaryImageNode extends DecoratorNode<React.ReactNode> {
    __fileId: number
    __src: string
    __alt: string

    static getType(): string {
        return "diary-image"
    }

    static clone(node: DiaryImageNode): DiaryImageNode {
        return new DiaryImageNode(node.__fileId, node.__src, node.__alt, node.__key)
    }

    static importJSON(serializedNode: SerializedDiaryImageNode): DiaryImageNode {
        return $createDiaryImageNode({
            alt: serializedNode.alt,
            fileId: serializedNode.fileId,
            src: serializedNode.src
        })
    }

    constructor(fileId: number, src: string, alt: string, key?: NodeKey) {
        super(key)
        this.__fileId = fileId
        this.__src = src
        this.__alt = alt
    }

    exportJSON(): SerializedDiaryImageNode {
        return {
            alt: this.__alt,
            fileId: this.__fileId,
            src: this.__src,
            type: "diary-image",
            version: 1
        }
    }

    createDOM(_config: EditorConfig): HTMLElement {
        const element = document.createElement("figure")
        element.className = styles.imageFrame
        return element
    }

    updateDOM(): false {
        return false
    }

    decorate(_editor: LexicalEditor, _config: EditorConfig): React.ReactNode {
        return <img className={styles.inlineImage} src={pictureSource(this.__src)} alt={this.__alt} />
    }
}

function $createDiaryImageNode(payload: {
    alt: string
    fileId: number
    src: string
}): DiaryImageNode {
    return $applyNodeReplacement(new DiaryImageNode(payload.fileId, payload.src, payload.alt))
}

const fileUrl = (fileId: number) => `/api/files/${fileId}`

const parseJsonObject = (value?: string | null) => {
    try {
        if (!value) {
            return null
        }

        const parsed = JSON.parse(value)
        return parsed && typeof parsed === "object" ? parsed : null
    } catch (_error) {
        return null
    }
}

function isSerializedEditorStateValue(value: any) {
    return value && typeof value === "object" && value.root && Array.isArray(value.root.children)
}

function isSerializedEditorState(value: string) {
    return isSerializedEditorStateValue(parseJsonObject(value))
}

const isDiaryRichTextDocumentValue = (value: any): value is DiaryRichTextDocument => {
    return value
        && typeof value === "object"
        && value.version === DIARY_RICH_TEXT_DOCUMENT_VERSION
        && Array.isArray(value.content)
}

export const isDiaryRichTextBody = (value?: string | null) => {
    return isDiaryRichTextDocumentValue(parseJsonObject(value))
}

export const isLexicalDiaryBody = (value?: string | null) => {
    return typeof value === "string" && isSerializedEditorState(value)
}

export const isStructuredDiaryBody = (value?: string | null) => {
    return isDiaryRichTextBody(value) || isLexicalDiaryBody(value)
}

const marksFromLexicalFormat = (format: number): DiaryRichTextMark[] | undefined => {
    const marks: DiaryRichTextMark[] = []

    if (format & 1) {
        marks.push("bold")
    }

    if (format & 2) {
        marks.push("italic")
    }

    if (format & 8) {
        marks.push("underline")
    }

    return marks.length > 0 ? marks : undefined
}

const lexicalFormatFromMarks = (marks?: DiaryRichTextMark[]) => {
    let format = 0

    if (marks?.includes("bold")) {
        format |= 1
    }

    if (marks?.includes("italic")) {
        format |= 2
    }

    if (marks?.includes("underline")) {
        format |= 8
    }

    return format
}

const createTextNode = (node: DiaryRichTextTextNode) => ({
    detail: 0,
    format: lexicalFormatFromMarks(node.marks),
    mode: "normal",
    style: "",
    text: node.text,
    type: "text",
    version: 1
})

const createParagraphNode = (children: any[]) => ({
    children,
    direction: null,
    format: "",
    indent: 0,
    textFormat: 0,
    textStyle: "",
    type: "paragraph",
    version: 1
})

const createHeadingNode = (level: DiaryRichTextHeadingLevel, children: any[]) => ({
    children,
    direction: null,
    format: "",
    indent: 0,
    tag: level === 3 ? "h3" : "h2",
    type: "heading",
    version: 1
})

const inlineNodesToLexicalChildren = (nodes: DiaryRichTextInlineNode[]) => {
    return nodes.map(node => {
        if ("text" in node) {
            return createTextNode(node)
        }

        return {
            type: "linebreak",
            version: 1
        }
    })
}

const lexicalStateFromDiaryDocument = (document: DiaryRichTextDocument) => {
    const children = document.content.map(block => {
        if (block.type === "heading") {
            return createHeadingNode(block.level, inlineNodesToLexicalChildren(block.children))
        }

        if (block.type === "image") {
            return {
                alt: block.alt ?? "",
                fileId: block.fileId,
                src: fileUrl(block.fileId),
                type: "diary-image",
                version: 1
            }
        }

        return createParagraphNode(inlineNodesToLexicalChildren(block.children))
    })

    return {
        root: {
            children,
            direction: null,
            format: "",
            indent: 0,
            type: "root",
            version: 1
        }
    }
}

const inlineNodesFromLexicalChildren = (children: any[]): DiaryRichTextInlineNode[] => {
    return children.flatMap(child => {
        if (child?.type === "text") {
            const node: DiaryRichTextTextNode = {
                text: child.text ?? ""
            }
            const marks = marksFromLexicalFormat(Number(child.format ?? 0))

            if (marks) {
                node.marks = marks
            }

            return node.text === "" ? [] : [node]
        }

        if (child?.type === "linebreak") {
            return [{ type: "lineBreak" }]
        }

        if (Array.isArray(child?.children)) {
            return inlineNodesFromLexicalChildren(child.children)
        }

        return []
    })
}

const diaryDocumentFromLexicalState = (state: any): DiaryRichTextDocument => {
    const content: DiaryRichTextBlock[] = (state?.root?.children ?? []).flatMap((node: any) => {
        if (node?.type === "diary-image" && typeof node.fileId === "number") {
            return [{
                type: "image",
                fileId: node.fileId,
                alt: node.alt || undefined
            }]
        }

        if (node?.type === "heading") {
            return [{
                type: "heading",
                level: node.tag === "h3" ? 3 : 2,
                children: inlineNodesFromLexicalChildren(node.children ?? [])
            }]
        }

        if (node?.type === "paragraph") {
            return [{
                type: "paragraph",
                children: inlineNodesFromLexicalChildren(node.children ?? [])
            }]
        }

        if (Array.isArray(node?.children)) {
            return [{
                type: "paragraph",
                children: inlineNodesFromLexicalChildren(node.children)
            }]
        }

        return []
    })

    return {
        version: DIARY_RICH_TEXT_DOCUMENT_VERSION,
        content: content.length > 0 ? content : [{
            type: "paragraph",
            children: []
        }]
    }
}

const inlineNodesFromPlainText = (text: string): DiaryRichTextInlineNode[] => {
    if (!text) {
        return []
    }

    return text.split("\n").flatMap((part, index) => {
        const nodes: DiaryRichTextInlineNode[] = []

        if (index > 0) {
            nodes.push({ type: "lineBreak" })
        }

        if (part) {
            nodes.push({ text: part })
        }

        return nodes
    })
}

const createDiaryDocumentFromPlainTextAndImages = (text: string, images: DiaryInlineImage[]): DiaryRichTextDocument => ({
    version: DIARY_RICH_TEXT_DOCUMENT_VERSION,
    content: [
        {
            type: "paragraph",
            children: inlineNodesFromPlainText(text)
        },
        ...images.map(image => ({
            type: "image" as const,
            fileId: image.fileId,
            alt: image.fileName
        }))
    ]
})

const diaryDocumentFromBody = (body?: string | null, images: DiaryInlineImage[] = []): DiaryRichTextDocument => {
    const parsed = parseJsonObject(body)

    if (isDiaryRichTextDocumentValue(parsed)) {
        return parsed
    }

    if (isSerializedEditorStateValue(parsed)) {
        return diaryDocumentFromLexicalState(parsed)
    }

    return createDiaryDocumentFromPlainTextAndImages(body ?? "", images)
}

export const createDiaryBodyFromPlainTextAndImages = (text: string, images: DiaryInlineImage[]) => {
    return JSON.stringify(createDiaryDocumentFromPlainTextAndImages(text, images))
}

const createInitialEditorState = (value: string, images: DiaryInlineImage[]) => {
    return JSON.stringify(lexicalStateFromDiaryDocument(diaryDocumentFromBody(value, images)))
}

function DiaryToolbarPlugin(props: {
    disabled?: boolean
}) {
    const [editor] = useLexicalComposerContext()
    const [isUploading, setIsUploading] = useState(false)
    const [isBold, setIsBold] = useState(false)
    const [isItalic, setIsItalic] = useState(false)
    const inputRef = useRef<HTMLInputElement | null>(null)

    useEffect(() => {
        return editor.registerUpdateListener(({ editorState }) => {
            editorState.read(() => {
                const selection = $getSelection()

                if (!$isRangeSelection(selection)) {
                    setIsBold(false)
                    setIsItalic(false)
                    return
                }

                setIsBold(selection.hasFormat("bold"))
                setIsItalic(selection.hasFormat("italic"))
            })
        })
    }, [editor])

    const runToolbarAction = (event: React.MouseEvent<HTMLButtonElement>, action: () => void) => {
        event.preventDefault()
        action()
    }

    const toolbarButtonClass = (active = false) => {
        return active ? `${styles.toolbarButton} ${styles.toolbarButtonActive}` : styles.toolbarButton
    }

    const setBlockType = (type: "paragraph" | "h2" | "h3") => {
        editor.update(() => {
            const selection = $getSelection()

            if (!$isRangeSelection(selection)) {
                return
            }

            if (type === "paragraph") {
                $setBlocksType(selection, () => $createParagraphNode())
                selection.setFormat(0)
                selection.setStyle("")
                return
            }

            $setBlocksType(selection, () => $createHeadingNode(type))
        })
    }

    const uploadFiles = async (files: FileList | null) => {
        if (!files || files.length === 0) {
            return
        }

        setIsUploading(true)

        try {
            for (const file of Array.from(files)) {
                const formData = new FormData()
                formData.append("file", file)

                const response = await postFormData<UploadedFile>("/api/files", formData)

                if (response.error) {
                    throw new Error(response.error.message)
                }

                if (response.payload) {
                    editor.update(() => {
                        $insertNodes([
                            $createDiaryImageNode({
                                alt: response.payload?.fileName ?? file.name,
                                fileId: response.payload.id,
                                src: response.payload.url
                            }),
                            $createParagraphNode()
                        ])
                    })
                }
            }
        } finally {
            setIsUploading(false)
            if (inputRef.current) {
                inputRef.current.value = ""
            }
        }
    }

    return (
        <div className={styles.toolbar}>
            <button
                className={toolbarButtonClass(isBold)}
                disabled={props.disabled}
                type="button"
                onMouseDown={event => runToolbarAction(event, () => editor.dispatchCommand(FORMAT_TEXT_COMMAND, "bold"))}>
                Bold
            </button>
            <button
                className={toolbarButtonClass(isItalic)}
                disabled={props.disabled}
                type="button"
                onMouseDown={event => runToolbarAction(event, () => editor.dispatchCommand(FORMAT_TEXT_COMMAND, "italic"))}>
                Italic
            </button>
            <span className={styles.toolbarDivider} />
            <button
                className={styles.toolbarButton}
                disabled={props.disabled}
                type="button"
                onMouseDown={event => runToolbarAction(event, () => setBlockType("h2"))}>
                H2
            </button>
            <button
                className={styles.toolbarButton}
                disabled={props.disabled}
                type="button"
                onMouseDown={event => runToolbarAction(event, () => setBlockType("h3"))}>
                H3
            </button>
            <button
                className={styles.toolbarButton}
                disabled={props.disabled}
                type="button"
                onMouseDown={event => runToolbarAction(event, () => setBlockType("paragraph"))}>
                Text
            </button>
            <span className={styles.toolbarDivider} />
            <button
                className={styles.toolbarButton}
                disabled={props.disabled || isUploading}
                type="button"
                onClick={() => inputRef.current?.click()}>
                {isUploading ? "Adding..." : "Add picture"}
            </button>
            <input
                ref={inputRef}
                className={styles.hiddenFileInput}
                type="file"
                accept="image/*"
                multiple
                onChange={event => uploadFiles(event.target.files)}
            />
        </div>
    )
}

export function DiaryLexicalEditor(props: {
    initialImages?: DiaryInlineImage[]
    onChange: (value: string) => void
    value: string
}) {
    const initialConfig = {
        editable: true,
        editorState: createInitialEditorState(props.value, props.initialImages ?? []),
        namespace: "DiaryEditor",
        nodes: [DiaryImageNode, HeadingNode],
        onError(error: Error) {
            throw error
        },
        theme: {
            heading: {
                h2: styles.editorHeading,
                h3: styles.editorSubheading
            },
            paragraph: styles.editorParagraph
        }
    }

    return (
        <LexicalComposer initialConfig={initialConfig}>
            <div className={styles.editorShell}>
                <DiaryToolbarPlugin />
                <RichTextPlugin
                    contentEditable={<ContentEditable className={styles.editorInput} />}
                    placeholder={<div className={styles.placeholder}>Write as much as you want...</div>}
                    ErrorBoundary={LexicalErrorBoundary}
                />
                <HistoryPlugin />
                <OnChangePlugin onChange={editorState => {
                    props.onChange(JSON.stringify(diaryDocumentFromLexicalState(editorState.toJSON())))
                }} />
            </div>
        </LexicalComposer>
    )
}

export const getDiaryBodyFileIds = (body: string) => {
    return Array.from(new Set(diaryDocumentFromBody(body).content
        .filter((block): block is Extract<DiaryRichTextBlock, { type: "image" }> => block.type === "image")
        .map(block => block.fileId)))
}

export const hasDiaryBodyContent = (body?: string | null) => {
    if (!body) {
        return false
    }

    if (!isStructuredDiaryBody(body)) {
        return body.trim() !== ""
    }

    return diaryDocumentFromBody(body).content.some(block => {
        if (block.type === "image") {
            return true
        }

        return block.children.some(child => "text" in child && child.text.trim() !== "")
    })
}

const renderInlineNode = (node: DiaryRichTextInlineNode, key: string) => {
    if (!("text" in node)) {
        return <br key={key} />
    }

    let content: React.ReactNode = node.text

    if (node.marks?.includes("bold")) {
        content = <strong>{content}</strong>
    }

    if (node.marks?.includes("italic")) {
        content = <em>{content}</em>
    }

    if (node.marks?.includes("underline")) {
        content = <u>{content}</u>
    }

    return <React.Fragment key={key}>{content}</React.Fragment>
}

const renderBlock = (block: DiaryRichTextBlock, key: string): React.ReactNode => {
    if (block.type === "image") {
        return (
            <figure className={styles.readonlyImageFrame} key={key}>
                <img className={styles.inlineImage} src={pictureSource(fileUrl(block.fileId))} alt={block.alt ?? ""} />
            </figure>
        )
    }

    const children = block.children.map((child, index) => renderInlineNode(child, `${key}-${index}`))

    if (block.type === "paragraph") {
        return (
            <p className={styles.readonlyParagraph} key={key}>
                {children}
            </p>
        )
    }

    if (block.level === 3) {
        return <h3 className={styles.readonlySubheading} key={key}>{children}</h3>
    }

    return <h2 className={styles.readonlyHeading} key={key}>{children}</h2>
}

export function DiaryBodyRenderer(props: {
    body: string
}) {
    if (!isStructuredDiaryBody(props.body)) {
        return <div className={styles.plainBody}>{props.body}</div>
    }

    const document = diaryDocumentFromBody(props.body)

    return (
        <div className={styles.readonlyBody}>
            {document.content.map((block, index) => renderBlock(block, String(index)))}
        </div>
    )
}
