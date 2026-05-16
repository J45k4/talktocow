import { diaryFileImageUrl, diaryFileUrl, diaryImageSource, diaryImageVariantUrl } from "../src/components/diary/diary-image-source"

it("builds diary file image variants", () => {
    expect(diaryFileUrl(42)).toBe("/api/files/42")
    expect(diaryFileImageUrl(42, "large")).toBe("/api/files/42?size=large")
})

it("does not put auth tokens into diary image URLs", () => {
    expect(diaryImageSource(diaryImageVariantUrl("/api/files/42", "large"))).toBe("http://localhost:12001/api/files/42?size=large")
    expect(diaryImageSource("https://app.example.test/api/files/42?size=medium")).toBe("https://app.example.test/api/files/42?size=medium")
    expect(diaryImageSource(diaryImageVariantUrl("https://api.example.test/api/files/42", "thumb"))).toBe("https://api.example.test/api/files/42?size=thumb")
})

it("preserves file-like browser URLs", () => {
    expect(diaryImageVariantUrl("blob:http://localhost/image-id", "large")).toBe("blob:http://localhost/image-id")
    expect(diaryImageSource("file:///tmp/upload.jpg")).toBe("file:///tmp/upload.jpg")
})
