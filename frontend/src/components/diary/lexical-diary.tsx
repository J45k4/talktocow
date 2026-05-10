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
    $createTextNode,
    $getSelection,
    $getRoot,
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

function isSerializedEditorState(value: string) {
    try {
        const parsed = JSON.parse(value)
        return parsed && typeof parsed === "object" && parsed.root && Array.isArray(parsed.root.children)
    } catch (_error) {
        return false
    }
}

export const isLexicalDiaryBody = (value?: string | null) => {
    return typeof value === "string" && isSerializedEditorState(value)
}

export const createDiaryBodyFromPlainTextAndImages = (text: string, images: DiaryInlineImage[]) => {
    return JSON.stringify({
        root: {
            children: [
                {
                    children: text
                        ? [{
                            detail: 0,
                            format: 0,
                            mode: "normal",
                            style: "",
                            text,
                            type: "text",
                            version: 1
                        }]
                        : [],
                    direction: null,
                    format: "",
                    indent: 0,
                    textFormat: 0,
                    textStyle: "",
                    type: "paragraph",
                    version: 1
                },
                ...images.map(image => ({
                    alt: image.fileName,
                    fileId: image.fileId,
                    src: image.url,
                    type: "diary-image",
                    version: 1
                }))
            ],
            direction: null,
            format: "",
            indent: 0,
            type: "root",
            version: 1
        }
    })
}

const createInitialEditorState = (value: string, images: DiaryInlineImage[]) => {
    if (isSerializedEditorState(value)) {
        return value
    }

    return (_editor: LexicalEditor) => {
        const root = $getRoot()
        root.clear()

        const paragraph = $createParagraphNode()
        if (value) {
            paragraph.append($createTextNode(value))
        }
        root.append(paragraph)

        for (const image of images) {
            root.append($createDiaryImageNode({
                alt: image.fileName,
                fileId: image.fileId,
                src: image.url
            }))
        }
    }
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
                    props.onChange(JSON.stringify(editorState.toJSON()))
                }} />
            </div>
        </LexicalComposer>
    )
}

const collectFileIdsFromNode = (node: any, result: Set<number>) => {
    if (node?.type === "diary-image" && typeof node.fileId === "number") {
        result.add(node.fileId)
    }

    if (Array.isArray(node?.children)) {
        node.children.forEach((child: any) => collectFileIdsFromNode(child, result))
    }
}

export const getDiaryBodyFileIds = (body: string) => {
    const result = new Set<number>()

    if (!isLexicalDiaryBody(body)) {
        return []
    }

    collectFileIdsFromNode(JSON.parse(body).root, result)
    return Array.from(result)
}

const getPlainTextFromNode = (node: any): string => {
    if (node?.type === "text") {
        return node.text ?? ""
    }

    if (node?.type === "linebreak") {
        return "\n"
    }

    if (Array.isArray(node?.children)) {
        return node.children.map(getPlainTextFromNode).join("")
    }

    return ""
}

export const hasDiaryBodyContent = (body?: string | null) => {
    if (!body) {
        return false
    }

    if (!isLexicalDiaryBody(body)) {
        return body.trim() !== ""
    }

    const parsed = JSON.parse(body)
    return getDiaryBodyFileIds(body).length > 0 || getPlainTextFromNode(parsed.root).trim() !== ""
}

const renderTextNode = (node: any, key: string) => {
    let content: React.ReactNode = node.text ?? ""
    const format = Number(node.format ?? 0)

    if (format & 1) {
        content = <strong>{content}</strong>
    }

    if (format & 2) {
        content = <em>{content}</em>
    }

    if (format & 8) {
        content = <u>{content}</u>
    }

    return <React.Fragment key={key}>{content}</React.Fragment>
}

const renderSerializedNode = (node: any, key: string): React.ReactNode => {
    if (node?.type === "text") {
        return renderTextNode(node, key)
    }

    if (node?.type === "linebreak") {
        return <br key={key} />
    }

    if (node?.type === "diary-image") {
        return (
            <figure className={styles.readonlyImageFrame} key={key}>
                <img className={styles.inlineImage} src={pictureSource(node.src)} alt={node.alt ?? ""} />
            </figure>
        )
    }

    if (node?.type === "paragraph") {
        return (
            <p className={styles.readonlyParagraph} key={key}>
                {(node.children ?? []).map((child: any, index: number) => renderSerializedNode(child, `${key}-${index}`))}
            </p>
        )
    }

    if (node?.type === "heading") {
        const children = (node.children ?? []).map((child: any, index: number) => renderSerializedNode(child, `${key}-${index}`))

        if (node.tag === "h3") {
            return <h3 className={styles.readonlySubheading} key={key}>{children}</h3>
        }

        return <h2 className={styles.readonlyHeading} key={key}>{children}</h2>
    }

    if (Array.isArray(node?.children)) {
        return node.children.map((child: any, index: number) => renderSerializedNode(child, `${key}-${index}`))
    }

    return null
}

export function DiaryBodyRenderer(props: {
    body: string
}) {
    if (!isLexicalDiaryBody(props.body)) {
        return <div className={styles.plainBody}>{props.body}</div>
    }

    const parsed = JSON.parse(props.body)

    return (
        <div className={styles.readonlyBody}>
            {parsed.root.children.map((child: any, index: number) => renderSerializedNode(child, String(index)))}
        </div>
    )
}
