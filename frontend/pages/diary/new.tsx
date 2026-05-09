import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import { postJson } from "../../src/api-methods"
import { PageContainer } from "../../src/components/page-container"
import styles from "../../src/components/diary/diary.module.css"

const todayAsDateInputValue = () => new Date().toISOString().slice(0, 10)

export default function NewDiaryEntryPage() {
    const navigate = useNavigate()
    const [entryDate, setEntryDate] = useState(todayAsDateInputValue())
    const [title, setTitle] = useState("New diary entry")
    const [body, setBody] = useState("")
    const [isPostingEntry, setIsPostingEntry] = useState(false)

    const postEntry = () => {
        const createdAt = entryDate ? `${entryDate}T12:00:00Z` : undefined

        setIsPostingEntry(true)

        postJson<any>("/api/diary/entry", {
            title,
            body,
            createdAt
        }).then(r => {
            if (r.payload) {
                navigate("/diary/entry/" + r.payload.id)
            }
        }).finally(() => {
            setIsPostingEntry(false)
        })
    }

    return (
        <PageContainer>
            <div className={styles.entryPage}>
                <div className={styles.entryEditorCard}>
                    <div className={styles.entryPageHeader}>
                        <button className={styles.secondaryButton} onClick={() => navigate("/diary")}>← Back</button>
                        <div>
                            <div className={styles.entryPageEyebrow}>Diary</div>
                            <h1>New diary entry</h1>
                        </div>
                    </div>
                    <div className={styles.draftForm}>
                        <label className={styles.draftField}>
                            <span>Entry date</span>
                            <input
                                className={styles.datePicker}
                                type="date"
                                value={entryDate}
                                onChange={e => setEntryDate(e.target.value)}
                            />
                        </label>
                        <label className={styles.draftField}>
                            <span>Title</span>
                            <input
                                className={styles.titleInput}
                                type="text"
                                value={title}
                                onChange={e => setTitle(e.target.value)}
                            />
                        </label>
                        <label className={styles.draftField}>
                            <span>Content</span>
                            <textarea
                                className={styles.longContentInput}
                                value={body}
                                placeholder="Write as much as you want..."
                                onChange={e => setBody(e.target.value)}
                            />
                        </label>
                        <div className={styles.draftActions}>
                            <button className={styles.primaryButton} onClick={postEntry} disabled={isPostingEntry}>
                                {isPostingEntry ? "Posting..." : "Post entry"}
                            </button>
                            <button className={styles.secondaryButton} onClick={() => navigate("/diary")} disabled={isPostingEntry}>
                                Cancel
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </PageContainer>
    )
}
