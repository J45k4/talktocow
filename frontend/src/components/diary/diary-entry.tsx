import { Edit } from "@material-ui/icons"
import Link from "next/link"
import React from "react"
import styles from "./diary-entry.module.css";

export const DiaryEntry = (props: {
    id: number
    title: string
    body: string
    postedAt: string
}) => {
    const d = new Date(props.postedAt)

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
                    <Edit />
                </Link>
            </div>
            <div className={styles.body}>
                {props.body}
            </div>
        </div>
    )
}