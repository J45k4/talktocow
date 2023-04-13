import React, { useEffect } from "react"
import { api } from "../api"

import styles from "./new-chat.module.css"
import { useReactVar } from "../use-reactive-var"
import { useUsers } from "../hokers"
import { creatingChatroomState, selectedUsersState } from "../state"
import Link from "next/link"
import { useRouter } from "next/router"
import { Button } from "../components"

const SelectUsers = () => {
	const users = useUsers()
	const selectedUsers = useReactVar(selectedUsersState)

	useEffect(() => {
		api.fetchUsers()
	}, [])

	return (
		<div className={styles.selectUsers}>
			Select users
			<div>
				{users.map(user => {
					const selected = selectedUsers.includes(user.id)

					return (
						<div 
							key={user.id} 
							className={styles.userRow}
							style={{
								backgroundColor: selected ? "grey" : "white",
							}}
							onClick={() => {
								if (selected) {
									selectedUsersState.set(selectedUsers.filter(id => id != user.id))
								} else {
									selectedUsersState.set([...selectedUsers, user.id])
								}
							}}>
							{user.name}
						</div>
					)
				})}
			</div>
		</div>
	)
}

export const NewChatForm = () => {
	const router = useRouter()

	const [newChatroomName, setNewChatroomName] = React.useState("")

	return (
		<div>
			Create new chat
			<div>
				<input 
					type="text" 
					placeholder="Chatroom name"
					value={newChatroomName}
					onChange={e => {
						setNewChatroomName(e.target.value)
					}} />
			</div>
			<SelectUsers />
			<div>
				<Link href="/chats">
					<Button
						title="Cancel"
					/>
				</Link>
				<Button onClick={() => {
					api.createChatroom({
						userIds: selectedUsersState.get(),
						name: newChatroomName,
					}).then(res => {
						router.push(`/chats/${res.payload.id}`)
					})
				}} title="Create chatroom" />
			</div>
		</div>
	)
}