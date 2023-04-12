import React from "react"
import { Chatroom } from "./chatroom"
import { Chats } from "../chatroom/chatrooms"

export const NarrowChatroom = (props: {
	chatroomId?: string
}) => {
	if (props.chatroomId) {
		return <Chatroom chatroomId={props.chatroomId} />
	}

	return <Chats />
}