import { FaEdit, FaTrash } from "react-icons/fa";
import React, { useCallback, useEffect } from "react"
import { Link } from "react-router-dom"
import { deleteJson, getJson, postJson } from "../../api-methods";
import { getSession } from "../../logic/session-manager";
import styles from "./diary-entry.module.css";

export const DiaryEntry = (props: {
    id: number
    title: string
    body: string
    postedAt: string
    postedByUserId: string
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

    const [commentOffset, setCommentOffset] = React.useState(0)

    const fetchComments = useCallback(() => {
        getJson<any>(`/api/diary/entry/${props.id}/comments?limit=50&offset=${commentOffset}`).then(data => {
            const newComments = [...comments, ...data.payload]
            setComments(newComments)
            setCommentOffset(newComments.length)
        })
    }, [commentOffset, comments, setComments])

    useEffect(() => {
        fetchComments()

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
        if (!window.confirm("Delete this diary entry?")) {
            return
        }

        deleteJson(`/api/diary/entry/${props.id}`).then(r => {
            if (!r.error) {
                props.onDelete()
            }
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
                    <div className={styles.title}>
                        {props.title}
                    </div>
                    {isAuthor && (
                        <div className={styles.actions}>
                            <Link className={styles.iconButton} to={"/diary/entry/" + props.id} aria-label="Edit diary entry">
                                <FaEdit />
                            </Link>
                            <button className={styles.iconButton} onClick={deleteEntry} aria-label="Delete diary entry">
                                <FaTrash />
                            </button>
                        </div>
                    )}
                </div>
            </div>
            <div className={styles.body}>
                {props.body}
            </div>
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
        </div>
    )
}
