import React from "react"
import { NewChatroomForm } from "../../src/components/chatroom/new_chatroom_form"
import { NavigationBar } from "../../src/components/navigation-bar"
import { PageContainer } from "../../src/components/page-container"

export default function NewChatroomPage() {
    return (
        <PageContainer>
            <NewChatroomForm />
        </PageContainer>
    )
}