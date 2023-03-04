import React, { useEffect, useState } from "react"
import { getJson, postJson } from "../../utility/talktocow-api-helpers"

export const NewChatroomForm = () => {
    const [users, setUsers] = useState<any[]>([])
    const [selected, setSelected] = useState<string[]>([])
    
    useEffect(() => {
        getJson<any>("/api/users").then(r => {
            setUsers(r.payload)
        })
    }, [])

    return (
        <div>
            <input type="text" />
            <div>
                {users.map(p => (
                    <div onClick={() => {
                        setSelected([...selected, p.id])
                    }} style={{
                        cursor: "pointer",
                        border: selected.includes(p.id) ? "1px solid black" : "none"
                    }} key={p.id}>
                        {p.name}
                    </div>
                ))}
            </div>
            <button onClick={() => {
                postJson<any>("/api/chatroom", {
                    userIds: selected
                }).then(r => {
                    if (r.payload) {
                        window.location.href = "/chatroom/" + r.payload.id
                    }
                })
            }}>
                Create
            </button>
        </div>
    )
}