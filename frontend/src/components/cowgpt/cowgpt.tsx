import { useRouter } from "next/router"
import React from "react"
import { useWindowWiderThan } from "../../use-window-width"
import { ws } from "../../ws"
import { ChatroomSearchButton } from "./chatroom-search-button"
import { CowGPTChatroom } from "./cowgpt-chatroom"
import { CowGPTChatrooms } from "./cowgpt-chatrooms"

const ChatPortition = (props: {
	chatroomId?: string
	wideScreen?: boolean
}) => {
	const router = useRouter()

	return (
		<div style={{
			display: "flex",
			flexDirection: "column",
			flexGrow: 1,
		}}>
			{props.wideScreen && (
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
					<ChatroomSearchButton />
				</div>
			</div>)}
			{!props.wideScreen && (
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
			)}
			{props.chatroomId &&
			<CowGPTChatroom chatroomId={props.chatroomId} />}
		</div>
	)
}

export const SmallScreenChatrooms = () => {
	return (
		<div>
			<ChatroomSearchButton />
			<CowGPTChatrooms />
		</div>
		
	)
}

export const CowGPT = (props: {
	chatroomId?: string
}) => {
	const wideScreen = useWindowWiderThan(500)

	return (
		<div style={{
			display: "flex",
			flexDirection: "row",
			flexGrow: 1,
			height: "calc(100vh - 100px)"
		}}>
			{wideScreen && <>
			<div>
				<div>
					<button onClick={() => {
						ws.send({
							type: "askQuestion"
						})
					}}>
						Ask question
					</button>
				</div>
				<div>
					<CowGPTChatrooms selectedChatroomId={props.chatroomId} />
				</div>
			</div>
			<ChatPortition 
				chatroomId={props.chatroomId} 
				wideScreen={wideScreen} />
			</>}
			{!wideScreen &&
			(props.chatroomId ? 
			<ChatPortition 
				chatroomId={props.chatroomId}
				wideScreen={wideScreen} /> : 
			<SmallScreenChatrooms />)}
		</div>
	)
}