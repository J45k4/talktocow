import React, { useState } from "react";
import { useSearchUsers } from "../../data_hooks";
import { postJson } from "../../utility/talktocow-api-helpers";

export const AddStudentForm = (props: {
	courseId: string
}) => {
	const [searchWord, setSearchWord] = useState("")
	const [selectedUser, setSelectedUser] = useState<string | null>(null)
	const users = useSearchUsers(searchWord)

	return (
		<div style={{ border: "solid 1px black", maxWidth: "300px" }}>
			<input type="text" value={searchWord} onChange={e => {
				setSearchWord(e.target.value)
			}} />

			<div>
				{users.map(user => (
					<div key={user.id} onClick={() => {
						setSelectedUser(user.id)
					}} style={{
						cursor: "pointer",
						border: user.id === selectedUser ? "solid 1px black" : "" 
					}}>
						{user.name}
					</div>
				))}
			</div>

			<button onClick={async () => {
				await postJson(`/api/course/${props.courseId}/student`, {
					userId: selectedUser
				})
			}}>
				Add student
			</button>
		</div>
	)
}