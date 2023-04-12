import React, { useState } from "react"
import { postJson } from "../../api-methods";

export const NewCourseForm = (props: {
	onCourseCreated: (courseId: string) => void;
}) => {
	const [ name, setName ] = useState("")

	return (
		<div>
			<div>
				Name
			</div>
			<div>
				<input type="text" value={name} onChange={e => {
					setName(e.target.value)
				}} />
			</div>
			<div>
				<button onClick={() => {
					postJson<any>("/api/course", {
						name: name
					}).then(r => {
						props.onCourseCreated(r.payload.id)
					})
				}}>
					Create
				</button>
			</div>
		</div>
	);
}