import { render, screen } from "@testing-library/react"
import { expect, it } from "vitest"
import { DiaryBodyRenderer, createDiaryBodyFromPlainTextImagesAudioAndVideo, getDiaryBodyFileIds, getDiaryBodyRecordingFileIds, getDiaryBodyVideoFileIds, hasDiaryBodyContent } from "../src/components/diary/lexical-diary"

it("renders diary video blocks and extracts video ids separately from image and recording ids", () => {
    const body = createDiaryBodyFromPlainTextImagesAudioAndVideo("", [{
        fileId: 12,
        fileName: "picture.jpg",
        url: "/api/files/12"
    }], [{
        durationMs: 65000,
        fileId: 44,
        fileName: "diary-recording.webm"
    }], [{
        fileId: 55,
        fileName: "clip.mp4"
    }])

    expect(getDiaryBodyFileIds(body)).toEqual([12])
    expect(getDiaryBodyRecordingFileIds(body)).toEqual([44])
    expect(getDiaryBodyVideoFileIds(body)).toEqual([55])
    expect(hasDiaryBodyContent(body)).toBe(true)

    render(<DiaryBodyRenderer body={body} />)

    expect(screen.getByText("clip.mp4")).toBeTruthy()
    expect(document.querySelector("video")?.getAttribute("src")).toBe("http://localhost:12001/api/files/55")
})
