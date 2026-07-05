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
import { DiaryImageSize, diaryFileImageUrl, diaryFileUrl, diaryImageSource, diaryImageVariantUrl } from "./diary-image-source"
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

export type DiaryInlineAudio = {
    durationMs?: number
    fileId: number
    fileName: string
}

export type DiaryInlineVideo = {
    fileId: number
    fileName: string
}

type SerializedDiaryImageNode = Spread<{
    alt: string
    fileId: number
    src: string
}, SerializedLexicalNode>

type SerializedDiaryAudioNode = Spread<{
    durationMs?: number
    fileId: number
    fileName: string
    src: string
}, SerializedLexicalNode>

type SerializedDiaryVideoNode = Spread<{
    fileId: number
    fileName: string
    src: string
}, SerializedLexicalNode>

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
            <img className={styles.inlineImage} src={diaryImageSource(diaryFileImageUrl(props.fileId))} alt={props.alt} />
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

export class DiaryAudioNode extends DecoratorNode<React.ReactNode> {
    __durationMs?: number
    __fileId: number
    __fileName: string
    __src: string

    static getType(): string {
        return "diary-audio"
    }

    static clone(node: DiaryAudioNode): DiaryAudioNode {
        return new DiaryAudioNode(node.__fileId, node.__src, node.__fileName, node.__durationMs, node.__key)
    }

    static importJSON(serializedNode: SerializedDiaryAudioNode): DiaryAudioNode {
        return $createDiaryAudioNode({
            durationMs: serializedNode.durationMs,
            fileId: serializedNode.fileId,
            fileName: serializedNode.fileName,
            src: serializedNode.src
        })
    }

    constructor(fileId: number, src: string, fileName: string, durationMs?: number, key?: NodeKey) {
        super(key)
        this.__durationMs = durationMs
        this.__fileId = fileId
        this.__fileName = fileName
        this.__src = src
    }

    exportJSON(): SerializedDiaryAudioNode {
        return {
            durationMs: this.__durationMs,
            fileId: this.__fileId,
            fileName: this.__fileName,
            src: this.__src,
            type: "diary-audio",
            version: 1
        }
    }

    createDOM(_config: EditorConfig): HTMLElement {
        const element = document.createElement("figure")
        element.className = styles.audioFrame
        return element
    }

    updateDOM(): false {
        return false
    }

    decorate(editor: LexicalEditor, _config: EditorConfig): React.ReactNode {
        return <EditableDiaryAudio editor={editor} nodeKey={this.__key} fileName={this.__fileName} fileId={this.__fileId} durationMs={this.__durationMs} />
    }
}

function formatAudioDuration(durationMs?: number) {
    if (!durationMs || durationMs < 1000) {
        return ""
    }

    const totalSeconds = Math.round(durationMs / 1000)
    const minutes = Math.floor(totalSeconds / 60)
    const seconds = totalSeconds % 60

    return `${minutes}:${seconds.toString().padStart(2, "0")}`
}

function EditableDiaryAudio(props: {
    durationMs?: number
    editor: LexicalEditor
    fileId: number
    fileName: string
    nodeKey: NodeKey
}) {
    const removeAudio = () => {
        props.editor.update(() => {
            $getNodeByKey(props.nodeKey)?.remove()
        })
    }

    return (
        <div className={styles.editableAudioFrame}>
            <div className={styles.audioMeta}>
                <span>{props.fileName}</span>
                {props.durationMs && <span>{formatAudioDuration(props.durationMs)}</span>}
            </div>
            <audio className={styles.audioPlayer} controls preload="metadata" src={diaryImageSource(diaryFileUrl(props.fileId))} />
            <button
                aria-label="Remove recording"
                className={styles.removeAudioButton}
                onClick={removeAudio}
                title="Remove recording"
                type="button">
                ×
            </button>
        </div>
    )
}

function $createDiaryAudioNode(payload: {
    durationMs?: number
    fileId: number
    fileName: string
    src: string
}): DiaryAudioNode {
    return $applyNodeReplacement(new DiaryAudioNode(payload.fileId, payload.src, payload.fileName, payload.durationMs))
}

