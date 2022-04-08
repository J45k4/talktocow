import { connEvent } from "./websocket-conn";

export const startNotificationHandler = () => {
    if (Notification) {
        Notification.requestPermission((status) => {
            console.log("status", status);
    
            if (status === "granted") {
                connEvent.subscribe({
                    next: v => {
                        console.log("v", v);
    
                        if (v.type === "messageFromServer") {
                            if (v.newChatroomMessage) {
                                new Notification("New message")
                            }
                        }
                    }
                })
            }
        })
    }
}