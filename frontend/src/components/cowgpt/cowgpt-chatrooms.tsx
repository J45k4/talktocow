import Link from "next/link"
import React from "react"
import { useFetch } from "../../use-fetch"

export const CowGPTChatrooms = () => {
	const { data } = useFetch({
		path: "/api/chatrooms",
	})

	return (
		<div>
			{data?.map((chatroom: any) => {
				return (
					<Link href={`/cowgpt/${chatroom.id}`} style={{
						textDecoration: "none",
					}}>
						<div key={chatroom.id} style={{
							cursor: "pointer",
							margin: "10px",
						}}>
							{chatroom.name}
						</div>
					</Link>
					
				)
			})}
		</div>
	)
}