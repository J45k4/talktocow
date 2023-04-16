import React, { useState, useEffect } from 'react';

import styles from "./chat-info.module.css";
import { useChatroomMembers } from '../hokers';
import { useCache } from '../cache';
import { cache } from '../cache';
import { api } from '../api';
import { getSession } from '../logic/session-manager';
import { createLogger } from '../logger';

const logger = createLogger("chat-info")

const NavigationSide = (props: {
	onClose: () => void	
}) => {
	return (
		<div className={styles.navigationSide}>
			<button onClick={props.onClose}>
				Close
			</button>
			<div>
				<div>
					Members
				</div>
			</div>
		</div>
	)
}

const AddMember = (props: {
	chatroomId: string
}) => {
	const users = useCache(cache.users)
	const memebers = useCache(cache.chatroom.get(props.chatroomId).members)

	return (
		<div>
			{users.map(u => {
				const selected = memebers.some(m => m.id === u.id)
				
				return (
					<div key={u.id}>
						{u.name}
						<input type="checkbox"
							checked={selected}
							onChange={(e) => {
								if (selected) {
									const session = getSession()

									if (session?.userId == u.id) {
										return
									}

									api.removeChatroomMember({
										chatroomId: props.chatroomId,
										userId: u.id,
									})

									return
								}

								api.addChatroomMember({
									chatroomId: props.chatroomId,
									userId: u.id,
								})
							}}
						/>
					</div>
				)
			})}
		</div>
	)
}

const ChatMemebers = (props: {
	chatroomId: string
	onAddMember?: () => void
}) => {
	const memebers = useChatroomMembers(props.chatroomId)
	
	return (
		<div>
			<div>
				<button onClick={props.onAddMember}>
					Add member
				</button>
			</div>
			{memebers.map(m => (
				<div key={m.id}>
					{m.name}
				</div>
			))}
		</div>
	)
}

const ChatInfoSide = (props: {
	chatroomId: string
}) => {
	const [state, setState] = useState<"members" | "add-member">("members")



	return (
		<div className={styles.chatInfoSide}>
			{state === "members" && 
			<ChatMemebers
				onAddMember={() => setState("add-member")}
				chatroomId={props.chatroomId}
			/>}

			{state === "add-member" && <AddMember
				chatroomId={props.chatroomId}
			/>}
			
		</div>
	)
}

export const ChatInfo = (props: {
	onClose: () => void
	chatroomId: string
}) => {
	return (
		<div className={styles.chatInfo}>
			<NavigationSide onClose={props.onClose} />
			<ChatInfoSide chatroomId={props.chatroomId} />
		</div>
	)
}
