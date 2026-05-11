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
    $getNodeByKey,
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

    decorate(editor: LexicalEditor, _config: EditorConfig): React.ReactNode {
        return <EditableDiaryImage editor={editor} nodeKey={this.__key} fileId={this.__fileId} alt={this.__alt} />
    }
}

function EditableDiaryImage(props: {
    alt: string
    editor: LexicalEditor
    fileId: number
    nodeKey: NodeKey
}) {
    const removeImage = () => {
        props.editor.update(() => {
            $getNodeByKey(props.nodeKey)?.remove()
        })
    }

    return (
        <div className={styles.editableImageFrame}>
            <img className={styles.inlineImage} src={pictureSource(fileUrl(props.fileId))} alt={props.alt} />
            <button
                aria-label="Remove picture"
                className={styles.removeImageButton}
                onClick={removeImage}
                title="Remove picture"
                type="button">
                ×
            </button>
        </div>
    )
}

function $createDiaryImageNode(payload: {
    alt: string
    fileId: number
    src: string
}): DiaryImageNode {
    return $applyNodeReplacement(new DiaryImageNode(payload.fileId, payload.src, payload.alt))
}

const fileUrl = (fileId: number, size: "thumb" | "medium" | "large" | "original" = "medium") => `/api/files/${fileId}?size=${size}`
const maxImageDimension = 1600
const imageUploadQuality = 0.82

const isHeicImageFile = (file: File) => {
    const name = file.name.toLowerCase()

    return file.type === "image/heic"
        || file.type === "image/heif"
        || name.endsWith(".heic")
        || name.endsWith(".heif")
}

const convertHeicToJpeg = async (file: File) => {
    const heic2any = (await import("heic2any")).default
    const result = await heic2any({
        blob: file,
        quality: imageUploadQuality,
        toType: "image/jpeg"
    })
    const blob = Array.isArray(result) ? result[0] : result

    return new File([blob], file.name.replace(/\.[^.]+$/, ".jpg"), {
        lastModified: file.lastModified,
        type: "image/jpeg"
    })
}

const resizeImageForUpload = async (file: File): Promise<File> => {
    if (isHeicImageFile(file)) {
        return resizeImageForUpload(await convertHeicToJpeg(file))
    }

    if (!file.type.startsWith("image/") || file.type === "image/gif" || file.type === "image/svg+xml") {
        return file
    }

    return new Promise<File>(resolve => {
        const image = new Image()
        const objectUrl = URL.createObjectURL(file)

        image.onload = () => {
            URL.revokeObjectURL(objectUrl)

            const scale = Math.min(1, maxImageDimension / Math.max(image.width, image.height))

            if (scale >= 1) {
                resolve(file)
                return
            }

            const canvas = document.createElement("canvas")
            canvas.width = Math.round(image.width * scale)
            canvas.height = Math.round(image.height * scale)
            const context = canvas.getContext("2d")

            if (!context) {
                resolve(file)
                return
            }

            context.drawImage(image, 0, 0, canvas.width, canvas.height)
            canvas.toBlob(blob => {
                if (!blob) {
                    resolve(file)
                    return
                }

                const resizedFile = new File([blob], file.name.replace(/\.[^.]+$/, ".jpg"), {
                    lastModified: file.lastModified,
                    type: "image/jpeg"
                })

                resolve(resizedFile.size < file.size ? resizedFile : file)
            }, "image/jpeg", imageUploadQuality)
        }

        image.onerror = () => {
            URL.revokeObjectURL(objectUrl)
            resolve(file)
        }

        image.src = objectUrl
    })
}

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

const blocksFromLexicalNode = (node: any): DiaryRichTextBlock[] => {
    if (node?.type === "diary-image" && typeof node.fileId === "number") {
        return [{
            type: "image",
            fileId: node.fileId,
            alt: node.alt || undefined
        }]
    }

    if (node?.type === "heading") {
        const nestedImages = (node.children ?? []).flatMap(blocksFromLexicalNode)
        const textChildren = inlineNodesFromLexicalChildren(node.children ?? [])
        const heading: DiaryRichTextBlock[] = textChildren.length > 0 ? [{
            type: "heading",
            level: node.tag === "h3" ? 3 : 2,
            children: textChildren
        }] : []

        return [...heading, ...nestedImages]
    }

    if (node?.type === "paragraph") {
        const nestedImages = (node.children ?? []).flatMap(blocksFromLexicalNode)
        const paragraph: DiaryRichTextBlock = {
            type: "paragraph",
            children: inlineNodesFromLexicalChildren(node.children ?? [])
        }

        return [paragraph, ...nestedImages]
    }

    if (Array.isArray(node?.children)) {
        return node.children.flatMap(blocksFromLexicalNode)
    }

    return []
}

