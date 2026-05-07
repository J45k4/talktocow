import { FaPlus } from "react-icons/fa"
import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import { postJson } from "../../api-methods"
import { DiaryEntryList } from "./diary-entry-list"
import styles from "./diary.module.css"

const todayAsDateInputValue = () => new Date().toISOString().slice(0, 10)

export const Diary = () => {
    const navigate = useNavigate()
    const [newEntryDate, setNewEntryDate] = useState(todayAsDateInputValue())
    const [newEntryTitle, setNewEntryTitle] = useState("New diary entry")
    const [newEntryBody, setNewEntryBody] = useState("")
    const [isCreatingEntry, setIsCreatingEntry] = useState(false)
    const [isPostingEntry, setIsPostingEntry] = useState(false)

    const resetDraft = () => {
        setNewEntryDate(todayAsDateInputValue())
        setNewEntryTitle("New diary entry")
        setNewEntryBody("")
        setIsCreatingEntry(false)
    }

    const postEntry = () => {
        const createdAt = newEntryDate ? `${newEntryDate}T12:00:00Z` : undefined

        setIsPostingEntry(true)

        postJson<any>("/api/diary/entry", {
            title: newEntryTitle,
            body: newEntryBody,
            createdAt
        }).then(r => {
            if (r.payload) {
                resetDraft()
                navigate("/diary/entry/" + r.payload.id)
            }
        }).finally(() => {
            setIsPostingEntry(false)
        })
    }

    return (
        <div className={styles.container}>
            <div className={styles.header}>
                <div className={styles.addButtonArea}>
                    <button className={styles.newEntryButton} onClick={() => {
                        setIsCreatingEntry(true)
                    }}>
                        <FaPlus />
                        New diary entry
                    </button>
                </div>
            </div>
            {isCreatingEntry && (
                <div className={styles.draftForm}>
                    <label className={styles.draftField}>
                        Entry date
                        <input
                            className={styles.datePicker}
                            type="date"
                            value={newEntryDate}
                            onChange={e => setNewEntryDate(e.target.value)}
                        />
                    </label>
                    <label className={styles.draftField}>
                        Title
                        <input
                            type="text"
                            value={newEntryTitle}
                            onChange={e => setNewEntryTitle(e.target.value)}
                        />
                    </label>
                    <label className={styles.draftField}>
                        Body
                        <textarea
                            value={newEntryBody}
                            onChange={e => setNewEntryBody(e.target.value)}
                        />
                    </label>
                    <div className={styles.draftActions}>
                        <button onClick={postEntry} disabled={isPostingEntry}>
                            {isPostingEntry ? "Posting..." : "Post"}
                        </button>
                        <button onClick={resetDraft} disabled={isPostingEntry}>
                            Cancel
                        </button>
                    </div>
                </div>
            )}
            <DiaryEntryList />
        </div>
    )
}
