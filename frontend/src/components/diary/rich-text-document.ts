export const DIARY_RICH_TEXT_DOCUMENT_VERSION = 1

export type DiaryRichTextDocumentVersion = typeof DIARY_RICH_TEXT_DOCUMENT_VERSION

export type DiaryRichTextMark = "bold" | "italic" | "underline"

export type DiaryRichTextTextNode = {
    text: string
    marks?: DiaryRichTextMark[]
}

export type DiaryRichTextLineBreakNode = {
    type: "lineBreak"
}

export type DiaryRichTextInlineNode = DiaryRichTextTextNode | DiaryRichTextLineBreakNode

export type DiaryRichTextParagraphBlock = {
    type: "paragraph"
    children: DiaryRichTextInlineNode[]
}

export type DiaryRichTextHeadingLevel = 2 | 3

export type DiaryRichTextHeadingBlock = {
    type: "heading"
    level: DiaryRichTextHeadingLevel
    children: DiaryRichTextInlineNode[]
}

export type DiaryRichTextImageBlock = {
    type: "image"
    fileId: number
    alt?: string
}

export type DiaryRichTextBlock =
    | DiaryRichTextParagraphBlock
    | DiaryRichTextHeadingBlock
    | DiaryRichTextImageBlock

export type DiaryRichTextDocument = {
    version: DiaryRichTextDocumentVersion
    content: DiaryRichTextBlock[]
}

export const createEmptyDiaryRichTextDocument = (): DiaryRichTextDocument => ({
    version: DIARY_RICH_TEXT_DOCUMENT_VERSION,
    content: [{
        type: "paragraph",
        children: []
    }]
})
