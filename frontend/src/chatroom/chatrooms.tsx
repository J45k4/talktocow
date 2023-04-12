import Link from "next/link"
import React, { useState } from "react"
import { BiEdit } from "react-icons/bi"
import { Chatroom } from "../types"
import { useChatrooms } from "../hokers"
import { api } from "../api"

const EditChatroomName = (props: {
	chatroomId: string
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
					})
				}
			}} />
	)
}

const ChatRow = (props: {
	selectedChatroomId: string
	chatroom: Chatroom
}) => {
	const [editMode, setEditMode] = React.useState(false)

	return (
		<div style={{
			cursor: "pointer",
			padding: "10px",
			whiteSpace: "nowrap",
			border: props.selectedChatroomId == props.chatroom.id ? "solid 1px black" : ""
		}}>
			<Link key={props.chatroom.id} href={`/chats/${props.chatroom.id}`} style={{
				textDecoration: "none",
				marginRight: "10px",
			}}>
				{editMode && 
				<EditChatroomName
					chatroomId={props.chatroom.id} />}
				{!editMode && props.chatroom.name}
			</Link>
			<BiEdit onClick={() => {
				setEditMode(!editMode)
			}} />
		</div>
	)
}

export const Chats = (props: {
	selectedChatroomId?: string
}) => {
	const chatrooms = useChatrooms()

	return (
		<div>
			{chatrooms.map((chatroom) => {
				return (
					<ChatRow 
						chatroom={chatroom}
						selectedChatroomId={props.selectedChatroomId}
						key={chatroom.id}
					/>
				)
			})}
		</div>
	)
}