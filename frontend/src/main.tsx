import React from "react"
import { createRoot } from "react-dom/client"
import { BrowserRouter } from "react-router-dom"

import App from "./App"
import { startNotificationHandler } from "./logic/notification-handler"
import { startOnlineWatch } from "./logic/online-indication"
import { waitForSessionReady } from "./logic/session-manager"
import "../styles/globals.css"

const startApp = async () => {
	await waitForSessionReady()

	startNotificationHandler()
	startOnlineWatch()

	createRoot(document.getElementById("root")!).render(
		<React.StrictMode>
			<BrowserRouter>
				<App />
			</BrowserRouter>
		</React.StrictMode>
	)
}

void startApp()
