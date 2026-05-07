import React from "react"
import { createRoot } from "react-dom/client"
import { BrowserRouter } from "react-router-dom"

import App from "./App"
import { startNotificationHandler } from "./logic/notification-handler"
import { startOnlineWatch } from "./logic/online-indication"
import "../styles/globals.css"

startNotificationHandler()
startOnlineWatch()

createRoot(document.getElementById("root")!).render(
	<React.StrictMode>
		<BrowserRouter>
			<App />
		</BrowserRouter>
	</React.StrictMode>
)
