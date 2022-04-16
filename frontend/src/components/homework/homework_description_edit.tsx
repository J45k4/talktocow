
import React from "react"

export const HomeworkDescriptionEdit = (props: {
	description: string
}) => {
	return (
		<div>
			<textarea value={props.description} />
		</div>
	)
}