export class DiaryVideoNode extends DecoratorNode<React.ReactNode> {
    __fileId: number
    __fileName: string
    __src: string

    static getType(): string {
        return "diary-video"
    }

    static clone(node: DiaryVideoNode): DiaryVideoNode {
        return new DiaryVideoNode(node.__fileId, node.__src, node.__fileName, node.__key)
    }

    static importJSON(serializedNode: SerializedDiaryVideoNode): DiaryVideoNode {
        return $createDiaryVideoNode({
            fileId: serializedNode.fileId,
            fileName: serializedNode.fileName,
            src: serializedNode.src
        })
    }

    constructor(fileId: number, src: string, fileName: string, key?: NodeKey) {
        super(key)
        this.__fileId = fileId
        this.__fileName = fileName
        this.__src = src
    }

    exportJSON(): SerializedDiaryVideoNode {
        return {
            fileId: this.__fileId,
            fileName: this.__fileName,
            src: this.__src,
            type: "diary-video",
            version: 1
        }
    }

    createDOM(_config: EditorConfig): HTMLElement {
        const element = document.createElement("figure")
        element.className = styles.videoFrame
        return element
    }

    updateDOM(): false {
        return false
    }

    decorate(editor: LexicalEditor, _config: EditorConfig): React.ReactNode {
        return <EditableDiaryVideo editor={editor} nodeKey={this.__key} fileName={this.__fileName} fileId={this.__fileId} />
    }
}

function EditableDiaryVideo(props: {
    editor: LexicalEditor
    fileId: number
    fileName: string
    nodeKey: NodeKey
}) {
    const removeVideo = () => {
        props.editor.update(() => {
            $getNodeByKey(props.nodeKey)?.remove()
        })
    }

    return (
        <div className={styles.editableVideoFrame}>
            <div className={styles.videoMeta}>
                <span>{props.fileName}</span>
            </div>
            <video className={styles.videoPlayer} controls playsInline preload="metadata" src={diaryImageSource(diaryFileUrl(props.fileId))} />
            <button
                aria-label="Remove video"
                className={styles.removeVideoButton}
                onClick={removeVideo}
                title="Remove video"
                type="button">
                ×
            </button>
        </div>
    )
}

function $createDiaryVideoNode(payload: {
    fileId: number
    fileName: string
    src: string
}): DiaryVideoNode {
    return $applyNodeReplacement(new DiaryVideoNode(payload.fileId, payload.src, payload.fileName))
}

const maxImageDimension = 1600
const imageUploadQuality = 0.82

const isHeicImageFile = (file: File) => {
    const name = file.name.toLowerCase()

    return file.type === "image/heic"
        || file.type === "image/heif"
        || name.endsWith(".heic")
        || name.endsWith(".heif")
}

const isSupportedImageUploadFile = (file: File) => {
    return file.type.startsWith("image/") || isHeicImageFile(file)
}

const supportedVideoUploadMimeTypes = new Set([
    "video/mp4",
    "video/quicktime",
    "video/webm"
])

const isSupportedVideoUploadFile = (file: File) => {
    const name = file.name.toLowerCase()
    return supportedVideoUploadMimeTypes.has(file.type)
        || name.endsWith(".mp4")
        || name.endsWith(".mov")
        || name.endsWith(".webm")
}

const supportedAudioUploadMimeTypes = new Set([
    "audio/mp4",
    "audio/ogg",
    "audio/wav",
    "audio/webm",
    "video/webm"
])

const isSupportedAudioUploadFile = (file: File) => {
    const name = file.name.toLowerCase()
    return supportedAudioUploadMimeTypes.has(file.type)
        || name.endsWith(".m4a")
        || name.endsWith(".ogg")
        || name.endsWith(".wav")
        || name.endsWith(".webm")
}

const isSupportedMediaUploadFile = (file: File) => {
    return isSupportedImageUploadFile(file) || isSupportedAudioUploadFile(file) || isSupportedVideoUploadFile(file)
}

