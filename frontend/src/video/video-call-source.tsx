import React, { useEffect, useRef, useState } from "react"
import { createLogger } from "../logger";

const logger = createLogger("")

export const VideoCallSource = () => {
	const [peerConnection, setPeerConnection] = useState<RTCPeerConnection>();
	const localStreamRef = useRef<MediaStream>();
	const videoRef = useRef<HTMLVideoElement>(null);

	useEffect(() => {
		async function setupPeerConnection() {
		  const stream = await navigator.mediaDevices.getUserMedia({ audio: true, video: true });
		  localStreamRef.current = stream;
	
		  const peerConnection = new RTCPeerConnection({
			iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
			
		  });
	
		  peerConnection.onicecandidate = (event) => {
			if (event.candidate) {
			  // send candidate to the remote peer
			}
		  };
	
		  peerConnection.ontrack = (event) => {
			if (videoRef.current) {
			  videoRef.current.srcObject = event.streams[0];
			}
		  };

		  
	
		  stream.getTracks().forEach((track) => {
			peerConnection.addTrack(track, stream);
		  });
	
		  setPeerConnection(peerConnection);
		}
	
		setupPeerConnection();

		return () => {
			peerConnection.close()
		}
	  }, []);

	return (
		<div>
			<video ref={videoRef} autoPlay playsInline muted />
		</div>
	)
}