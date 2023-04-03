import React from "react"
import { WideChatroom } from "../../chatroom/wide-chatroom"
import { useWindowWiderThan } from "../../use-window-width"
import { NarrowChatroom } from "../../chatroom/narrow-chatroom"

import styles from "./cowgpt.module.css"

const Content = (props: {
	chatroomId?: string
}) => {
	const wideScreen = useWindowWiderThan(500)

	if (wideScreen) {
		return <WideChatroom chatroomId={props.chatroomId} />
	}

	return <NarrowChatroom chatroomId={props.chatroomId} />
}

export const CowGPT = (props: {
	chatroomId?: string
}) => {
	return (
		<div className={styles.chatrooms}>
			<Content chatroomId={props.chatroomId} />
		</div>	
	)
}