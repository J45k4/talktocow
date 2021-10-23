import React from "react";
import Head from 'next/head'
import { startOnlineWatch } from "../src/logic/online-indication";
import { FrontPage } from "../src/components/front-page";

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