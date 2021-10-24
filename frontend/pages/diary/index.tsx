import { DiaryEntryList } from "../../src/components/diary/diary-entry-list"
import React from "react"
import { postJson } from "../../src/utility/talktocow-api-helpers"

export default function DiaryPage() {
    return (
        <div>
            <button onClick={() => {
                postJson("/api/diary/entry", {
                    title: "New diary entry",
                    body: ""
                })
            }}>
                Create new
            </button>
            <DiaryEntryList />
        </div>
    )
}