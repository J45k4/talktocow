import React, { useEffect, useState } from "react"

import { DiaryEntry } from "./diary-entry"
import { getJson } from "../../utility/talktocow-api-helpers"

export const DiaryEntryList = (props: {

}) => {
    const [entries, setEntries] = useState<any[]>([])

    useEffect(() => {
        getJson<any>("/api/diary/entries").then(r => {
            console.log("Response", r)
            setEntries(r.payload)
        })
    }, [])

    return (
        <div>
            {entries.map(p => (
                <DiaryEntry
                    key={p.id}
                    id={p.id}
                    title={p.title}
                    body={p.body}
                    postedAt={p.createdAt}

                />
            ))}
        </div>
    )
}