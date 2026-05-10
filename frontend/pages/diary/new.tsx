import React, { useState } from "react"
import { useNavigate, useSearchParams } from "react-router-dom"
import { deleteJson, postJson } from "../../src/api-methods"
import { PageContainer } from "../../src/components/page-container"
import { createDiaryBodyFromPlainTextAndImages, DiaryLexicalEditor, getDiaryBodyFileIds } from "../../src/components/diary/lexical-diary"
import styles from "../../src/components/diary/diary.module.css"

const todayAsDateInputValue = () => new Date().toISOString().slice(0, 10)

export default function NewDiaryEntryPage() {
    const navigate = useNavigate()
    const [searchParams] = useSearchParams()
    const initialDate = searchParams.get("date") ?? todayAsDateInputValue()
    const [entryDate, setEntryDate] = useState(initialDate)
    const [title, setTitle] = useState("New diary entry")
    const [label, setLabel] = useState("")
    const [body, setBody] = useState(createDiaryBodyFromPlainTextAndImages("", []))
    const [isPostingEntry, setIsPostingEntry] = useState(false)
    const [saveError, setSaveError] = useState("")

    const postEntry = async () => {
        const createdAt = entryDate ? `${entryDate}T12:00:00Z` : undefined

        setIsPostingEntry(true)
        setSaveError("")

        const pictureFileIds: number[] = []

        try {
            pictureFileIds.push(...getDiaryBodyFileIds(body))

            const r = await postJson<any>("/api/diary/entry", {
                title,
                body,
                label: label || undefined,
                createdAt,
                pictureFileIds
            })

            if (r.error) {
                await Promise.all(pictureFileIds.map(fileId => deleteJson(`/api/files/${fileId}`)))
                setSaveError(r.error.message)
                return
            }

            if (r.payload) {
                navigate("/diary")
            }
        } catch (error) {
            await Promise.all(pictureFileIds.map(fileId => deleteJson(`/api/files/${fileId}`)))
            setSaveError(error instanceof Error ? error.message : String(error))
        } finally {
            setIsPostingEntry(false)
        }
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
                            <span>Label</span>
                            <input
                                type="text"
                                value={label}
                                placeholder="e.g. Sauna day"
                                onChange={e => setLabel(e.target.value)}
                            />
                        </label>
                        <div className={styles.draftField}>
                            <span>Content</span>
                            <DiaryLexicalEditor value={body} onChange={setBody} />
                        </div>
                        {saveError && <div className={styles.saveError}>{saveError}</div>}
                        <div className={styles.draftActions}>
                            <button className={styles.secondaryButton} onClick={() => navigate("/diary")} disabled={isPostingEntry}>
                                Cancel
                            </button>
                            <button className={styles.primaryButton} onClick={postEntry} disabled={isPostingEntry}>
                                {isPostingEntry ? "Posting..." : "Post entry"}
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </PageContainer>
    )
}
