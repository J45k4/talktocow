import React from "react"
import { Chatroom } from "./chatroom"
import { Chats } from "./chats"

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
			<div style={{
				flexGrow: 1,
				overflowX: "hidden",
				overflowY: "auto",
			}}>
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