const resizeImageForUpload = async (file: File): Promise<File> => {
    if (isHeicImageFile(file)) {
        return file
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

const insertUploadedImage = (editor: LexicalEditor, uploadedFile: UploadedFile, originalFile: File) => {
    editor.update(() => {
        $insertNodes([
            $createDiaryImageNode({
                alt: uploadedFile.fileName ?? originalFile.name,
                fileId: uploadedFile.id,
                src: diaryFileImageUrl(uploadedFile.id)
            }),
            $createParagraphNode()
        ])
    })
}

const insertUploadedAudio = (editor: LexicalEditor, uploadedFile: UploadedFile, originalFile: File) => {
    editor.update(() => {
        $insertNodes([
            $createDiaryAudioNode({
                fileId: uploadedFile.id,
                fileName: uploadedFile.fileName ?? originalFile.name,
                src: diaryFileUrl(uploadedFile.id)
            }),
            $createParagraphNode()
        ])
    })
}

const insertUploadedVideo = (editor: LexicalEditor, uploadedFile: UploadedFile, originalFile: File) => {
    editor.update(() => {
        $insertNodes([
            $createDiaryVideoNode({
                fileId: uploadedFile.id,
                fileName: uploadedFile.fileName ?? originalFile.name,
                src: diaryFileUrl(uploadedFile.id)
            }),
            $createParagraphNode()
        ])
    })
}

type UploadMediaKind = "audio" | "image" | "video"

const mediaUploadKind = (file: File, preferredKind?: UploadMediaKind): UploadMediaKind => {
    const canUploadAudio = isSupportedAudioUploadFile(file)
    const canUploadVideo = isSupportedVideoUploadFile(file)

    if (preferredKind === "audio" && isSupportedAudioUploadFile(file)) {
        return "audio"
    }

    if (preferredKind === "image" && isSupportedImageUploadFile(file)) {
        return "image"
    }

    if (preferredKind === "video" && isSupportedVideoUploadFile(file)) {
        return "video"
    }

    if (canUploadAudio && file.type.startsWith("audio/")) {
        return "audio"
    }

    if (canUploadVideo && file.type.startsWith("video/")) {
        return "video"
    }

    if (canUploadAudio && !canUploadVideo) {
        return "audio"
    }

    if (canUploadVideo) {
        return "video"
    }

    return "image"
}

type UploadState = {
    current: number
    total: number
}

const uploadMediaFile = async (editor: LexicalEditor, file: File, preferredKind?: UploadMediaKind) => {
    const kind = mediaUploadKind(file, preferredKind)
    const uploadFile = kind === "image" ? await resizeImageForUpload(file) : file
    const formData = new FormData()
    formData.append(kind === "audio" ? "recording" : kind === "video" ? "video" : "file", uploadFile)

    const response = await postFormData<UploadedFile>(
        kind === "audio" ? "/api/diary/recording" : kind === "video" ? "/api/diary/video" : "/api/files",
        formData
    )

    if (response.error) {
        throw new Error(response.error.message)
    }

    if (!response.payload) {
        throw new Error("Upload did not return a file")
    }

    if (kind === "audio") {
        insertUploadedAudio(editor, response.payload, file)
        return
    }

	if (kind === "video") {
		insertUploadedVideo(editor, response.payload, file)
		return
	}

    insertUploadedImage(editor, response.payload, file)
}

function UploadIndicator(props: {
    uploadState: UploadState | null
}) {
    if (!props.uploadState) {
        return null
    }

    const label = props.uploadState.total > 1
        ? `Uploading ${props.uploadState.current} of ${props.uploadState.total}`
        : "Uploading"

    return (
        <div className={styles.uploadIndicator} role="status" aria-live="polite">
            <span className={styles.uploadSpinner} aria-hidden="true" />
            <span>{label}</span>
        </div>
    )
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
                src: diaryFileImageUrl(block.fileId),
                type: "diary-image",
                version: 1
            }
        }

        if (block.type === "audio") {
            return {
                durationMs: block.durationMs,
                fileId: block.fileId,
                fileName: block.fileName ?? "Diary recording",
                src: diaryFileUrl(block.fileId),
                type: "diary-audio",
                version: 1
            }
        }

        if (block.type === "video") {
            return {
                fileId: block.fileId,
                fileName: block.fileName ?? "Diary video",
                src: diaryFileUrl(block.fileId),
                type: "diary-video",
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

    if (node?.type === "diary-audio" && typeof node.fileId === "number") {
        return [{
            type: "audio",
            fileId: node.fileId,
            fileName: node.fileName || undefined,
            durationMs: typeof node.durationMs === "number" ? node.durationMs : undefined
        }]
    }

    if (node?.type === "diary-video" && typeof node.fileId === "number") {
        return [{
            type: "video",
            fileId: node.fileId,
            fileName: node.fileName || undefined
        }]
    }

    if (node?.type === "heading") {
        const nestedMedia = (node.children ?? []).flatMap(blocksFromLexicalNode)
        const textChildren = inlineNodesFromLexicalChildren(node.children ?? [])
        const heading: DiaryRichTextBlock[] = textChildren.length > 0 ? [{
            type: "heading",
            level: node.tag === "h3" ? 3 : 2,
            children: textChildren
        }] : []

        return [...heading, ...nestedMedia]
    }

    if (node?.type === "paragraph") {
        const nestedMedia = (node.children ?? []).flatMap(blocksFromLexicalNode)
        const paragraph: DiaryRichTextBlock = {
            type: "paragraph",
            children: inlineNodesFromLexicalChildren(node.children ?? [])
        }

        return [paragraph, ...nestedMedia]
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

const createDiaryDocumentFromPlainTextImagesAudioAndVideo = (text: string, images: DiaryInlineImage[], audio: DiaryInlineAudio[] = [], videos: DiaryInlineVideo[] = []): DiaryRichTextDocument => ({
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
        })),
        ...audio.map(recording => ({
            type: "audio" as const,
            durationMs: recording.durationMs,
            fileId: recording.fileId,
            fileName: recording.fileName
        })),
        ...videos.map(video => ({
            type: "video" as const,
            fileId: video.fileId,
            fileName: video.fileName
        }))
    ]
})

