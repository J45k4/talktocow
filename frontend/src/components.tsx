import React, { CSSProperties } from "react"

export const Button = (props: {
	title: string
	style?: CSSProperties
	onClick?: () => void
}) => {
	return (
		<button style={{
			height: "30px",
			border: "solid 1px black",
			fontSize: "1em",
		}}
		onClick={() => {
			if (props.onClick) {
				props.onClick()
			}
		}}
		>
			{props.title}
		</button>
	)
}