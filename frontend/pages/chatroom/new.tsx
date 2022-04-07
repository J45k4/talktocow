import React from "react"
import { NewChatroomForm } from "../../src/components/chatroom/new_chatroom_form"
import { NavigationBar } from "../../src/components/navigation_bar"
import { PageContainer } from "../../src/components/page_container"

export default function NewChatroomPage() {
    return (
        <PageContainer>
            <NavigationBar />
            <NewChatroomForm />
        </PageContainer>
    )
}