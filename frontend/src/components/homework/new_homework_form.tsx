import React, { useState } from "react"
import { postJson } from "../../utility/talktocow-api-helpers";

export const NewHomeworkForm = (props: {
	courseId: string
	onHomeworkCreated: (homeworkId: string) => void
}) => {
	const [title, setTitle] = useState("")
	
	return (
		<div>
			<div>
				Title
			</div>
			<div>
				<input type="text" value={title} onChange={e => {
					setTitle(e.target.value)
				}} />
			</div>
			<div>
				<button onClick={() => {
					postJson<any>(`/api/course/${props.courseId}/homework`, {
						title: title
					}).then(r => {
						props.onHomeworkCreated(r.payload.id)
					})
				}}>
					Create
				</button>
			</div>
		</div>
	);
}