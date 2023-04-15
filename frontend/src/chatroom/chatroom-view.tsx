import React from "react"
import { WideChatroom } from "./wide-chatroom"
import { useWindowWiderThan } from "../use-window-width"
import { NarrowChatroom } from "./narrow-chatroom"

import styles from "./chatroom-view.module.css"

const Content = (props: {
	chatroomId?: string
}) => {
	const wideScreen = useWindowWiderThan(500)
	// const creatingChatroom = useReactVar(creatingChatroomState)

	// if (creatingChatroom) {
	// 	return <NewChatForm />
	// }

	if (wideScreen) {
		return <WideChatroom chatroomId={props.chatroomId} />
	}

	return <NarrowChatroom chatroomId={props.chatroomId} />
}

export const ChatroomView = (props: {
	chatroomId?: string
}) => {
	return (
		<div className={styles.chatrooms}>
			<Content chatroomId={props.chatroomId} />
		</div>	
	)
}