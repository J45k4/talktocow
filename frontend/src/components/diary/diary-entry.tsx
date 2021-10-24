import { Edit } from "@material-ui/icons"
import Link from "next/link"
import React from "react"

export const DiaryEntry = (props: {
    id: number
    title: string
    body: string
    postedAt: string
}) => {
    return (
        <div>
            <div>
                {props.title}
                <Link href={"/diary/entry/" + props.id}>
                    <Edit />
                </Link>
            </div>
            <div>
                {props.body}
            </div>
        </div>
    )
}