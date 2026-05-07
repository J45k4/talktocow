import React, { useCallback, useEffect, useState } from "react"

import { DiaryEntry } from "./diary-entry"
import { getJson } from "../../api-methods"
import { PullToRefresh } from "../pull-to-refresh"

import styles from "./diary-entry-list.module.css"

import InfiniteScroll from "react-infinite-scroll-component"

export const DiaryEntryList = () => {
    const [offset, setOffset] = useState(0)
    const [count, setCount] = useState(0)

    const [entries, setEntries] = useState<any[]>([])

    const fetchFirstPage = useCallback(async () => {
        const [entriesResponse, countResponse] = await Promise.all([
            getJson<any>("/api/diary/entries?offset=0"),
            getJson<any>("/api/diary/entries/count")
        ])

        const newEntries = entriesResponse.payload ?? []

        setEntries(newEntries)
        setOffset(newEntries.length)
        setCount(countResponse.payload?.count ?? 0)
    }, [])

    useEffect(() => {
        fetchFirstPage()
    }, [fetchFirstPage])

    return (
        <PullToRefresh onRefresh={() => window.location.reload()}>
            <div className={styles.container}>
                <InfiniteScroll
                    dataLength={entries.length}
                    next={() => {
                        console.log("Next")

                        getJson<any>("/api/diary/entries?offset=" + offset).then(r => {
                            const newEntries = r.payload ?? []
                            setEntries([...entries, ...newEntries])
                            setOffset(offset + newEntries.length)
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
        </PullToRefresh>
    )
}
