import React from "react"
import { Chatroom } from "../src/components/chatroom"
import { NavigationBar } from "../src/components/navigation_bar"

export default function ChatroomsPage() {
    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
            height: "100%"
        }}>
            <NavigationBar />
            <div style={{
                flexGrow: 1
            }}>
                <Chatroom chatroomId="1"/>
            </div>
        </div>
    )
}