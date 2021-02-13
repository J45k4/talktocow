import { Chatroom } from "./chatroom"
import { IsLoggedIn, IsNotLoggedIn, useIsLoggedIn } from "./isloggedin"
import { LoginForm } from "./login-form"

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
                <Chatroom chatroomId="1"/>
            </IsLoggedIn>
            <IsNotLoggedIn>
                <LoginForm />
            </IsNotLoggedIn>
        </div>
    )
}