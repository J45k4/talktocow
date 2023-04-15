import Link from "next/link"
import React, { useEffect, useState } from "react"
import { BiEdit } from "react-icons/bi"
import { Chatroom } from "../types"
import { api } from "../api"
import { cache, useCache } from "../cache"

import styles from "./chats.module.css"

const EditChatroomName = (props: {
	chatroomId: string
	onEditCompleted: () => void
}) => {
	const [name, setName] = useState("")

	return (
		<input type="text" 
			value={name}
			onChange={e => {
				setName(e.target.value)
			}}
			onKeyDown={e => {
				if (e.key == "Enter") {
					api.patchChatroom({
						chatroomId: props.chatroomId,
						name,
					}).then(() => {
						props.onEditCompleted()
					})
				}

				if (e.key == "Escape") {
					props.onEditCompleted()
				}	
			}} />
	)
}

const ChatRow = (props: {
	selectedChatroomId: string
	chatroom: Chatroom
	editing: boolean
	onStartEditing: () => void
	onEditCompleted: () => void
}) => {
	return (
		<Link key={props.chatroom.id} href={`/chats/${props.chatroom.id}`} style={{
			textDecoration: "none",
			marginRight: "10px",
		}}>
			<div style={{
				display: "flex",
				cursor: "pointer",
				padding: "10px",
				whiteSpace: "nowrap",
				border: props.selectedChatroomId == props.chatroom.id ? "solid 1px black" : ""
			}} onKeyDown={e => {
				if (e.key == "Escape") {
					props.onEditCompleted()
				}
			}}>
				<div style={{
					flex: 1,
					marginRight: "15px"
				}}>
					{props.editing && 
					<EditChatroomName
						chatroomId={props.chatroom.id}
						onEditCompleted={props.onEditCompleted} />}
					{!props.editing && props.chatroom.name}
				</div>
				<div>
					<BiEdit onClick={(e) => {
						e.stopPropagation()
						e.preventDefault()
						props.onStartEditing()
					}} />
				</div>
			</div>
		</Link>
	)
}

export const Chats = (props: {
	selectedChatroomId?: string
}) => {
	const [editingChatroomId, setEditingChatroomId] = useState("")

	const chatrooms = useCache(cache.myChatrooms)

	cache.chatroom.get("1234").get()

	useEffect(() => {
		window.onkeydown = (e) => {
			if (e.key == "Escape") {
				setEditingChatroomId("")
			}
		}
	}, [setEditingChatroomId])

	console.log("chatrooms", chatrooms)

	return (
		<div className={styles.chats}>
			{chatrooms.map((chatroom) => {
				return (
					<ChatRow 
						chatroom={chatroom}
						selectedChatroomId={props.selectedChatroomId}
						key={chatroom.id}
						editing={editingChatroomId == chatroom.id}
						onEditCompleted={() => {
							setEditingChatroomId("")
						}}
						onStartEditing={() => {
							setEditingChatroomId(chatroom.id)
						}}
					/>
				)
			})}
		</div>
	)
}