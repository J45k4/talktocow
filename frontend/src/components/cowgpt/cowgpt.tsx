import React from "react"
import { BsSearch } from "react-icons/bs"
import { sendMessage } from "../../ws"
import { CowGPTChatroom } from "./cowgpt-chatroom"
import { CowGPTChatrooms } from "./cowgpt-chatrooms"

export const CowGPT = (props: {
	chatroomId?: string
}) => {
	return (
		<div style={{
			display: "flex",
			flexDirection: "row",
			flexGrow: 1,
		}}>
			<div>
				<div>
					<button onClick={() => {
						sendMessage({
							type: "askQuestion"
						})
					}}>
						Ask question
					</button>
				</div>
				<div>
					<CowGPTChatrooms />
				</div>
			</div>
			<div style={{
				display: "flex",
				flexDirection: "column",
				flexGrow: 1,
			}}>
				<div>
					<div style={{
						border: "solid 1px #8E8E8E",
						flexGrow: 1,
						marginLeft: "1em",
						marginRight: "1em",
						marginBottom: "0.2em",
						fontSize: "1.5em",
						padding: "0.2em",
					}}>
						<BsSearch />
					</div>
					
				</div>
				{props.chatroomId &&
				<CowGPTChatroom chatroomId={props.chatroomId} />}
			</div>
		</div>
	)
}