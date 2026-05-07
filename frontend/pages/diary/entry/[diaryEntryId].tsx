import React, { useEffect, useState } from "react"
import { getJson, putJson } from "../../../src/api-methods"

import { useNavigate, useParams } from "react-router-dom"

export default function DiaryEntryPage() {
    const navigate = useNavigate()
    const { diaryEntryId } = useParams()

    const [entry, setEntry] = useState<any>()

    useEffect(() => {
        getJson("/api/diary/entry/" + diaryEntryId).then(r => {
            setEntry(r.payload)
        })
    }, [diaryEntryId])

    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
            padding: "10px",
            height: "90%"
        }}>
            <div style={{

            }}>
                <button onClick={() => {
                    navigate(-1)
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
            <div style={{
                flexGrow: 1,
                display: "flex",
                flexDirection: "column",
            }}>
                <textarea value={entry?.body} style={{
                    flexGrow: 1
                }} onChange={e => {
                    setEntry({
                        ...entry,
                        body: e.target.value
                    })
                }} />                
            </div>
            <div style={{

            }}>
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
        </div>
    )
}
