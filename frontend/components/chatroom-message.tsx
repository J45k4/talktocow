import styles from "./chatroom-message.module.css"
import { MessageIndicator } from "./icons/message-indicators"


export const YourMessage = (props: {
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
                            {`${writenAt.getHours()}:${writenAt.getMinutes()}`}
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
}

export const ParticipantMessage = (props: {
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
                            {`${writenAt.getHours()}:${writenAt.getMinutes()}`}
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
}