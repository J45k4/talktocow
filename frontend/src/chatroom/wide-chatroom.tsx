import React from "react"
import { Chatroom } from "./chatroom"
import { Chats } from "../chatroom/chatrooms"

import styles from "./wide-chatroom.module.css"
import Link from "next/link"
import { Button } from "../components"

const LeftSide = (props: {
	chatroomId?: string
}) => {
	return (
		<div className={styles.leftSide}>
			<div>
				<Link href="/chats/new">
					<Button
						title="New Chatroom"
					/>
				</Link>
			</div>
			<div>
				<Chats selectedChatroomId={props.chatroomId} />
			</div>
		</div>
	)
}

const RightSide = (props: {
	chatroomId?: string
}) => {
	return (
		<div className={styles.rightSide}>
			{/* <div style={{
				border: "solid 1px #8E8E8E",
				marginLeft: "1em",
				marginRight: "1em",
				marginBottom: "0.2em",
				fontSize: "1.5em",
				padding: "0.2em",
			}}>
				<ChatroomSearchButton />
			</div> */}
			{props.chatroomId &&
			<Chatroom chatroomId={props.chatroomId} />}
		</div>
	)
}

export const WideChatroom = (props: {
	chatroomId?: string
}) => {
	return (
		<div className={styles.wideChatroom}>
			<LeftSide chatroomId={props.chatroomId} />
			<RightSide chatroomId={props.chatroomId} />
		</div>	
	)
}