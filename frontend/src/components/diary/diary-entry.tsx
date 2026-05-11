import { FaEdit, FaTrash } from "react-icons/fa";
import React, { useCallback, useEffect } from "react"
import { Link } from "react-router-dom"
import { deleteJson, getJson, postJson } from "../../api-methods";
import { getSession } from "../../logic/session-manager";
import { resolveServerUrl } from "../../utility";
import { Modal } from "../modal";
import styles from "./diary-entry.module.css";

type DiaryEntryPicture = {
    id: number
    fileName: string
    url: string
}

const pictureSource = (url: string) => {
    return resolveServerUrl(url)
}

export const DiaryEntry = (props: {
    id: number
    title: string
    body: string
    postedAt: string
    postedByUserId: string
    label?: string
    onDelete: () => void
}) => {
    const d = new Date(props.postedAt)
    const day = d.toLocaleDateString("en-US", { day: "2-digit" })
    const month = d.toLocaleDateString("en-US", { month: "short" })
    const year = d.getFullYear()
    const weekday = d.toLocaleDateString("en-US", { weekday: "long" })
    const fullDate = d.toLocaleDateString("en-US", {
        weekday: "long",
        day: "numeric",
        month: "long",
        year: "numeric"
    })
    const isAuthor = getSession().userId === props.postedByUserId

    const [newComment, setNewComment] = React.useState("")

    const [comments, setComments] = React.useState<any[]>([])
    const [pictures, setPictures] = React.useState<DiaryEntryPicture[]>([])
    const [previewPicture, setPreviewPicture] = React.useState<DiaryEntryPicture | null>(null)

    const [commentOffset, setCommentOffset] = React.useState(0)
    const [isConfirmingDelete, setIsConfirmingDelete] = React.useState(false)
    const [deleteError, setDeleteError] = React.useState("")

    const fetchComments = useCallback(() => {
        getJson<any>(`/api/diary/entry/${props.id}/comments?limit=50&offset=${commentOffset}`).then(data => {
            const newComments = [...comments, ...data.payload]
            setComments(newComments)
            setCommentOffset(newComments.length)
        })
    }, [commentOffset, comments, setComments])

    useEffect(() => {
        fetchComments()

        getJson<DiaryEntryPicture[]>(`/api/diary/entry/${props.id}/pictures`).then(data => {
            setPictures(data.payload ?? [])
        })

        getJson<any>(`/api/diary/entry/${props.id}/comments/count`).then(data => {
            setCommentOffset(data.payload.count)
        })
    }, [])

    const postComment = useCallback(() => {
        postJson(`/api/diary/entry/${props.id}/comment`, {
            diaryEntryId: props.id,
            commentText: newComment
        }).then(() => {
            setNewComment("")
            fetchComments()
        })
    }, [newComment, props.id, fetchComments])

    const deleteEntry = useCallback(() => {
        deleteJson(`/api/diary/entry/${props.id}`).then(r => {
            if (r.error && r.error.code !== 404 && r.error.message.toLowerCase() !== "not found") {
                setDeleteError(r.error.message)
                return
            }

            setIsConfirmingDelete(false)
            props.onDelete()
        })
    }, [props])

    return (
        <div className={styles.entry}>
            <div className={styles.header}>
                <time className={styles.dateBadge} dateTime={props.postedAt} aria-label={fullDate}>
                    <span className={styles.weekday}>{weekday}</span>
                    <span className={styles.dateMain}>
                        <span className={styles.day}>{day}</span>
                        <span className={styles.month}>{month}</span>
                    </span>
                    <span className={styles.year}>{year}</span>
                </time>
                <div className={styles.titleRow}>
                    <div>
                        {props.label && <div className={styles.labelChip}>{props.label}</div>}
                        <div className={styles.title}>
                            {props.title}
                        </div>
                    </div>
                    {isAuthor && (
                        <div className={styles.actions}>
                            <Link className={styles.iconButton} to={"/diary/entry/" + props.id} aria-label="Edit diary entry">
                                <FaEdit />
                            </Link>
                            <button className={styles.iconButton} onClick={() => setIsConfirmingDelete(true)} aria-label="Delete diary entry">
                                <FaTrash />
                            </button>
                        </div>
                    )}
                </div>
            </div>
            <div className={styles.body}>
                {props.body}
            </div>
            {pictures.length > 0 && (
                <div className={styles.pictureGrid}>
                    {pictures.map(picture => (
                        <button className={styles.pictureButton} onClick={() => setPreviewPicture(picture)} key={picture.id} type="button">
                            <img src={pictureSource(picture.url)} alt={picture.fileName} />
                        </button>
                    ))}
                </div>
            )}
            <div>
                <div>
                    {comments.map(p => (
                        <div key={p.id} className={styles.comment}>
                            <div className={styles.commentAuthor}>
                                {p.userName}
                            </div>
                            <div className={styles.commentBody}>
                                {p.commentText}
                            </div>                            
                        </div>
                    ))}
                </div>
                <div>
                    <input type="text" value={newComment} onChange={e => {
                        setNewComment(e.target.value)
                    }} onKeyDown={(e) => {
                        if (e.key === "Enter") {
                            postComment()
                        }
                    }} />
                    <button onClick={() => {
                        postComment()
                    }}>
                        Comment
                    </button>
                </div>
            </div>
            <Modal isOpen={isConfirmingDelete} title="Delete diary entry?" onClose={() => setIsConfirmingDelete(false)}>
                <div className={styles.modalContent}>
                    <p>Delete “{props.title}”?</p>
                    <div className={styles.modalActions}>
                        <button className={styles.cancelButton} onClick={() => setIsConfirmingDelete(false)}>Cancel</button>
                        <button className={styles.dangerButton} onClick={deleteEntry}>Delete</button>
                    </div>
                </div>
            </Modal>
            <Modal isOpen={deleteError !== ""} title="Could not delete entry" onClose={() => setDeleteError("")}>
                <div className={styles.modalContent}>
                    <p>{deleteError}</p>
                    <div className={styles.modalActions}>
                        <button className={styles.cancelButton} onClick={() => setDeleteError("")}>OK</button>
                    </div>
                </div>
            </Modal>
            {previewPicture && (
                <div className={styles.fullscreenPreview} onClick={() => setPreviewPicture(null)}>
                    <button className={styles.closePreviewButton} onClick={() => setPreviewPicture(null)} type="button">×</button>
                    <img src={pictureSource(previewPicture.url)} alt={previewPicture.fileName} onClick={event => event.stopPropagation()} />
                </div>
            )}
        </div>
    )
}
