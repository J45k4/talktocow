import { memo } from "react"
import styles from "./chatroom-message.module.css"
import { MessageIndicator } from "./icons/message-indicators"

export const formatMessageTime = (writenAt: Date) => {
    const hours = writenAt.getHours();
    const minues = writenAt.getMinutes();

    return `${hours}:${minues < 10 ? "0" + minues : minues}`
}

export const YourMessage = memo((props: {
    writenAt: string
    messageText: string
    status: "notsend" | "serverReceived" | "participantsReceived" | "participantsRead"
}) => {
    const writenAt = new Date(props.writenAt)

    return (
        <div className={styles.yourMessageLine}>
            <div className={styles.yourMessage}>
                <div className={styles.yourMessageText}>
                    {props.messageText}
                </div>
                <div style={{
                    display: "flex"
                }}>
                    <div className={styles.yourMessageTime}>
                        <span>
                            {isNaN(writenAt.getTime()) ? "" : formatMessageTime(writenAt)}
                        </span>
                        <span>
                            {props.status !== "notsend" && 
                                <MessageIndicator status={props.status} />}
                        </span>
                    </div>
                </div>
            </div>
        </div>
    )
})

export const ParticipantMessage = memo((props: {
    writenAt: string
    messageText: string
    status: "notsend" | "serverReceived" | "participantsReceived" | "participantsRead"
}) => {
    const writenAt = new Date(props.writenAt)

    return (
        <div className={styles.participantMessageLine}>
            <div className={styles.participantMessage}>
                <div className={styles.participantMessageText}>
                    {props.messageText}
                </div>
                <div style={{
                    display: "flex"
                }}>
                    <div className={styles.participantMessageTime}>
                        <span>
                            {isNaN(writenAt.getTime()) ? "" : formatMessageTime(writenAt)}
                        </span>
                    </div>
                </div>
            </div>
        </div>
    )
})