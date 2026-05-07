import React from "react"
import { Link, useLocation } from "react-router-dom"
import { clearSession, getSession, SessionChangeNotify, subscribeToSessionEvents, unsubscribeToSessionEvents } from "../logic/session-manager"
import { IsLoggedIn } from "./isloggedin"
import { useAddPasskey } from "../use-add-passkey"
import { useEffect, useState } from "react"

const NavigationBarItem = (props: {
    href: string
    text: string
}) => {
    const location = useLocation()

    const pathname = location.pathname || ""

    return (
        <div style={{
            padding: "20px",
            border: pathname.includes(props.href) ? "solid 1px black" : ""
        }}>
            <Link to={props.href}>
                {props.text}
            </Link>
        </div>
    )
}

export const NavigationBar = () => {
	const [authMethod, setAuthMethod] = useState(getSession().authMethod)
	const {
		addPasskey,
		loading,
		error
	} = useAddPasskey()

	useEffect(() => {
		const handle = (session: SessionChangeNotify) => {
			setAuthMethod(session.authMethod)
		}

		subscribeToSessionEvents(handle)

		return () => {
			unsubscribeToSessionEvents(handle)
		}
	}, [])

    return (
        <div style={{
			height: "55px",
            display: "flex",
            flexDirection: "row",
            justifyContent: "space-between"
        }}>
            <IsLoggedIn>
                <div style={{
                    display: "flex",
                    flexDirection: "row"
                }}>
                    {/* <NavigationBarItem href="/chatrooms" text="Chatrooms" /> */}
					<NavigationBarItem href="/chats" text="Chats" />
                    <NavigationBarItem href="/diary" text="Diary" />
					<NavigationBarItem href="/courses" text="Courses" />
					<NavigationBarItem href="/profile" text="Profile" />
                </div>
                <div style={{
                    display: "flex",
					gap: "14px",
                    alignSelf: "center",
                    paddingRight: "20px",
					alignItems: "center"
                }}>
					{authMethod !== "passkey" ? (
						<button onClick={addPasskey} disabled={loading}>
							{loading ? "Waiting for passkey..." : "Add passkey"}
						</button>
					) : null}
					{error ? (
						<div style={{
							color: "red"
						}}>
							{error}
						</div>
					) : null}
					<div style={{
						cursor: "pointer"
					}} onClick={() => {
						clearSession()
					}}>
						Logout
					</div>
                </div>
            </IsLoggedIn>
            
        </div>
    )
}
