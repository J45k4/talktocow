import React from "react"
import { NavigationBar } from "./navigation_bar"

export const PageContainer = (props: {
    children: any
}) => {
    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
            height: "100%"
        }}>
            <NavigationBar />

            {props.children}
        </div>
    )
}