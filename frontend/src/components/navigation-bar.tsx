import React, { useRef } from "react"
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
    const [username, setUsername] = useState(getSession().username)
    const [isMenuOpen, setIsMenuOpen] = useState(false)
    const menuRef = useRef<HTMLDivElement>(null)
    const {
        addPasskey,
        loading,
        error
    } = useAddPasskey()

    useEffect(() => {
        const handle = (session: SessionChangeNotify) => {
            setAuthMethod(session.authMethod)
            setUsername(session.username)
        }

        subscribeToSessionEvents(handle)

        return () => {
            unsubscribeToSessionEvents(handle)
        }
    }, [])

    useEffect(() => {
        if (!isMenuOpen) {
            return
        }

        const closeOnOutsideClick = (event: MouseEvent) => {
            if (menuRef.current && !menuRef.current.contains(event.target as Node)) {
                setIsMenuOpen(false)
            }
        }

        document.addEventListener("mousedown", closeOnOutsideClick)

        return () => {
            document.removeEventListener("mousedown", closeOnOutsideClick)
        }
    }, [isMenuOpen])

    return (
        <nav className={styles.navigationBar}>
            <IsLoggedIn>
                <div className={styles.navItems}>
                    {/* <NavigationBarItem href="/chatrooms" text="Chatrooms" /> */}
                    <NavigationBarItem href="/" text="Home" />
                    <NavigationBarItem href="/chats" text="Chats" />
                    <NavigationBarItem href="/diary" text="Diary" />
                    <NavigationBarItem href="/courses" text="Courses" />
                </div>
                <div className={styles.actions} ref={menuRef}>
                    <button className={styles.menuButton} onClick={() => setIsMenuOpen(open => !open)} aria-expanded={isMenuOpen} aria-haspopup="menu">
                        {username || "Menu"} ▾
                    </button>
                    {isMenuOpen ? (
                        <div className={styles.dropdownMenu} role="menu">
                            {authMethod !== "passkey" ? (
                                <button className={styles.dropdownItem} onClick={addPasskey} disabled={loading} role="menuitem">
                                    {loading ? "Waiting..." : "Add passkey"}
                                </button>
                            ) : null}
                            <button className={styles.dropdownItem} onClick={() => {
                                clearSession()
                                setIsMenuOpen(false)
                            }} role="menuitem">
                                Logout
                            </button>
                            {error ? (
                                <div className={styles.error}>
                                    {error}
                                </div>
                            ) : null}
                        </div>
                    ) : null}
                </div>
            </IsLoggedIn>
        </nav>
    )
}
