import { Conn } from "neffos.js";
import { CSSProperties, useState } from "react"
import { sendMessageToChatroom } from "../logic/chatroom-manager";

import styles from "./chatroom.module.css";
import { ConnectionIndicator } from "./connection-indicator";

export const Chatroom = (props: {
    style?: CSSProperties
    chatroomId: string
}) => {
    const [currentMessage, setCurrentMessage] = useState("");

    const handleSendMessage = () => {
        sendMessageToChatroom(props.chatroomId, currentMessage)

        setCurrentMessage("")
    }

    const messages = [];

    for (let i = 0; i < 35; i++) {
        messages.push(
            <div style={{
                height: "50px",
                flexShrink: 0
            }}>
                a
            </div>
        )
    }

    return (
        <div className={styles.root}>
            <div className={styles.header}>
                <ConnectionIndicator />
            </div>
            <div className={styles.body}>
                {messages}
            </div>
            <div className={styles.footer}>
                <div className={styles.emojiButton}>
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="currentColor" d="M9.153 11.603c.795 0 1.439-.879 1.439-1.962s-.644-1.962-1.439-1.962-1.439.879-1.439 1.962.644 1.962 1.439 1.962zm-3.204 1.362c-.026-.307-.131 5.218 6.063 5.551 6.066-.25 6.066-5.551 6.066-5.551-6.078 1.416-12.129 0-12.129 0zm11.363 1.108s-.669 1.959-5.051 1.959c-3.505 0-5.388-1.164-5.607-1.959 0 0 5.912 1.055 10.658 0zM11.804 1.011C5.609 1.011.978 6.033.978 12.228s4.826 10.761 11.021 10.761S23.02 18.423 23.02 12.228c.001-6.195-5.021-11.217-11.216-11.217zM12 21.354c-5.273 0-9.381-3.886-9.381-9.159s3.942-9.548 9.215-9.548 9.548 4.275 9.548 9.548c-.001 5.272-4.109 9.159-9.382 9.159zm3.108-9.751c.795 0 1.439-.879 1.439-1.962s-.644-1.962-1.439-1.962-1.439.879-1.439 1.962.644 1.962 1.439 1.962z"></path></svg>
                </div>
                <div className={styles.messageText}>
                    <input className={styles.input}
                        onFocus={() => {
                            // setDisplayEmojis(false)
                        }}
                        value={currentMessage} onChange={(e: any) => {
                            setCurrentMessage(e.target.value)
                        }}
                        onKeyDown={(e) => {
                            if (e.key === "Enter") {
                                handleSendMessage()
                            }
                        }}
                    />
                </div>
                <div className={styles.sendButton}
                    onClick={() => {
                        handleSendMessage()
                    }}>
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="currentColor" d="M1.101 21.757L23.8 12.028 1.101 2.3l.011 7.912 13.623 1.816-13.623 1.817-.011 7.912z"></path></svg>
                </div>
            </div>
        </div>
    )
}