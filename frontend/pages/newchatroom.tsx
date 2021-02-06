import { useEffect, useState } from "react"
import { startOnlineWatch } from "../logic/online-indication";
import { getJson } from "../utility/api";
import { MessageFromServer, subscribeToMessages, unsubscribeToMessages, useSocketConn } from "../websocket-conn";

startOnlineWatch();

const Newchatroom = () => {
    const [ displayEmojis, setDisplayEmojis ] = useState(false);

    const { sendMessage } = useSocketConn()

    const [currentMessage, setCurrentMessage] = useState("");

    const [allMesages, setAllMessages] = useState<string[]>([]);

    useEffect(() => {
        function handle(msg: MessageFromServer) {
            const transmittedAtDate = new Date(msg.newChatroomMessage.transmittedAt)

            const transmittedAtDateString = transmittedAtDate.toLocaleTimeString()

            setAllMessages([
                transmittedAtDateString + " " + msg.newChatroomMessage.fromUserName + ": " + msg.newChatroomMessage.messageText,
                ...allMesages,
                
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

            setAllMessages(r.map(msg => {
                const transmittedAtDate = new Date(msg.transmited_at)

                const transmittedAtDateString = transmittedAtDate.toLocaleTimeString()

                return transmittedAtDateString + " " + msg.name + ": " + msg.message_text
            }))
        })
    }, [])

    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
            position: "fixed",
            left: "0px",
            top: "0x",
            right: "0px",
            bottom: "0px",
            border: "solid 1px black",
            width: "100%",
            height: "100%"
        }}>
            <div style={{
                backgroundColor: "#2a2f32",
                color: "#e1e2e3",
                height: "62px"
            }}>

            </div>
            <div style={{
                flexGrow: 1,
                backgroundColor: "#525252",
                display: "flex",
                flexDirection: "column-reverse",
                overflowY: "auto",
                flexShrink: 0

            }}>
                {/* <div> */}
                {allMesages.map(p => (
                    <div style={{
                        marginTop: "10px",
                        marginBottom: "10px",
                        marginLeft: "10px",
                        marginRight: "10px",
                    }}>
                    <span style={{
                        backgroundColor: p.includes("Li") ? "#056162" : "#0d1418",
                        textAlign: "right",
                        color: "white",
                        borderRadius: "5px",
                        padding: "5px"
                    }}>
                        {p}
                    </span>
                    </div>
                ))}
                {/* </div> */}
            </div>
            <div style={{
                display: "flex",
                height: "62px",
                backgroundColor: "#2a2f32",
                justifyContent: "center",
                flexShrink: 0
            }}> 
                <div style={{
                    color: "#9b9fa2",
                    alignSelf: "center",
                    paddingLeft: "1em",
                    paddingRight: "1em"
                }}
                    onClick={(e) => {
                        e.preventDefault()

                        if (displayEmojis) {
                            setDisplayEmojis(false)

                            return
                        }

                        setDisplayEmojis(true)
                    }}
                >
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="currentColor" d="M9.153 11.603c.795 0 1.439-.879 1.439-1.962s-.644-1.962-1.439-1.962-1.439.879-1.439 1.962.644 1.962 1.439 1.962zm-3.204 1.362c-.026-.307-.131 5.218 6.063 5.551 6.066-.25 6.066-5.551 6.066-5.551-6.078 1.416-12.129 0-12.129 0zm11.363 1.108s-.669 1.959-5.051 1.959c-3.505 0-5.388-1.164-5.607-1.959 0 0 5.912 1.055 10.658 0zM11.804 1.011C5.609 1.011.978 6.033.978 12.228s4.826 10.761 11.021 10.761S23.02 18.423 23.02 12.228c.001-6.195-5.021-11.217-11.216-11.217zM12 21.354c-5.273 0-9.381-3.886-9.381-9.159s3.942-9.548 9.215-9.548 9.548 4.275 9.548 9.548c-.001 5.272-4.109 9.159-9.382 9.159zm3.108-9.751c.795 0 1.439-.879 1.439-1.962s-.644-1.962-1.439-1.962-1.439.879-1.439 1.962.644 1.962 1.439 1.962z"></path></svg>
                </div>
                <div style={{
                    flexGrow: 1,
                    alignSelf: "center",
                    paddingRight: "1em",
                    paddingLeft: "1em"
                }}>
                    <input style={{ 
                        width: "100%",
                        backgroundColor: "#33383b",
                        border: "none",
                        height: "3em",
                        color: "white",
                        outline: "none",
                        fontSize: "15px",
                        paddingLeft: "5px",
                        paddingRight: "5px"
                    }}
                    onFocus={() => {
                        setDisplayEmojis(false)
                    }}
                    value={currentMessage} onChange={(e: any) => {
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
                    />
                </div>
                <div style={{
                    alignSelf: "center",
                    paddingLeft: "1em",
                    paddingRight: "1em",
                    color: "#828689"
                }}
                onClick={() => {
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
                }}
                >
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="currentColor" d="M1.101 21.757L23.8 12.028 1.101 2.3l.011 7.912 13.623 1.816-13.623 1.817-.011 7.912z"></path></svg>
                </div>
            </div>
            {displayEmojis &&
            <div>
                <div style={{
                    backgroundColor: "#2a2f32",
                    fontSize: "35px"
                }}>
                😀 😃 😄 😁 😆 😅 😂 🤣 🥲 ☺️ 😊 😇 🙂 🙃 😉 😌 😍 🥰 😘 😗 😙 😚 😋 😛 😝 😜 🤪 🤨 🧐 🤓 😎 🥸 🤩 🥳 😏 😒 😞 😔 😟 😕 🙁 ☹️ 😣 😖 😫 😩 🥺 😢 😭 😤 😠 😡 🤬 🤯 😳 🥵 🥶
                </div>
            </div>}
        </div>
    )
}

export default Newchatroom