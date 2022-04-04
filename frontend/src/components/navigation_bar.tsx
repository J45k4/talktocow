import Link from "next/link"
import { useRouter } from "next/router"
import React from "react"

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
            display: "flex"
        }}>
            <NavigationBarItem href="/chatrooms" text="Chatrooms" />
            <NavigationBarItem href="/diary" text="Diary" />
        </div>
    )
}