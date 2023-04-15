import React from "react"
import { IsLoggedIn, IsNotLoggedIn } from "./isloggedin"
import { LoginForm } from "./login-form"
import { NavigationBar } from "./navigation-bar"

import styles from "./page-container.module.css"

export const PageContainer = (props: {
    children: any
}) => {
    return (
        <div className={styles.pageContainer}>
            <NavigationBar />
            <IsNotLoggedIn>
                <LoginForm />
            </IsNotLoggedIn>
            <IsLoggedIn>
				<div className={styles.pageContainerContent}>
					{props.children}
				</div>

				{/* {props.children} */}
            </IsLoggedIn>    
        </div>
    )
}