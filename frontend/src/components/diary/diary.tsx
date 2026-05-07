import { FaPlus } from "react-icons/fa"
import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import { postJson } from "../../api-methods"
import { DiaryEntryList } from "./diary-entry-list"
import styles from "./diary.module.css"

const todayAsDateInputValue = () => new Date().toISOString().slice(0, 10)

export const Diary = () => {
    const navigate = useNavigate()
    const [newEntryDate, setNewEntryDate] = useState(todayAsDateInputValue())

    return (
        <div className={styles.container}>
            <div className={styles.header}>
                <div className={styles.addButtonArea}>
                    <label className={styles.datePickerLabel}>
                        Entry date
                        <input
                            className={styles.datePicker}
                            type="date"
                            value={newEntryDate}
                            onChange={e => setNewEntryDate(e.target.value)}
                        />
                    </label>
                    <FaPlus style={{
                        fontSize: "40px"
                    }} onClick={() => {
                        const createdAt = newEntryDate ? `${newEntryDate}T12:00:00Z` : undefined

                        postJson<any>("/api/diary/entry", {
                            title: "New diary entry",
                            body: "",
                            createdAt
                        }).then(r => {
                            if (r.payload) {
                                navigate("/diary/entry/" + r.payload.id)
                            }
                        })
                    }} />
                </div>
                
            </div>
            <DiaryEntryList />
        </div>
    )
}