const diaryDocumentFromBody = (body?: string | null, images: DiaryInlineImage[] = [], audio: DiaryInlineAudio[] = [], videos: DiaryInlineVideo[] = []): DiaryRichTextDocument => {
    const parsed = parseJsonObject(body)

    if (isDiaryRichTextDocumentValue(parsed)) {
        return parsed
    }

    if (isSerializedEditorStateValue(parsed)) {
        return diaryDocumentFromLexicalState(parsed)
    }

    return createDiaryDocumentFromPlainTextImagesAudioAndVideo(body ?? "", images, audio, videos)
}

export const createDiaryBodyFromPlainTextAndImages = (text: string, images: DiaryInlineImage[]) => {
    return JSON.stringify(createDiaryDocumentFromPlainTextImagesAudioAndVideo(text, images))
}

export const createDiaryBodyFromPlainTextImagesAndAudio = (text: string, images: DiaryInlineImage[], audio: DiaryInlineAudio[]) => {
    return JSON.stringify(createDiaryDocumentFromPlainTextImagesAudioAndVideo(text, images, audio))
}

export const createDiaryBodyFromPlainTextImagesAudioAndVideo = (text: string, images: DiaryInlineImage[], audio: DiaryInlineAudio[], videos: DiaryInlineVideo[]) => {
    return JSON.stringify(createDiaryDocumentFromPlainTextImagesAudioAndVideo(text, images, audio, videos))
}

const createInitialEditorState = (value: string, images: DiaryInlineImage[]) => {
    return JSON.stringify(lexicalStateFromDiaryDocument(diaryDocumentFromBody(value, images)))
}

