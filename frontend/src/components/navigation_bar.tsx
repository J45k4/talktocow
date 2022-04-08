import Link from "next/link"
import { useRouter } from "next/router"
import React from "react"
import { clearSession } from "../logic/session-manager"
import { IsLoggedIn } from "./isloggedin"

const NavigationBarItem = (props: {
    href: string
    text: string
}) => {
    const router = useRouter()

    const pathname = router.pathname || ""

    return (
        <div style={{
            padding: "20px",
            border: pathname.includes(props.href) ? "solid 1px black" : ""
        }}>
            <Link href={props.href}>
                {props.text}
            </Link>
        </div>
    )
}

export const NavigationBar = () => {
    return (
        <div style={{
            display: "flex",
            flexDirection: "row",
            justifyContent: "space-between"
        }}>
            <IsLoggedIn>
                <div style={{
                    display: "flex",
                    flexDirection: "row"
                }}>
                    <NavigationBarItem href="/chatrooms" text="Chatrooms" />
                    <NavigationBarItem href="/diary" text="Diary" />
                </div>
                <div style={{
                    //display: "flex",
                    cursor: "pointer",
                    alignSelf: "center",
                    paddingRight: "20px"
                }} onClick={e => {
                    clearSession()
                }}>
                    Logout
                </div>
            </IsLoggedIn>
            
        </div>
    )
}