import React from "react"
import { Link } from "react-router-dom"
import { Chatroom } from "../src/components/chatroom"
import { ChatroomList } from "../src/components/chatroom/chatroom_list"
import { NavigationBar } from "../src/components/navigation-bar"
import { PageContainer } from "../src/components/page-container"

export default function ChatroomsPage() {
    return (
        <PageContainer>
            <div style={{
                flexGrow: 1
            }}>
                <Link to={"/chatroom/new"} style={{
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
