import React, { useEffect } from "react";
import { ws } from "../src/ws";

export default function CallPage() {
	useEffect(() => {
		const peer = new RTCPeerConnection()

		const createOffer = async () => {
			const offerRes = await peer.createOffer()

			ws.send({
				type: "createWebRTCOffer",
				sdp: offerRes.sdp || "",
				userId: "",
				deviceId: ""
			})
		}

		peer.onicecandidate = (e) => {
			
		}

		 peer.createOffer()
	}, [])

	return <div>Call</div>;
}
