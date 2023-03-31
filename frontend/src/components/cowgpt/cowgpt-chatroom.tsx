import React, { useState } from "react"
import { ChatroomMessages } from "../../chatroom/chatroom-messages"
import { ChatroomSendMessage } from "../../chatroom/chatroom-send-message"

export const CowGPTChatroom = (props: {
	chatroomId: string
}) => {
	return (
		<>
			<div style={{
				display: "flex",
				flexGrow: 1,
				padding: "1em",
				overflow: "auto",
			}}>
				<ChatroomMessages chatroomId={props.chatroomId} />
			</div>
			<div style={{
				display: "flex",
			}}>
				<ChatroomSendMessage chatroomId={props.chatroomId} />
			</div>
		</>
	)
}