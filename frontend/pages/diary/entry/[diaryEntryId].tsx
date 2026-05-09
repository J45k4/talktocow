import React, { useEffect, useState } from "react"
import { deleteJson, getJson, postFormData, putJson } from "../../../src/api-methods"

import { useNavigate, useParams } from "react-router-dom"
import { PageContainer } from "../../../src/components/page-container"
import { getSession } from "../../../src/logic/session-manager"
import { resolveServerUrl } from "../../../src/utility"
import styles from "../../../src/components/diary/diary.module.css"

type DiaryEntryPicture = {
    id: number
    fileName: string
    url: string
}

const pictureSource = (url: string) => {
    const token = getSession().token
    const separator = url.includes("?") ? "&" : "?"
    return resolveServerUrl(token ? `${url}${separator}token=${encodeURIComponent(token)}` : url)
}

export default function DiaryEntryPage() {
    const navigate = useNavigate()
    const { diaryEntryId } = useParams()

    const [entry, setEntry] = useState<any>()
    const [pictures, setPictures] = useState<DiaryEntryPicture[]>([])
    const [isSaving, setIsSaving] = useState(false)
    const [isUploadingPicture, setIsUploadingPicture] = useState(false)
    const [saveError, setSaveError] = useState("")

    const fetchPictures = () => {
        getJson<DiaryEntryPicture[]>(`/api/diary/entry/${diaryEntryId}/pictures`).then(r => {
            setPictures(r.payload ?? [])
        })
    }

    useEffect(() => {
        getJson("/api/diary/entry/" + diaryEntryId).then(r => {
            setEntry(r.payload)
        })
        fetchPictures()
    }, [diaryEntryId])

    const uploadPictures = (files: FileList | null) => {
        if (!files || files.length === 0) {
            return
        }

        setIsUploadingPicture(true)
        setSaveError("")

        Promise.all(Array.from(files).map(file => {
            const formData = new FormData()
            formData.append("picture", file)
            return postFormData(`/api/diary/entry/${diaryEntryId}/picture`, formData)
        })).then(results => {
            const error = results.find(result => result.error)?.error

            if (error) {
                setSaveError(error.message)
                return
            }

            fetchPictures()
        }).finally(() => {
            setIsUploadingPicture(false)
        })
    }

    const deletePicture = (pictureId: number) => {
        deleteJson(`/api/diary/entry/${diaryEntryId}/picture/${pictureId}`).then(r => {
            if (r.error) {
                setSaveError(r.error.message)
                return
            }

            setPictures(currentPictures => currentPictures.filter(picture => picture.id !== pictureId))
        })
    }

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
            mask: ["title", "body", "label"]
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
                            <label className={styles.draftField}>
                                <span>Content</span>
                                <textarea
                                    className={styles.longContentInput}
                                    value={entry.body ?? ""}
                                    placeholder="Write as much as you want..."
                                    onChange={e => setEntry({
                                        ...entry,
                                        body: e.target.value
                                    })}
                                />
                            </label>
                            <div className={styles.draftField}>
                                <span>Pictures</span>
                                {pictures.length > 0 && (
                                    <div className={styles.pictureGrid}>
                                        {pictures.map(picture => (
                                            <div className={styles.pictureTile} key={picture.id}>
                                                <img src={pictureSource(picture.url)} alt={picture.fileName} />
                                                <button className={styles.removePictureButton} onClick={() => deletePicture(picture.id)} type="button">
                                                    Remove
                                                </button>
                                            </div>
                                        ))}
                                    </div>
                                )}
                                <input
                                    type="file"
                                    accept="image/*"
                                    multiple
                                    disabled={isUploadingPicture}
                                    onChange={e => uploadPictures(e.target.files)}
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
