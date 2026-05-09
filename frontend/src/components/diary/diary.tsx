import { FaPlus } from "react-icons/fa"
import React from "react"
import { useNavigate } from "react-router-dom"
import { DiaryEntryList } from "./diary-entry-list"
import styles from "./diary.module.css"

export const Diary = () => {
    const navigate = useNavigate()

    return (
        <div className={styles.container}>
            <div className={styles.header}>
                <div className={styles.addButtonArea}>
                    <button className={styles.newEntryButton} onClick={() => {
                        navigate("/diary/new")
                    }}>
                        <FaPlus />
                        New diary entry
                    </button>
                </div>
            </div>
            <DiaryEntryList />
        </div>
    )
}
