import { ArrowBack, Add } from "@material-ui/icons"
import { useRouter } from 'next/router'
import React from "react"
import { postJson } from "../../utility/talktocow-api-helpers"
import { DiaryEntryList } from "./diary-entry-list"
import Link from "next/link"
import styles from "./diary.module.css"

export const Diary = () => {
    const router = useRouter()

    return (
        <div>
            <div className={styles.header}>
                <Link href="/">
                    <ArrowBack style={{
                        fontSize: "40px"
                    }} onClick={() => {
                        console.log("go to frontpage")
                        router.push("/")
                    }} />
                </Link>
                <div className={styles.title}>
                    Diary
                </div>
                <div className={styles.addButtonArea}>
                    <Add style={{
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