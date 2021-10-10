import React, { useEffect, useState } from "react"
import { isSocketConnected, SocketStatusChanged, subscribeToSocketStatusChanged, unsubscribeToSocketStatusChanged } from "../logic/websocket-conn"

export const useIsSocketConnected = () => {
    const [isConnected, setIsConnected] = useState(isSocketConnected())

    useEffect(() => {
        function handle(e: SocketStatusChanged) {
            setIsConnected(e.connected)
        }

        subscribeToSocketStatusChanged(handle)

        return () => {
            unsubscribeToSocketStatusChanged(handle)
        }
    }, [])

    return isConnected
}

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