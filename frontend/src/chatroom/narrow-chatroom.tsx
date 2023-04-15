import React from "react"
import { Chatroom } from "./chatroom"
import { Chats } from "./chats"
import Link from "next/link"
import { Button } from "../components"

export const NarrowChatroom = (props: {
	chatroomId?: string
}) => {
	if (props.chatroomId) {
		return <Chatroom chatroomId={props.chatroomId} />
	}

	return (
		<div>
			<Link href="/chats/new">
				<Button 
					title="New chat"
				/>
			</Link>
			<Chats />
		</div>
	)
}