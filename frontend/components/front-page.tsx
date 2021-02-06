import Link from "next/link"
import { Fragment } from "react"
import { Chatroom } from "./chatroom"
import { ConnectionIndicator } from "./connection-indicator"
import { IsLoggedIn, IsNotLoggedIn } from "./isloggedin"
import { LoginForm } from "./login-form"
import { LogoutButton } from "./logout-button"

export const FrontPage = () => {
    // return (
    //     <div style={{
    //         position: "absolute",
    //         top: "0px",
    //         right: "0px",
    //         bottom: "0px",
    //         left: "0px"
    //     }}>
    //         <Chatroom chatroomId="123"/>
    //     </div>
    // )
    
    return (
        <div>
            {/* <IsLoggedIn> */}
                <div style={{
                    position: "absolute",
                    top: "0px",
                    right: "0px",
                    bottom: "0px",
                    left: "0px"
                }}>
                    <Chatroom chatroomId="123"/>
                </div>
            {/* </IsLoggedIn> */}
            <IsNotLoggedIn>
                <LoginForm />
            </IsNotLoggedIn>
        </div>
    )
}