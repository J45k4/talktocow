import { DiaryEntryList } from "../../src/components/diary/diary-entry-list"
import React from "react"
import { postJson } from "../../src/utility/talktocow-api-helpers"
import { useRouter } from "next/dist/client/router"

export default function DiaryPage() {
    const router = useRouter()

    return (
        <div>
            <button onClick={() => {
                postJson<any>("/api/diary/entry", {
                    title: "New diary entry",
                    body: ""
                }).then(r => {
                    if (r.payload) {
                        router.push("/diary/entry/" + r.payload.id)
                    }
                })
            }}>
                Create new
            </button>
            <DiaryEntryList />
        </div>
    )
}