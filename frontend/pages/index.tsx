import React, { useEffect, useState } from "react";

import { useRouter } from "next/router";
import Link from "next/link";
import Head from 'next/head'
import { postJson } from "../utility/api";
import { Workbox } from "workbox-window";
import { startOnlineWatch } from "../logic/online-indication";
import { ConnectionIndicator } from "../components/connection-indicator";
import { IsLoggedIn, IsNotLoggedIn } from "../components/isloggedin";
import { LoginForm } from "../components/login-form";
import { LogoutButton } from "../components/logout-button";
import { FrontPage } from "../components/front-page";

startOnlineWatch()


const isServer = () => typeof window === 'undefined';

const Index = () => {
	// useEffect(() => {
	// 	// const w = new WebSocket("ws://localhost:12001/api/socket")

	// 	// w.onopen = () => {
	// 	// 	console.log("new connection opened");
	// 	// }

	// 	// w.onclose = () => {
	// 	// 	console.log("Connection closed")
	// 	// }
	// 	Notification.requestPermission().then(() => {
	// 		// var notification = new Notification("Cow send you a cute message");
	// 		// ServiceWorkerRegistration.showNotification("Hello")
	// 	})

	// }, [])

	// useEffect(() => {
	// 	const wb = new Workbox("sw.js", { scope: "/" });

	// 	wb.register();
	// }, []);

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