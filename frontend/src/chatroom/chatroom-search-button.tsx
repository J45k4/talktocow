import { BsSearch } from "react-icons/bs"

export const ChatroomSearchButton = () => {
	return (
		<BsSearch onClick={() => {
			console.log("search button clicked")
		}} />
	)
}