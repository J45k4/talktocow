import Link from "next/link"
import React from "react"
import { Chatroom } from "../src/components/chatroom"
import { ChatroomList } from "../src/components/chatroom/chatroom_list"
import { NavigationBar } from "../src/components/navigation_bar"
import { PageContainer } from "../src/components/page_container"

export default function ChatroomsPage() {
    return (
        <PageContainer>
            <div style={{
                flexGrow: 1
            }}>
                <Link href={"/chatroom/new"} style={{
					marginBottom: "10px"
				}}>
                    <button>
                        New
                    </button>
                </Link>
                
                <ChatroomList />
            </div>
        </PageContainer>

    )
}