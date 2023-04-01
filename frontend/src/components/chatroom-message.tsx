import { memo } from "react"
import styles from "./chatroom-message.module.css"
import { MessageIndicator } from "./icons/message-indicators"

export const formatMessageTime = (writtenAt: Date) => {
    const hours = writtenAt.getHours();
    const minues = writtenAt.getMinutes();

    return `${hours}:${minues < 10 ? "0" + minues : minues}`
}

export const YourMessage = memo(function YourMessage(props: {
    writtenAt: string
    messageText: string
    status: "notsend" | "serverReceived" | "participantsReceived" | "participantsRead"
}) {
    const writtenAt = new Date(props.writtenAt)

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
                            {isNaN(writtenAt.getTime()) ? "" : formatMessageTime(writtenAt)}
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

export const ParticipantMessage = memo(function ParticipantMessage(props: {
    writtenAt: string
    messageText: string
    status: "notsend" | "serverReceived" | "participantsReceived" | "participantsRead"
}) {
    const writtenAt = new Date(props.writtenAt)

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
                            {isNaN(writtenAt.getTime()) ? "" : formatMessageTime(writtenAt)}
                        </span>
                    </div>
                </div>
            </div>
        </div>
    )
})