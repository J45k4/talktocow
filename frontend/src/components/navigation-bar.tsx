import React from "react"
import { Link, useLocation } from "react-router-dom"
import { clearSession, getSession, SessionChangeNotify, subscribeToSessionEvents, unsubscribeToSessionEvents } from "../logic/session-manager"
import { IsLoggedIn } from "./isloggedin"
import { useAddPasskey } from "../use-add-passkey"
import { useEffect, useState } from "react"
import styles from "./navigation-bar.module.css"

const NavigationBarItem = (props: {
    href: string
    text: string
}) => {
    const location = useLocation()

    const pathname = location.pathname || ""
    const isActive = pathname === props.href || pathname.startsWith(`${props.href}/`)

    return (
        <Link className={`${styles.navItem} ${isActive ? styles.activeNavItem : ""}`} to={props.href}>
            {props.text}
        </Link>
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
        <nav className={styles.navigationBar}>
            <IsLoggedIn>
                <div className={styles.navItems}>
                    {/* <NavigationBarItem href="/chatrooms" text="Chatrooms" /> */}
                    <NavigationBarItem href="/chats" text="Chats" />
                    <NavigationBarItem href="/diary" text="Diary" />
                    <NavigationBarItem href="/courses" text="Courses" />
                </div>
                <div className={styles.actions}>
                    {authMethod !== "passkey" ? (
                        <button className={styles.passkeyButton} onClick={addPasskey} disabled={loading}>
                            {loading ? "Waiting..." : "Add passkey"}
                        </button>
                    ) : null}
                    {error ? (
                        <div className={styles.error}>
                            {error}
                        </div>
                    ) : null}
                    <button className={styles.logoutButton} onClick={() => {
                        clearSession()
                    }}>
                        Logout
                    </button>
                </div>
            </IsLoggedIn>
        </nav>
    )
}
