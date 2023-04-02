// import React from "react"
// import { ChatroomSearchButton } from "../components/cowgpt/chatroom-search-button"
// import { CowGPTChatroom } from "../components/cowgpt/cowgpt-chatroom"

// const LeftSide = () => {
// 	return (
// 		<div>
// 			<div>
// 				<button onClick={() => {
// 					ws.send({
// 						type: "askQuestion"
// 					})
// 				}}>
// 					Ask question
// 				</button>
// 			</div>
// 			<div>
// 				<CowGPTChatrooms selectedChatroomId={props.chatroomId} />
// 			</div>
// 		</div>
// 	)
// }

// const RightSide = (props: {
// 	chatroomId?: string
// }) => {
// 	return (
// 		<div>
// 			<div style={{
// 					border: "solid 1px #8E8E8E",
// 					flexGrow: 1,
// 					marginLeft: "1em",
// 					marginRight: "1em",
// 					marginBottom: "0.2em",
// 					fontSize: "1.5em",
// 					padding: "0.2em",
// 			}}>
// 				<ChatroomSearchButton />
// 			</div>
// 			{props.chatroomId &&
// 			<CowGPTChatroom chatroomId={props.chatroomId} />}
// 		</div>
// 	)
// }

// export const WideChatroom = () => {
// 	return (
// 		<div>
// 			<ChatPortition 
// 				chatroomId={props.chatroomId} 
// 				wideScreen={wideScreen} />
// 		</div>
// 	)
// }