import React, { useEffect, useState } from "react"

import { DiaryEntry } from "./diary-entry"
import { getJson } from "../../utility/talktocow-api-helpers"

import styles from "./diary-entry-list.module.css"

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
        <div className={styles.container}>
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