function DiaryToolbarPlugin(props: {
    disabled?: boolean
    onUploadStateChange: (uploadState: UploadState | null) => void
    uploadState: UploadState | null
}) {
    const [editor] = useLexicalComposerContext()
    const [isBold, setIsBold] = useState(false)
    const [isItalic, setIsItalic] = useState(false)
    const imageInputRef = useRef<HTMLInputElement | null>(null)
    const audioInputRef = useRef<HTMLInputElement | null>(null)
    const videoInputRef = useRef<HTMLInputElement | null>(null)

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

    const uploadFiles = async (files: FileList | File[] | null, preferredKind?: UploadMediaKind) => {
        const selectedFiles = Array.from(files ?? []).filter(isSupportedMediaUploadFile)

        if (selectedFiles.length === 0) {
            return
        }

        props.onUploadStateChange({
            current: 1,
            total: selectedFiles.length
        })

        try {
            for (const [index, file] of selectedFiles.entries()) {
                props.onUploadStateChange({
                    current: index + 1,
                    total: selectedFiles.length
                })
                await uploadMediaFile(editor, file, preferredKind)
            }
        } finally {
            props.onUploadStateChange(null)
            if (imageInputRef.current) {
                imageInputRef.current.value = ""
            }
            if (audioInputRef.current) {
                audioInputRef.current.value = ""
            }
            if (videoInputRef.current) {
                videoInputRef.current.value = ""
            }
        }
    }

    const isUploading = props.uploadState !== null

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
                onClick={() => imageInputRef.current?.click()}>
                {isUploading ? "Adding..." : "Add picture"}
            </button>
            <button
                className={styles.toolbarButton}
                disabled={props.disabled || isUploading}
                type="button"
                onClick={() => videoInputRef.current?.click()}>
                {isUploading ? "Adding..." : "Add video"}
            </button>
            <button
                className={styles.toolbarButton}
                disabled={props.disabled || isUploading}
                type="button"
                onClick={() => audioInputRef.current?.click()}>
                {isUploading ? "Adding..." : "Add audio"}
            </button>
            <input
                ref={imageInputRef}
                className={styles.hiddenFileInput}
                type="file"
                accept="image/*,.heic,.heif"
                multiple
                onChange={event => uploadFiles(event.target.files, "image")}
            />
            <input
                ref={videoInputRef}
                className={styles.hiddenFileInput}
                type="file"
                accept="video/mp4,video/webm,video/quicktime,.mp4,.webm,.mov"
                multiple
                onChange={event => uploadFiles(event.target.files, "video")}
            />
            <input
                ref={audioInputRef}
                className={styles.hiddenFileInput}
                type="file"
                accept="audio/mp4,audio/ogg,audio/wav,audio/webm,video/webm,.m4a,.ogg,.wav,.webm"
                multiple
                onChange={event => uploadFiles(event.target.files, "audio")}
            />
        </div>
    )
}

