import { render, screen } from "@testing-library/react"
import { expect, it } from "vitest"
import { DiaryBodyRenderer, createDiaryBodyFromPlainTextImagesAndAudio, getDiaryBodyFileIds, getDiaryBodyRecordingFileIds, hasDiaryBodyContent } from "../src/components/diary/lexical-diary"

it("renders diary audio blocks and extracts recording ids separately from picture ids", () => {
    const body = createDiaryBodyFromPlainTextImagesAndAudio("", [{
        fileId: 12,
        fileName: "picture.jpg",
        url: "/api/files/12"
    }], [{
        durationMs: 65000,
        fileId: 44,
        fileName: "diary-recording.webm"
    }])

    expect(getDiaryBodyFileIds(body)).toEqual([12])
    expect(getDiaryBodyRecordingFileIds(body)).toEqual([44])
    expect(hasDiaryBodyContent(body)).toBe(true)

    render(<DiaryBodyRenderer body={body} />)

    expect(screen.getByText("diary-recording.webm")).toBeTruthy()
    expect(screen.getByText("1:05")).toBeTruthy()
    expect(document.querySelector("audio")?.getAttribute("src")).toBe("http://localhost:12001/api/files/44")
})
