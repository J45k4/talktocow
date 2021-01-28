import React, { useEffect, useState } from "react";

import { useRouter } from "next/router";
import Link from "next/link";
import Head from 'next/head'
import { postJson } from "../utility/api";
import { IsLoggedIn, IsNotLoggedIn } from "../components/auth";
import { Workbox } from "workbox-window";


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

	useEffect(() => {
		const wb = new Workbox("sw.js", { scope: "/" });

		wb.register();
	}, []);

	const [username, setUsername] = useState();
	const [password, setPassword] = useState();

	const [currentlyLogining, setCurrentlyLogining] = useState(false);
	const [loginError, setLoginError] = useState("");

	const router = useRouter()

	return (
		<div>
			<Head>
				<link rel="manifest" href="/manifest.json" />
				<title>Talktocow 🥰 </title>
				<meta name="viewport" content="initial-scale=1.0, width=device-width" />
			</Head>

			<div style={{
				display: "flex",
				justifyContent: "center"
			}}>
				{/* <h1>
					<div>
						Hello cute cat
					</div>
					<div>
						love you a lot
					</div>
					<div>
						🥰🥰🥰🥰
					  </div>
				</h1> */}

				<div>
					<IsLoggedIn>
						<div>
							<h1>Hello lovely cat 😘</h1>
							<div>
								<Link href="/chatroom">
									Enter chatroom
							</Link>
							</div>
							<div>
								<Link href="/newchatroom">
									New chatroom
							</Link>
							</div>
							<div>
								<button onClick={() => {
									localStorage.removeItem("token");

									window.location.reload();
								}}>
									Log out
								</button>
							</div>
						</div>
					</IsLoggedIn>
					<IsNotLoggedIn>
						<div>
							<h1>
								Please login cute cat 🥰
							</h1>
						</div>
						<div>
							<div>
								<label>
									Username
								</label>
							</div>
							<input type="text" value={username} onChange={(e: any) => {
								setUsername(e.target.value)
							}} />
						</div>
						<div>
							<div>
								<label>
									Password
								</label>
							</div>
							<input type="password" value={password} onChange={(e: any) => {
								setPassword(e.target.value)
							}} />
						</div>
						<div>
							<div style={{
								color: "red"
							}}>
								{loginError}
							</div>
							{currentlyLogining ? 
								<div>
									Logining...
								</div>
							:
							<button onClick={async () => {
								setCurrentlyLogining(true);

								let res

								try {
									res = await postJson("/api/login", {
										username: username,
										password: password
									})
								} catch (e) {
									console.log("cathing erro", e);

									setCurrentlyLogining(false);
									setLoginError(e.message);

									return
								}

								if (res.errorMessage) {
									console.error(res.errorMessage)

									return
								}

								localStorage.setItem("token", res.token)

								window.location.reload();
							}}>
								Login
							</button>}
						</div>
					</IsNotLoggedIn>
				</div>
			</div>
		</div>
	)
}

export default Index