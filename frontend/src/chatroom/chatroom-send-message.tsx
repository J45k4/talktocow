import React, { useState, useCallback } from "react"
import { sendMessageToChatroom } from "../chatroom-message-managers"
import { Button } from "../components/button"

export const ChatroomSendMessage = (props: {
	chatroomId: string
}) => {
	const [newMessage, setNewMessage] = useState("")

	const sendMessage = useCallback(() => {
		if (newMessage === "") {
			return
		}

		sendMessageToChatroom(
			props.chatroomId,
			newMessage
		)

		setNewMessage("")
	}, [newMessage])

	return (
		<div style={{
			border: "solid 1px #8E8E8E",
			flexGrow: 1,
			margin: "1em",
			fontSize: "1.5em",
			padding: "0.2em",
			display: "flex",
			flexDirection: "row",
		}}>
			<input style={{
				border: "none",
				width: "100%",
				height: "100%",
				outline: "none",
			}} onKeyDown={e => {
				if (e.key === "Enter") {
					console.log("Enter")

					sendMessage()
				}
			}}
				value={newMessage}
				onChange={e => {
					setNewMessage(e.target.value)
				}}
			/>
			<Button onClick={() => {
				sendMessage()
			}}>
				Send
			</Button>
		</div>
	)
}