import { useRouter } from "next/router";
import { ChatroomView } from "../../src/chatroom/chatroom-view";
import { PageContainer } from "../../src/components/page_container";

export default function CowGPTChatroomPage() {
	const chatroomId = useRouter().query.chatroomId as string

	return (
		<PageContainer>
			<ChatroomView chatroomId={chatroomId} />
		</PageContainer>
	)
}