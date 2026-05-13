import { resolveServerUrl } from "../../utility"

export type DiaryImageSize = "thumb" | "medium" | "large" | "original"

const externalSchemePattern = /^[a-z][a-z\d+\-.]*:/i
const passthroughSchemePattern = /^(blob|data|file):/i

const canAppendServerImageParams = (url: string) => !passthroughSchemePattern.test(url)

const appendQueryParam = (url: string, name: string, value?: string) => {
    if (!value || !canAppendServerImageParams(url)) {
        return url
    }

    const hashIndex = url.indexOf("#")
    const beforeHash = hashIndex === -1 ? url : url.slice(0, hashIndex)
    const hash = hashIndex === -1 ? "" : url.slice(hashIndex)
    const separator = beforeHash.includes("?") ? "&" : "?"

    return `${beforeHash}${separator}${encodeURIComponent(name)}=${encodeURIComponent(value)}${hash}`
}

export const diaryFileUrl = (fileId: number) => `/api/files/${fileId}`

export const diaryImageVariantUrl = (url: string, size: DiaryImageSize) => {
    return appendQueryParam(url, "size", size)
}

export const diaryFileImageUrl = (fileId: number, size: DiaryImageSize = "medium") => {
    return diaryImageVariantUrl(diaryFileUrl(fileId), size)
}

export const diaryImageSource = (url: string) => {
    if (externalSchemePattern.test(url)) {
        return url
    }

    return resolveServerUrl(url)
}
