import { FrontPage } from "../src/components/front-page";
import React from "react";
import { startOnlineWatch } from "../src/logic/online-indication";
import { PageContainer } from "../src/components/page-container";
import { Diary } from "../src/components/diary/diary";

startOnlineWatch()



const Index = () => {
	return (
		<PageContainer>
			<Diary />
		</PageContainer>
	)
}

export default Index