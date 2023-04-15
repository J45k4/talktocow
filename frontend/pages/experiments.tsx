import React from "react"
import { PageContainer } from "../src/components/page-container"
import { Window } from "../src/window"
import { VideoCapture } from "../src/video/video-capture"

export default function ExperimentsPage() {
	const [show, setShow] = React.useState(true)

	return (
		<PageContainer>
			{show &&
			<Window title="Test" onClose={() => {
				setShow(false)
			}}>
				<VideoCapture
					constraints={{
						video: true,
					}}
					style={{
						width: "100%",
					}}
				/>
			</Window>}
		</PageContainer>
	)
}