const diaryDocumentFromLexicalState = (state: any): DiaryRichTextDocument => {
    const content: DiaryRichTextBlock[] = (state?.root?.children ?? []).flatMap(blocksFromLexicalNode)

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

    const uploadFiles = async (files: FileList | File[] | null) => {
        const selectedFiles = Array.from(files ?? []).filter(file => file.type.startsWith("image/"))

        if (selectedFiles.length === 0) {
            return
        }

        setIsUploading(true)

        try {
            for (const file of selectedFiles) {
                const uploadFile = await resizeImageForUpload(file)
                const formData = new FormData()
                formData.append("file", uploadFile)

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
                                src: fileUrl(response.payload.id)
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

function DiaryImageDropPlugin() {
    const [editor] = useLexicalComposerContext()
    const [isDraggingImage, setIsDraggingImage] = useState(false)
    const [isUploading, setIsUploading] = useState(false)
    const dragDepth = useRef(0)

    const imageFilesFromDataTransfer = (dataTransfer: DataTransfer) => {
        return Array.from(dataTransfer.files).filter(file => file.type.startsWith("image/"))
    }

    const hasImageDragItems = (dataTransfer: DataTransfer) => {
        return Array.from(dataTransfer.items ?? []).some(item => item.kind === "file" && item.type.startsWith("image/"))
    }

    const uploadFiles = async (files: File[]) => {
        if (files.length === 0) {
            return
        }

        setIsUploading(true)

        try {
            for (const file of files) {
                const uploadFile = await resizeImageForUpload(file)
                const formData = new FormData()
                formData.append("file", uploadFile)

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
                                src: fileUrl(response.payload.id)
                            }),
                            $createParagraphNode()
                        ])
                    })
                }
            }
        } finally {
            setIsUploading(false)
        }
    }

    return (
        <div
            className={isDraggingImage ? `${styles.dropTarget} ${styles.dropTargetActive}` : styles.dropTarget}
            onDragEnter={event => {
                if (!hasImageDragItems(event.dataTransfer)) {
                    return
                }

                event.preventDefault()
                dragDepth.current += 1
                setIsDraggingImage(true)
            }}
            onDragOver={event => {
                if (!hasImageDragItems(event.dataTransfer)) {
                    return
                }

                event.preventDefault()
                event.dataTransfer.dropEffect = "copy"
            }}
            onDragLeave={event => {
                if (!hasImageDragItems(event.dataTransfer)) {
                    return
                }

                event.preventDefault()
                dragDepth.current = Math.max(0, dragDepth.current - 1)

                if (dragDepth.current === 0) {
                    setIsDraggingImage(false)
                }
            }}
            onDrop={event => {
                const files = imageFilesFromDataTransfer(event.dataTransfer)

                if (files.length === 0) {
                    return
                }

                event.preventDefault()
                dragDepth.current = 0
                setIsDraggingImage(false)
                void uploadFiles(files)
            }}>
            <RichTextPlugin
                contentEditable={<ContentEditable className={styles.editorInput} />}
                placeholder={<div className={styles.placeholder}>Write as much as you want...</div>}
                ErrorBoundary={LexicalErrorBoundary}
            />
            {(isDraggingImage || isUploading) && (
                <div className={styles.dropOverlay}>
                    {isUploading ? "Adding pictures..." : "Drop pictures to add them"}
                </div>
            )}
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
                <DiaryImageDropPlugin />
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

const renderBlock = (block: DiaryRichTextBlock, key: string, onImageClick?: (image: { alt?: string, fileId: number }) => void): React.ReactNode => {
    if (block.type === "image") {
        const image = <img className={styles.inlineImage} src={pictureSource(fileUrl(block.fileId))} alt={block.alt ?? ""} />

        return (
            <figure className={styles.readonlyImageFrame} key={key}>
                {onImageClick ? (
                    <button className={styles.imagePreviewButton} onClick={() => onImageClick({ alt: block.alt, fileId: block.fileId })} type="button">
                        {image}
                    </button>
                ) : image}
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
    onImageClick?: (image: { alt?: string, fileId: number }) => void
}) {
    if (!isStructuredDiaryBody(props.body)) {
        return <div className={styles.plainBody}>{props.body}</div>
    }

    const document = diaryDocumentFromBody(props.body)

    return (
        <div className={styles.readonlyBody}>
            {document.content.map((block, index) => renderBlock(block, String(index), props.onImageClick))}
        </div>
    )
}
