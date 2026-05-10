import React, { useEffect, useState } from "react"
import { getJson, putJson } from "../../../src/api-methods"

import { useNavigate, useParams } from "react-router-dom"
import { PageContainer } from "../../../src/components/page-container"
import { createDiaryBodyFromPlainTextAndImages, DiaryLexicalEditor, getDiaryBodyFileIds, isStructuredDiaryBody } from "../../../src/components/diary/lexical-diary"
import styles from "../../../src/components/diary/diary.module.css"

type DiaryEntryPicture = {
    id: number
    fileId: number
    fileName: string
    url: string
}

export default function DiaryEntryPage() {
    const navigate = useNavigate()
    const { diaryEntryId } = useParams()

    const [entry, setEntry] = useState<any>()
    const [isSaving, setIsSaving] = useState(false)
    const [saveError, setSaveError] = useState("")

    useEffect(() => {
        Promise.all([
            getJson<any>("/api/diary/entry/" + diaryEntryId),
            getJson<DiaryEntryPicture[]>(`/api/diary/entry/${diaryEntryId}/pictures`)
        ]).then(([entryResponse, picturesResponse]) => {
            const loadedEntry = entryResponse.payload
            const pictures = picturesResponse.payload ?? []

            if (!loadedEntry) {
                setEntry(loadedEntry)
                return
            }

            setEntry({
                ...loadedEntry,
                body: isStructuredDiaryBody(loadedEntry.body)
                    ? loadedEntry.body
                    : createDiaryBodyFromPlainTextAndImages(loadedEntry.body ?? "", pictures.map(picture => ({
                        fileId: picture.fileId,
                        fileName: picture.fileName,
                        url: picture.url
                    })))
            })
        })
    }, [diaryEntryId])

    const updateEntry = () => {
        if (!entry) {
            return
        }

        setIsSaving(true)
        setSaveError("")

        putJson("/api/diary/entry/" + diaryEntryId, {
            title: entry.title,
            body: entry.body,
            label: entry.label || undefined,
            pictureFileIds: getDiaryBodyFileIds(entry.body ?? ""),
            mask: ["title", "body", "label", "pictureFileIds"]
        }).then(r => {
            if (r.error) {
                setSaveError(r.error.message)
                return
            }

            if (r.payload) {
                setEntry(r.payload)
            }
        }).finally(() => {
            setIsSaving(false)
        })
    }

    return (
        <PageContainer>
            <div className={styles.entryPage}>
                <div className={styles.entryEditorCard}>
                    <div className={styles.entryPageHeader}>
                        <button className={styles.secondaryButton} onClick={() => navigate(-1)}>← Back</button>
                        <div>
                            <div className={styles.entryPageEyebrow}>Diary</div>
                            <h1>Edit diary entry</h1>
                        </div>
                    </div>

                    {!entry && (
                        <div className={styles.emptyEditorState}>Loading entry…</div>
                    )}

                    {entry && (
                        <div className={styles.draftForm}>
                            <label className={styles.draftField}>
                                <span>Title</span>
                                <input
                                    className={styles.titleInput}
                                    type="text"
                                    value={entry.title ?? ""}
                                    onChange={e => setEntry({
                                        ...entry,
                                        title: e.target.value
                                    })}
                                />
                            </label>
                            <label className={styles.draftField}>
                                <span>Label</span>
                                <input
                                    type="text"
                                    value={entry.label ?? ""}
                                    placeholder="e.g. Sauna day"
                                    onChange={e => setEntry({
                                        ...entry,
                                        label: e.target.value
                                    })}
                                />
                            </label>
                            <div className={styles.draftField}>
                                <span>Content</span>
                                <DiaryLexicalEditor
                                    value={entry.body ?? ""}
                                    onChange={body => setEntry({
                                        ...entry,
                                        body
                                    })}
                                />
                            </div>
                            {saveError && <div className={styles.saveError}>{saveError}</div>}
                            <div className={styles.draftActions}>
                                <button className={styles.secondaryButton} onClick={() => navigate(-1)} disabled={isSaving}>
                                    Cancel
                                </button>
                                <button className={styles.primaryButton} onClick={updateEntry} disabled={isSaving}>
                                    {isSaving ? "Saving..." : "Save changes"}
                                </button>
                            </div>
                        </div>
                    )}
                </div>
            </div>
        </PageContainer>
    )
}
