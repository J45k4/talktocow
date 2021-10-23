import React from "react";
import Head from 'next/head'
import { startOnlineWatch } from "../src/logic/online-indication";
import { FrontPage } from "../src/components/front-page";

startOnlineWatch()

const isServer = () => typeof window === 'undefined';

const Index = () => {
	return (
		<React.Fragment>
			<Head>
				<link rel="manifest" href="/manifest.json" />
				<title>Talktocow 🥰 </title>
				<meta name="viewport" content="initial-scale=1.0, maximum-scale=1.0, width=device-width, user-scalable=no" />
			</Head>
			{!isServer() && <FrontPage />}
					
		</React.Fragment>
	)
}

export default Index