function DiaryMediaDropPlugin(props: {
    onUploadStateChange: (uploadState: UploadState | null) => void
    uploadState: UploadState | null
}) {
    const [editor] = useLexicalComposerContext()
    const [isDraggingMedia, setIsDraggingMedia] = useState(false)
    const dragDepth = useRef(0)

    const mediaFilesFromDataTransfer = (dataTransfer: DataTransfer) => {
        return Array.from(dataTransfer.files).filter(isSupportedMediaUploadFile)
    }

    const hasFileDragItems = (dataTransfer: DataTransfer) => {
        return Array.from(dataTransfer.items ?? []).some(item => item.kind === "file")
    }

    const uploadFiles = async (files: File[]) => {
        if (files.length === 0) {
            return
        }

        props.onUploadStateChange({
            current: 1,
            total: files.length
        })

        try {
            for (const [index, file] of files.entries()) {
                props.onUploadStateChange({
                    current: index + 1,
                    total: files.length
                })
                await uploadMediaFile(editor, file)
            }
        } finally {
            props.onUploadStateChange(null)
        }
    }

    const isUploading = props.uploadState !== null

    return (
        <div
            className={isDraggingMedia ? `${styles.dropTarget} ${styles.dropTargetActive}` : styles.dropTarget}
            onDragEnter={event => {
                if (!hasFileDragItems(event.dataTransfer)) {
                    return
                }

                event.preventDefault()
                dragDepth.current += 1
                setIsDraggingMedia(true)
            }}
            onDragOver={event => {
                if (!hasFileDragItems(event.dataTransfer)) {
                    return
                }

                event.preventDefault()
                event.dataTransfer.dropEffect = "copy"
            }}
            onDragLeave={event => {
                if (!hasFileDragItems(event.dataTransfer)) {
                    return
                }

                event.preventDefault()
                dragDepth.current = Math.max(0, dragDepth.current - 1)

                if (dragDepth.current === 0) {
                    setIsDraggingMedia(false)
                }
            }}
            onDrop={event => {
                const files = mediaFilesFromDataTransfer(event.dataTransfer)

                if (files.length === 0) {
                    return
                }

                event.preventDefault()
                dragDepth.current = 0
                setIsDraggingMedia(false)
                void uploadFiles(files)
            }}>
            <RichTextPlugin
                contentEditable={<ContentEditable className={styles.editorInput} />}
                placeholder={<div className={styles.placeholder}>Write as much as you want...</div>}
                ErrorBoundary={LexicalErrorBoundary}
            />
            {(isDraggingMedia || isUploading) && (
                <div className={styles.dropOverlay}>
					{isUploading ? "Adding media..." : "Drop pictures, audio, or videos to add them"}
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
    const [uploadState, setUploadState] = useState<UploadState | null>(null)
    const initialConfig = {
        editable: true,
        editorState: createInitialEditorState(props.value, props.initialImages ?? []),
        namespace: "DiaryEditor",
        nodes: [DiaryAudioNode, DiaryImageNode, DiaryVideoNode, HeadingNode],
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
                <DiaryToolbarPlugin onUploadStateChange={setUploadState} uploadState={uploadState} />
                <DiaryMediaDropPlugin onUploadStateChange={setUploadState} uploadState={uploadState} />
                <HistoryPlugin />
                <OnChangePlugin onChange={editorState => {
                    props.onChange(JSON.stringify(diaryDocumentFromLexicalState(editorState.toJSON())))
                }} />
                <UploadIndicator uploadState={uploadState} />
            </div>
        </LexicalComposer>
    )
}

export const getDiaryBodyFileIds = (body: string) => {
    return Array.from(new Set(diaryDocumentFromBody(body).content
        .filter((block): block is Extract<DiaryRichTextBlock, { type: "image" }> => block.type === "image")
        .map(block => block.fileId)))
}

export const getDiaryBodyRecordingFileIds = (body: string) => {
    return Array.from(new Set(diaryDocumentFromBody(body).content
        .filter((block): block is Extract<DiaryRichTextBlock, { type: "audio" }> => block.type === "audio")
        .map(block => block.fileId)))
}

export const getDiaryBodyVideoFileIds = (body: string) => {
    return Array.from(new Set(diaryDocumentFromBody(body).content
        .filter((block): block is Extract<DiaryRichTextBlock, { type: "video" }> => block.type === "video")
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
        if (block.type === "audio" || block.type === "image" || block.type === "video") {
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

const diaryEntryPictureImageUrl = (entryId: number, pictureId: number, size: DiaryImageSize) => {
    return diaryImageVariantUrl(`/api/diary/entry/${entryId}/picture/${pictureId}`, size)
}

function ReadonlyDiaryAudio(props: {
    durationMs?: number
    fileId: number
    fileName?: string
    isCompact?: boolean
}) {
    if (props.isCompact) {
        return (
            <div className={styles.compactAudioBlock}>
                <span>Recording</span>
                {props.durationMs && <span>{formatAudioDuration(props.durationMs)}</span>}
            </div>
        )
    }

    return (
        <figure className={styles.readonlyAudioFrame}>
            <div className={styles.audioMeta}>
                <span>{props.fileName ?? "Diary recording"}</span>
                {props.durationMs && <span>{formatAudioDuration(props.durationMs)}</span>}
            </div>
            <audio className={styles.audioPlayer} controls preload="metadata" src={diaryImageSource(diaryFileUrl(props.fileId))} />
        </figure>
    )
}

function ReadonlyDiaryVideo(props: {
    fileId: number
    fileName?: string
    isCompact?: boolean
}) {
    if (props.isCompact) {
        return (
            <div className={styles.compactVideoBlock}>
                <span>Video</span>
            </div>
        )
    }

    return (
        <figure className={styles.readonlyVideoFrame}>
            <div className={styles.videoMeta}>
                <span>{props.fileName ?? "Diary video"}</span>
            </div>
            <video className={styles.videoPlayer} controls playsInline preload="metadata" src={diaryImageSource(diaryFileUrl(props.fileId))} />
        </figure>
    )
}

function ReadonlyDiaryImage(props: {
    alt: string
    className: string
    fallbackUrl?: string
    src: string
}) {
    const [src, setSrc] = useState(props.src)
    const [didFallback, setDidFallback] = useState(false)

    return (
        <img
            className={props.className}
            src={src}
            alt={props.alt}
            onError={() => {
                if (!props.fallbackUrl || didFallback) {
                    return
                }

                setDidFallback(true)
                setSrc(props.fallbackUrl)
            }}
        />
    )
}

export type DiaryImageFallback = {
    fileId?: number
    fileName?: string
    id: number
}

const findImageFallback = (block: Extract<DiaryRichTextBlock, { type: "image" }>, fallbacks?: DiaryImageFallback[]) => {
    return fallbacks?.find(fallback => fallback.fileId === block.fileId)
        ?? fallbacks?.find(fallback => fallback.fileName === block.alt)
}

const renderBlock = (block: DiaryRichTextBlock, key: string, options: {
    imageFallbacks?: DiaryImageFallback[]
    isCompact?: boolean
    legacyDiaryEntryId?: number
    onImageClick?: (image: { alt?: string, fileId: number }) => void
}): React.ReactNode => {
    if (block.type === "image") {
        const size = options.isCompact ? "thumb" : "medium"
        const imageUrl = diaryImageSource(diaryFileImageUrl(block.fileId, size))
        const fallback = findImageFallback(block, options.imageFallbacks)
        const fallbackPictureId = fallback?.id ?? block.fileId
        const fallbackUrl = options.legacyDiaryEntryId
            ? diaryImageSource(diaryEntryPictureImageUrl(options.legacyDiaryEntryId, fallbackPictureId, size))
            : undefined
        const image = <ReadonlyDiaryImage className={options.isCompact ? styles.compactInlineImage : styles.inlineImage} src={imageUrl} fallbackUrl={fallbackUrl} alt={block.alt ?? ""} />

        return (
            <figure className={options.isCompact ? styles.compactImageFrame : styles.readonlyImageFrame} key={key}>
                {options.onImageClick ? (
                    <button className={styles.imagePreviewButton} onClick={() => options.onImageClick?.({ alt: block.alt, fileId: block.fileId })} type="button">
                        {image}
                    </button>
                ) : image}
            </figure>
        )
    }

    if (block.type === "audio") {
        return <ReadonlyDiaryAudio durationMs={block.durationMs} fileId={block.fileId} fileName={block.fileName} isCompact={options.isCompact} key={key} />
    }

    if (block.type === "video") {
        return <ReadonlyDiaryVideo fileId={block.fileId} fileName={block.fileName} isCompact={options.isCompact} key={key} />
    }

    const children = block.children.map((child, index) => renderInlineNode(child, `${key}-${index}`))

    if (block.type === "paragraph") {
        return (
            <p className={options.isCompact ? styles.compactParagraph : styles.readonlyParagraph} key={key}>
                {children}
            </p>
        )
    }

    if (block.level === 3) {
        return <h3 className={options.isCompact ? styles.compactSubheading : styles.readonlySubheading} key={key}>{children}</h3>
    }

    return <h2 className={options.isCompact ? styles.compactHeading : styles.readonlyHeading} key={key}>{children}</h2>
}

export function DiaryBodyRenderer(props: {
    body: string
    imageFallbacks?: DiaryImageFallback[]
    legacyDiaryEntryId?: number
    maxBlocks?: number
    onImageClick?: (image: { alt?: string, fileId: number }) => void
    variant?: "full" | "compact"
}) {
    const isCompact = props.variant === "compact"

    if (!isStructuredDiaryBody(props.body)) {
        return <div className={isCompact ? styles.compactPlainBody : styles.plainBody}>{props.body}</div>
    }

    const document = diaryDocumentFromBody(props.body)
    const blocks = props.maxBlocks ? document.content.slice(0, props.maxBlocks) : document.content

    return (
        <div className={isCompact ? styles.readonlyBodyCompact : styles.readonlyBody}>
            {blocks.map((block, index) => renderBlock(block, String(index), {
                imageFallbacks: props.imageFallbacks,
                isCompact,
                legacyDiaryEntryId: props.legacyDiaryEntryId,
                onImageClick: props.onImageClick
            }))}
        </div>
    )
}
