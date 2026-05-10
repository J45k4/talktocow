import React, { useState } from "react"
import { useNavigate, useSearchParams } from "react-router-dom"
import { deleteJson, postFormData, postJson } from "../../src/api-methods"
import { PageContainer } from "../../src/components/page-container"
import styles from "../../src/components/diary/diary.module.css"

const todayAsDateInputValue = () => new Date().toISOString().slice(0, 10)

type UploadedFile = {
    id: number
}

export default function NewDiaryEntryPage() {
    const navigate = useNavigate()
    const [searchParams] = useSearchParams()
    const initialDate = searchParams.get("date") ?? todayAsDateInputValue()
    const [entryDate, setEntryDate] = useState(initialDate)
    const [title, setTitle] = useState("New diary entry")
    const [label, setLabel] = useState("")
    const [body, setBody] = useState("")
    const [pictures, setPictures] = useState<File[]>([])
    const [isPostingEntry, setIsPostingEntry] = useState(false)
    const [saveError, setSaveError] = useState("")

    const uploadPictures = async () => {
        const pictureFileIds: number[] = []

        for (const picture of pictures) {
            const formData = new FormData()
            formData.append("file", picture)

            const response = await postFormData<UploadedFile>("/api/files", formData)

            if (response.error) {
                throw new Error(response.error.message)
            }

            if (response.payload) {
                pictureFileIds.push(response.payload.id)
            }
        }

        return pictureFileIds
    }

    const postEntry = async () => {
        const createdAt = entryDate ? `${entryDate}T12:00:00Z` : undefined

        setIsPostingEntry(true)
        setSaveError("")

        const pictureFileIds: number[] = []

        try {
            pictureFileIds.push(...await uploadPictures())

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
                        <label className={styles.draftField}>
                            <span>Content</span>
                            <textarea
                                className={styles.longContentInput}
                                value={body}
                                placeholder="Write as much as you want..."
                                onChange={e => setBody(e.target.value)}
                            />
                        </label>
                        <label className={styles.draftField}>
                            <span>Pictures</span>
                            <input
                                type="file"
                                accept="image/*"
                                multiple
                                onChange={e => setPictures(Array.from(e.target.files ?? []))}
                            />
                        </label>
                        {pictures.length > 0 && (
                            <div className={styles.pictureSelectionList}>
                                {pictures.map(picture => (
                                    <div className={styles.pictureSelectionItem} key={`${picture.name}-${picture.size}`}>
                                        {picture.name}
                                    </div>
                                ))}
                            </div>
                        )}
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
