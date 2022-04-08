import React from "react"
import { getSession } from "../logic/session-manager"
import { IsLoggedIn, IsNotLoggedIn, useIsLoggedIn } from "./isloggedin"
import { LoginForm } from "./login-form"
import { NavigationBar } from "./navigation_bar"

import styles from "./page-container.module.css"

export const PageContainer = (props: {
    children: any
}) => {
    // const isLoggedin = useIsLoggedIn()

    return (
        <div className={styles.pageContainer}>
            <NavigationBar />
            

            {/* <IsLoggedIn>
                <NavigationBar />
                {props.children}
            </IsLoggedIn> */}
            <IsNotLoggedIn>
                
                <LoginForm />
            </IsNotLoggedIn>
            <IsLoggedIn>
            {props.children}
            </IsLoggedIn>
            
        </div>
    )
}