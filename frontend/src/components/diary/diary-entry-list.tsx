import React, { useEffect, useState } from "react"

import { DiaryEntry } from "./diary-entry"
import { getJson } from "../../utility/talktocow-api-helpers"

import styles from "./diary-entry-list.module.css"

import InfiniteScroll from "react-infinite-scroll-component"

export const DiaryEntryList = (props: {

}) => {
    const [offset, setOffset] = useState(0)
    const [count, setCount] = useState(0)

    const [entries, setEntries] = useState<any[]>([])

    useEffect(() => {
        getJson<any>("/api/diary/entries?offset=" + offset).then(r => {
            setEntries(r.payload)
            setOffset(offset + r.payload.length)
        })

        getJson<any>("/api/diary/entries/count").then(r => {
            setCount(r.payload.count)
        })
    }, [])

    return (
        <div className={styles.container}>
            <InfiniteScroll
                dataLength={entries.length}
                next={() => {
                    console.log("Next")

                    getJson<any>("/api/diary/entries?offset=" + offset).then(r => {
                        setEntries([...entries, ...r.payload])
                        setOffset(offset + r.payload.length)
                    })
                }}
                hasMore={offset < count}
                loader={<h4>loading</h4>}>
                {entries.map(p => (
                    <DiaryEntry
                        key={p.id}
                        id={p.id}
                        title={p.title}
                        body={p.body}
                        postedAt={p.createdAt}

                    />
                ))}
            </InfiniteScroll>
        </div>
    )
}