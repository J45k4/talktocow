import React, { useState } from "react"
import { sendMessage } from "../../ws"

export const CowGPTChatroom = (props: {
	chatroomId: string
}) => {
	const [newMessage, setNewMessage] = useState("")

	return (
		<>
			<div style={{
				display: "flex",
				flexGrow: 1,
				padding: "1em",
			}}>
				<div>

				</div>
			</div>
			<div style={{
				display: "flex",

			}}>
				<div style={{
					border: "solid 1px #8E8E8E",
					flexGrow: 1,
					margin: "2em",
					fontSize: "1.5em",
					padding: "0.2em",
				}}>
					<input style={{
						border: "none",
						width: "100%",
						height: "100%",
					}} onKeyDown={e => {
						if (e.key === "Enter") {
							console.log("Enter")

							sendMessage({
								type: "sendMessage",
								chatroomId: parseInt(props.chatroomId, 10),
								message: newMessage
							})
						}
					}} 
					value={newMessage}
					onChange={e => {
						setNewMessage(e.target.value)
					}}
					/>
				</div>
			</div>
		</>
	)
}