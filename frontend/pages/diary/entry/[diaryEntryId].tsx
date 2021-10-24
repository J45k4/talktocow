import React, { useEffect, useState } from "react"
import { getJson, postJson, putJson } from "../../../src/utility/talktocow-api-helpers"

import { useRouter } from "next/dist/client/router"

export default function DiaryEntryPage() {
    const router = useRouter()

    console.log("query", router.query)

    const diaryEntryId = router.query.diaryEntryId

    const [entry, setEntry] = useState<any>()

    useEffect(() => {
        getJson("/api/diary/entry/" + diaryEntryId).then(r => {
            setEntry(r.payload)
        })
    }, [diaryEntryId])

    console.log("entry", entry)

    return (
        <div>
            <div>
                <button onClick={() => {
                    router.back()
                }}>
                    back
                </button>
                Title
                <input type="text" value={entry?.title} onChange={e => {
                    setEntry({
                        ...entry,
                        title: e.target.value
                    })
                }} />
            </div>
            <div>
                Body
                <textarea value={entry?.body} onChange={e => {
                    setEntry({
                        ...entry,
                        body: e.target.value
                    })
                }} />
            </div>
            <button onClick={() => {
                putJson("/api/diary/entry/" + diaryEntryId, {
                    title: entry.title,
                    body: entry.body,
                    mask: ["title", "body"]
                })
            }}>
                Update
            </button>
        </div>
    )
}