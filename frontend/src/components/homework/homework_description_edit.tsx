
import React from "react"

export const HomeworkDescriptionEdit = (props: {
	description: string
	onChange: (description: string) => void
}) => {
	return (
		<div>
			<textarea value={props.description} onChange={e => {
				props.onChange(e.target.value)
			}} />
		</div>
	)
}