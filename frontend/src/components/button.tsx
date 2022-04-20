import React from "react"

export const Button = (props: {
	children: any
	onClick?: () => void
}) => {
	return (
		<button style={{
			maxWidth: "70px",
		}} onClick={props.onClick}>
			{props.children}
		</button>
	)
}