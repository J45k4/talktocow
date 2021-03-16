import { useState } from "react"
import { Chatroom } from "./chatroom"
import { IsLoggedIn, IsNotLoggedIn, useIsLoggedIn } from "./isloggedin"
import { LoginForm } from "./login-form"
import { ValentinesDayGift } from "./valentines-day-gift"

export const FrontPage = () => {
    return (
        <div style={{
            position: "absolute",
            top: "0px",
            right: "0px",
            bottom: "0px",
            left: "0px"
        }}>
            <IsLoggedIn>
                <Chatroom chatroomId="1"/>
            </IsLoggedIn>
            <IsNotLoggedIn>
                <LoginForm />
            </IsNotLoggedIn>
        </div>
    )
}