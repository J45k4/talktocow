import { useRouter } from "next/router"
import React from "react"
import { Chatroom } from "../../src/components/chatroom"
import { PageContainer } from "../../src/components/page-container"

export default function ChatroomPage() {
    const router = useRouter()

    const chatroomId = router.query.chatroomId as string
    
    return (
        <PageContainer>
            {chatroomId &&
                <Chatroom chatroomId={chatroomId} />}
        </PageContainer>
    )
}