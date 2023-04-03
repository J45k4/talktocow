import { useRouter } from "next/router"
import React from "react"
import { Chatroom } from "./chatroom"
import { CowGPTChatrooms } from "../components/cowgpt/cowgpt-chatrooms"
import { ChatroomSearchButton } from "./chatroom-search-button"

export const NarrowChatroom = (props: {
	chatroomId?: string
}) => {
	const router = useRouter()

	if (props.chatroomId) {
		return (
			<>
				<div style={{
					display: "flex",
					flexDirection: "row",
					marginLeft: "1em",
					marginRight: "1em",
				}}>
					<div style={{
						flexGrow: 1
					}}>
						<button onClick={() => {
							router.push("/cowgpt")
						}}>
							Hamburger
						</button>
					</div>
					<div style={{
						flexGrow: 1,
						display: "flex",
						justifyContent: "flex-end"
					}}>
						<ChatroomSearchButton />
					</div>
				</div>
				<Chatroom chatroomId={props.chatroomId} />
			</>
		)
	}

	return <CowGPTChatrooms />
}