import { dial } from "neffos.js";
import React, { useEffect, useState } from "react";
import { ConnectionIndicator } from "../components/connection-indicator";
import { MessageFromServer, subscribeToMessages, unsubscribeToMessages, useMessages, useSocketConn } from "../websocket-conn";
import { getJson } from "../utility/api";


const chatroom = () => {
    const { sendMessage } = useSocketConn()

    const [currentMessage, setCurrentMessage] = useState("");

    const [allMesages, setAllMessages] = useState<string[]>([]);

    useEffect(() => {
        function handle(msg: MessageFromServer) {
            const transmittedAtDate = new Date(msg.newChatroomMessage.transmittedAt)

            const transmittedAtDateString = transmittedAtDate.toLocaleString()

            setAllMessages([
                ...allMesages,
                transmittedAtDateString + " " + msg.newChatroomMessage.fromUserName + ": " + msg.newChatroomMessage.messageText,
            ])
        }

        subscribeToMessages(handle)

        return () => {
            unsubscribeToMessages(handle)
        }
    }, [allMesages])

    useEffect(() => {
        getJson("/api/messages").then(r => {
            console.log(r)

            setAllMessages(r.reverse().map(msg => {
                const transmittedAtDate = new Date(msg.transmited_at)

                const transmittedAtDateString = transmittedAtDate.toLocaleString()

                return transmittedAtDateString + " " + msg.name + ": " + msg.message_text
            }))
        })
    }, [])

    return (
        <div style={{
            position: "fixed",
            display: "flex",
            alignItems: "flex-end",
            top: "0px",
            left: "40px",
            right: "0px",
            bottom: "10px"
        }}>
            <div>
                <div style={{
                    overflowY: "auto"
                }}>
                    {allMesages.map(p => (
                        <div>
                            {p}
                        </div>
                    ))}
                </div>
                <div>
                    <input type="text" value={currentMessage} onChange={(e: any) => {
                        setCurrentMessage(e.target.value)
                    }}
                    onKeyDown={(e) => {
                        if (e.key === "Enter") {
                            const res = sendMessage({
                                messageToChatroom: {
                                    messageText: currentMessage,
                                    createTime: new Date().toISOString(),
                                    transmitTime: new Date().toISOString()
                                }
                            })
    
                            if (res === false) {
                                return
                            }

                            setCurrentMessage("")
                        }
                    }}
                    style={{
                        height: "2em",
                        width: "200px"
                    }} /> <button onClick={() => {
                        const res = sendMessage({
                            messageToChatroom: {
                                messageText: currentMessage,
                                createTime: new Date().toISOString(),
                                transmitTime: new Date().toISOString()
                            }
                        })

                        if (res === false) {
                            return
                        }

                        setCurrentMessage("")
                    }} style={{
                        height: "2.5em",
                        
                    }}>
                        Send message
                    </button>
                </div>
            </div>
        </div>

    )
}

export default chatroom