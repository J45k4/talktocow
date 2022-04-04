import { useState } from "react"
import { Chatroom } from "./chatroom"
import { Diary } from "./diary/diary"
import { IsLoggedIn, IsNotLoggedIn, useIsLoggedIn } from "./isloggedin"
import { LoginForm } from "./login-form"
import { NavigationBar } from "./navigation_bar"
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
                <NavigationBar />
                <Diary /> 
                {/* <Chatroom chatroomId="1"/> */}
            </IsLoggedIn>
            <IsNotLoggedIn>
                <LoginForm />
            </IsNotLoggedIn>
        </div>
    )
}