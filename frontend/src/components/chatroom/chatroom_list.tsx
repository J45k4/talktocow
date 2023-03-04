import Link from "next/link"
import React, { useEffect, useState } from "react"
import { getJson } from "../../utility/talktocow-api-helpers"

export const ChatroomList = () => {
    const [chatrooms, setChatrooms] = useState<any[]>([])

    useEffect(() => {
        getJson<any>("/api/chatrooms").then(r => {
            setChatrooms(r.payload)
        })
    }, [])

    if (chatrooms.length === 0) {
        return (
            <div>
                No chatrooms
            </div>
        )
    }

    return (
        <div>
            {chatrooms.map(p => (
                <Link href={`/chatroom/${p.id}`} key={p.id}>
                    <div key={p.id} style={{ cursor: "pointer" }}>
                        {"chatroom " + p.id}
                    </div>
                </Link>
            ))}
        </div>
    )
}