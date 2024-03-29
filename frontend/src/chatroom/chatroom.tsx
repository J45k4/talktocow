import React, { useCallback, useEffect, useRef, useState } from "react"

import styles from "./chatroom.module.css"
import { BsSearch, BsSnapchat } from "react-icons/bs"
import { IoIosCall } from "react-icons/io"
import { useChatroom, useChatroomMembers } from "../hokers"
import { useChatroomMessages } from "../use-chatroom-messages"
import { loadChatroomMessages, sendMessageToChatroom } from "../chatroom-message-managers"
import { ChatroomMessageRow } from "./chatroom-row"
import { Button } from "../components/button"
import { ChatInfo } from "./chat-info"
import { cache, useCache } from "../cache"
import { createLogger } from "../logger"
import { ws } from "../ws"

const logger = createLogger("chatroom")

const ChatroomSearch = () => {
	const [searching, setSearching] = useState(false)

	return (
		<div className={styles.chatroomTitleRightSide}>
			<input
				hidden={!searching}
				style={{
					fontSize: "20px",
				}}
				className={searching ?
					styles.showChatroomSearchInput :
					styles.hideChatroomSearchInput}
				type="text" />
			<BsSearch style={{
				width: "25px",
				height: "25px",
				padding: "10px",
				paddingRight: "15px",
				cursor: "pointer",
			}} onClick={(e) => {
				e.stopPropagation()
				e.preventDefault()
				setSearching(!searching)
			}} />
		</div>
	)
}

const ChatroomTitle = (props: {
	chatroomId: string
	onOpenInfo?: () => void
}) => {
	const members = useCache(cache.chatroom.get(props.chatroomId).members)
	const chatroom = useChatroom(props.chatroomId)

	return (
		<div className={styles.chatroomTitle}>
			<div className={styles.chatroomTitleLeftSide}
				onClick={props.onOpenInfo}>
				<BsSnapchat style={{
					width: "100%",
					height: "100%",
				}} />
			</div>
			<div className={styles.chatroomTitleCenter}>
				<div onClick={props.onOpenInfo}>
					<div className={styles.chatroomChatGroup}>
						{chatroom?.name}
					</div>
					<div className={styles.chatroomChatGroupMembers}>
						{members.map(m => m.name).join(", ")}
					</div>
				</div>
				
			</div>
			<ChatroomSearch />
			<IoIosCall style={{
				width: "25px",
				height: "25px",
				padding: "10px",
				cursor: "pointer",
			}} onClick={() => {
				ws.send({
					type: "createCall",
					chatroomId: props.chatroomId,
				})
			}} />
		</div>
	)
}

const ChatroomMessages = (props: {
	chatroomId: string
}) => {
	const messages = useChatroomMessages(props.chatroomId)
	const messagesEndRef = useRef(null);

	const scrollToBottom = () => {
		messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
	};

	useEffect(() => {
		logger.debug("scrolling to bottom")

		setTimeout(() => {
			scrollToBottom();
		}, 100);
	}, [messages]);

	useEffect(() => {
		loadChatroomMessages(props.chatroomId, 20)
	}, [props.chatroomId])

	return (
		<div className={styles.chatroomMessages}>
			{messages.map((p, index) => (
				<ChatroomMessageRow 
					key={p.reference} 
					chatroomMessage={p}
					grayBackground={index % 2 === 0} />
			))}
			<div ref={messagesEndRef} />
		</div>
	)
}

const ChatroomSendMessage = (props: {
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
			// flexGrow: 1,
			margin: "1em",
			fontSize: "1.5em",
			padding: "0.2em",
			display: "flex",
			flexDirection: "row",
		}}>
			<textarea style={{
				border: "none",
				width: "100%",
				height: "100%",
				outline: "none",
				resize: "none",
			}} 
			rows={1}
			onInput={e => {
				const target = e.currentTarget;
				target.style.height = "auto"; // reset the height to auto
				target.style.height = `${target.scrollHeight}px`; // set the height to the scrollHeight
				setNewMessage(target.value);
			}}
			onKeyDown={e => {
				if (e.key === "Enter" && !e.shiftKey) {
					console.log("Enter")

					sendMessage()
				}

				if (e.key === "Enter" && e.shiftKey) {
					logger.info("Shift + Enter")

					setNewMessage(newMessage + "\n")
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

export const Chatroom = (props: {
	chatroomId: string
}) => {
	const [infoVisible, setInfoVisible] = useState(false)

	return (
		<div className={styles.chatroom}>
			{infoVisible &&
			<div style={{
				position: "absolute",
				width: "500px",
				height: "500px",
				backgroundColor: "white",
				border: "solid 1px black",
			}}>
				<ChatInfo
					chatroomId={props.chatroomId}
					onClose={() => setInfoVisible(false)} />
			</div>}
			<ChatroomTitle 
				chatroomId={props.chatroomId}
				onOpenInfo={() => setInfoVisible(true)}
			/>
			<ChatroomMessages chatroomId={props.chatroomId} />
			<ChatroomSendMessage chatroomId={props.chatroomId} />
		</div>
	)
}