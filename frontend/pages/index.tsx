import { FrontPage } from "../src/components/front-page";
import React from "react";
import { PageContainer } from "../src/components/page-container";
import { Diary } from "../src/components/diary/diary";

const Index = () => {
	return (
		<PageContainer>
			<Diary />
		</PageContainer>
	)
}

export default Index
