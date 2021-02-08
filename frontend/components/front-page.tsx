import Link from "next/link"
import { Fragment } from "react"
import { Chatroom } from "./chatroom"
import { ConnectionIndicator } from "./connection-indicator"
import { IsLoggedIn, IsNotLoggedIn, useIsLoggedIn } from "./isloggedin"
import { LoginForm } from "./login-form"
import { LogoutButton } from "./logout-button"

export const FrontPage = () => {
    const isloggedin = useIsLoggedIn()

    console.log("FrontPage isLoggedIn", isloggedin)
    
    return (
        <div style={{
            position: "absolute",
            top: "0px",
            right: "0px",
            bottom: "0px",
            left: "0px"
        }}>
            <IsLoggedIn>
                <Chatroom chatroomId="123"/>
            </IsLoggedIn>
            <IsNotLoggedIn>
                <LoginForm />
            </IsNotLoggedIn>
        </div>
    )
}