import { diaryFileImageUrl, diaryFileUrl, diaryImageSource, diaryImageVariantUrl } from "../src/components/diary/diary-image-source"

it("builds diary file image variants", () => {
    expect(diaryFileUrl(42)).toBe("/api/files/42")
    expect(diaryFileImageUrl(42, "large")).toBe("/api/files/42?size=large")
})

it("does not put auth tokens into diary image URLs", () => {
    expect(diaryImageSource(diaryImageVariantUrl("/api/files/42", "large"))).toBe("/api/files/42?size=large")
})

it("preserves file-like browser URLs instead of appending server-only params", () => {
    expect(diaryImageVariantUrl("blob:http://localhost/image-id", "large")).toBe("blob:http://localhost/image-id")
    expect(diaryImageSource("file:///tmp/upload.jpg")).toBe("file:///tmp/upload.jpg")
})

it("keeps absolute API file URLs absolute without adding auth params", () => {
    expect(diaryImageSource(diaryImageVariantUrl("https://api.example.test/api/files/42", "thumb"))).toBe("https://api.example.test/api/files/42?size=thumb")
})
