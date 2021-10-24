import { FrontPage } from "../src/components/front-page";
import React from "react";
import { startOnlineWatch } from "../src/logic/online-indication";

startOnlineWatch()

const isServer = () => typeof window === 'undefined';

const Index = () => {
	return (
		<React.Fragment>
			{!isServer() && <FrontPage />}
		</React.Fragment>
	)
}

export default Index