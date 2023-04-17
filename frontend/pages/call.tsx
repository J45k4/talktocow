import React, { useEffect } from "react";
import { ws } from "../src/ws";

export default function CallPage() {
	useEffect(() => {
		const peer = new RTCPeerConnection()

		const createOffer = async () => {
			const offerRes = await peer.createOffer()

			ws.send({
				type: "createOffer",
				sdp: offerRes.sdp
			})
		}

		peer.onicecandidate = (e) => {
			
		}

		 peer.createOffer()
	}, [])

	return <div>Call</div>;
}