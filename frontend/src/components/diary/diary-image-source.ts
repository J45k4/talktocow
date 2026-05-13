import { getSession } from "../../logic/session-manager"
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

const isCrossOriginUrl = (url: string, currentOrigin = globalThis.location?.origin) => {
    if (!currentOrigin || !canAppendServerImageParams(url)) {
        return false
    }

    try {
        return new URL(url, currentOrigin).origin !== currentOrigin
    } catch (_error) {
        return false
    }
}

export const diaryFileUrl = (fileId: number) => `/api/files/${fileId}`

export const diaryImageVariantUrl = (url: string, size: DiaryImageSize) => {
    return appendQueryParam(url, "size", size)
}

export const diaryFileImageUrl = (fileId: number, size: DiaryImageSize = "medium") => {
    return diaryImageVariantUrl(diaryFileUrl(fileId), size)
}

export const diaryImageSource = (url: string, token = getSession().token, currentOrigin?: string) => {
    const resolvedUrl = externalSchemePattern.test(url) ? url : resolveServerUrl(url)

    if (!isCrossOriginUrl(resolvedUrl, currentOrigin)) {
        return resolvedUrl
    }

    return appendQueryParam(resolvedUrl, "token", token)
}
