import { CSSProperties } from "react"
import { BsSearch } from "react-icons/bs"

export const ChatroomSearchButton = (props: {
	style?: CSSProperties
}) => {
	return (
		<BsSearch style={props.style} onClick={() => {
			console.log("search button clicked")
		}} />
	)
}