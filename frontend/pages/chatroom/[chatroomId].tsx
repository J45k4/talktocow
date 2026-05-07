import React from "react"
import { useParams } from "react-router-dom"
import { Chatroom } from "../../src/components/chatroom"
import { PageContainer } from "../../src/components/page-container"

export default function ChatroomPage() {
    const { chatroomId } = useParams()
    
    return (
        <PageContainer>
            {chatroomId &&
                <Chatroom chatroomId={chatroomId} />}
        </PageContainer>
    )
}
