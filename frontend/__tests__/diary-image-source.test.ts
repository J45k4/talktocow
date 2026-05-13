import { diaryFileImageUrl, diaryFileUrl, diaryImageSource, diaryImageVariantUrl } from "../src/components/diary/diary-image-source"

it("builds diary file image variants", () => {
    expect(diaryFileUrl(42)).toBe("/api/files/42")
    expect(diaryFileImageUrl(42, "large")).toBe("/api/files/42?size=large")
})

it("does not put auth tokens into same-origin diary image URLs", () => {
    expect(diaryImageSource(diaryImageVariantUrl("/api/files/42", "large"), "token value", "https://app.example.test")).toBe("/api/files/42?size=large")
    expect(diaryImageSource("https://app.example.test/api/files/42?size=medium", "token value", "https://app.example.test")).toBe("https://app.example.test/api/files/42?size=medium")
})

it("adds the token to cross-origin diary image URLs", () => {
    expect(diaryImageSource(diaryImageVariantUrl("https://api.example.test/api/files/42", "thumb"), "token value", "https://app.example.test")).toBe("https://api.example.test/api/files/42?size=thumb&token=token%20value")
})

it("preserves file-like browser URLs instead of appending server-only params", () => {
    expect(diaryImageVariantUrl("blob:http://localhost/image-id", "large")).toBe("blob:http://localhost/image-id")
    expect(diaryImageSource("file:///tmp/upload.jpg", "token value", "https://app.example.test")).toBe("file:///tmp/upload.jpg")
})
