import { FaPlus } from "react-icons/fa"
import { useRouter } from 'next/router'
import React from "react"
import { postJson } from "../../api-methods"
import { DiaryEntryList } from "./diary-entry-list"
import styles from "./diary.module.css"

export const Diary = () => {
    const router = useRouter()

    return (
        <div className={styles.container}>
            <div className={styles.header}>
                <div className={styles.addButtonArea}>
                    <FaPlus style={{
                        fontSize: "40px"
                    }} onClick={() => {
                        postJson<any>("/api/diary/entry", {
                            title: "New diary entry",
                            body: ""
                        }).then(r => {
                            if (r.payload) {
                                router.push("/diary/entry/" + r.payload.id)
                            }
                        })
                    }} />
                </div>
                
            </div>
            <DiaryEntryList />
        </div>
    )
}