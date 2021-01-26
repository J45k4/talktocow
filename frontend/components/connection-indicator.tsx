import { useIsSocketConnected } from "../websocket-conn"
import React from "react"

export const ConnectionIndicator = () => {
    const isConnected = useIsSocketConnected()

    if (isConnected === true) {
        return (
            <div style={{ color: "green" }}>
                Connected
            </div>
        )
    }

    return (
        <div style={{
            color: "red"
        }}>
            Disconnected
        </div>
    )
}