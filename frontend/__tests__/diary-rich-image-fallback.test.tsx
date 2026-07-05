import { fireEvent, render, screen, waitFor } from "@testing-library/react"
import { MemoryRouter } from "react-router-dom"
import { beforeEach, expect, it, vi } from "vitest"
import { DiaryEntry } from "../src/components/diary/diary-entry"
import { DIARY_RICH_TEXT_DOCUMENT_VERSION } from "../src/components/diary/rich-text-document"
import { setSession } from "../src/logic/session-manager"

vi.mock("../src/api-methods", () => ({
    deleteJson: vi.fn(async () => ({})),
    getJson: vi.fn(async (path: string) => {
        if (path.includes("/pictures")) {
            return {
                payload: [{
                    id: 29,
                    fileId: 123,
                    fileName: "IMG_4719.jpg",
                    url: "/api/files/123"
                }]
            }
        }

        if (path.includes("/comments/count")) {
            return { payload: { count: 0 } }
        }

        if (path.includes("/comments")) {
            return { payload: [] }
        }

        return { payload: null }
    }),
    postJson: vi.fn(async () => ({}))
}))

beforeEach(() => {
    setSession({ username: "teemu", userId: "19" })
})

it("falls back to attached diary picture route when a rich body image points at a missing file id", async () => {
    const body = JSON.stringify({
        version: DIARY_RICH_TEXT_DOCUMENT_VERSION,
        content: [{
            type: "image",
            fileId: 999,
            alt: "IMG_4719.jpg"
        }]
    })

    render(
        <MemoryRouter>
            <DiaryEntry
                id={77}
                title="Bicycle day"
                body={body}
                postedAt="2026-05-22T12:00:00Z"
                postedByUserId="19"
                postedByUserName="teemu"
                label="Bicycle day"
                onDelete={() => undefined}
            />
        </MemoryRouter>
    )

    const image = await screen.findByAltText("IMG_4719.jpg") as HTMLImageElement
    expect(image.getAttribute("src")).toBe("http://localhost:12001/api/files/999?size=medium")

    await waitFor(() => {
        expect(screen.getByText("By teemu")).toBeTruthy()
    })

    fireEvent.error(image)

    await waitFor(() => {
        expect(image.getAttribute("src")).toBe("http://localhost:12001/api/diary/entry/77/picture/29?size=medium")
    })
})
