import { FaEdit } from "react-icons/fa";
import Link from "next/link"
import React, { useCallback, useEffect } from "react"
import { getJson, postJson } from "../../utility/talktocow-api-helpers";
import styles from "./diary-entry.module.css";

export const DiaryEntry = (props: {
    id: number
    title: string
    body: string
    postedAt: string
}) => {
    const d = new Date(props.postedAt)

    const [newComment, setNewComment] = React.useState("")

    const [comments, setComments] = React.useState([])

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

    return (
        <div className={styles.entry}>
            <div className={styles.date}>
                {d.toLocaleDateString("us", { weekday: "long" }) + " "}
                {d.getDate() + " "}
                {d.toLocaleDateString("us", { month: "long" }) + " "}
                {d.getFullYear()}
            </div>
            <div className={styles.title}>
                {props.title}
                <Link href={"/diary/entry/" + props.id}>
                    <FaEdit />
                </Link>
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