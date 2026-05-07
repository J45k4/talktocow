import { useParams } from "react-router-dom";
import { ChatroomView } from "../../src/chatroom/chatroom-view";
import { PageContainer } from "../../src/components/page-container";

export default function CowGPTChatroomPage() {
	const { chatroomId } = useParams()

	return (
		<PageContainer>
			<ChatroomView chatroomId={chatroomId} />
		</PageContainer>
	)
}
