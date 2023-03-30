import { useRouter } from "next/router";
import { CowGPT } from "../../src/components/cowgpt/cowgpt";
import { PageContainer } from "../../src/components/page_container";

export default function CowGPTChatroomPage() {
	const chatroomId = useRouter().query.chatroomId as string

	return (
		<PageContainer>
			<CowGPT chatroomId={chatroomId} />
		</PageContainer>
	)
}