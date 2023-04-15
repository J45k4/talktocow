import React, { useEffect, useRef, useState, CSSProperties } from 'react';
import { createLogger } from '../logger';

const logger = createLogger("VideoCapture")

interface VideoCaptureProps {
	constraints: MediaStreamConstraints;
	style?: CSSProperties
}

export const VideoCapture: React.FC<VideoCaptureProps> = ({
	constraints,
	style
}) => {
	const [mediaStream, setMediaStream] = useState<MediaStream | null>(null);
	const videoRef = useRef<HTMLVideoElement>(null);

	useEffect(() => {
		navigator.mediaDevices.getUserMedia(constraints)
			.then(mediaStream => {
				setMediaStream(mediaStream);
			})
			.catch(error => {
				// Handle errors
			});
	}, [constraints]);

	useEffect(() => {
		if (videoRef.current && mediaStream) {
			videoRef.current.srcObject = mediaStream;
			videoRef.current.play();

			return () => {
				logger.info("Cleaning up video capture");
				if (videoRef.current) {
					videoRef.current.srcObject = null;
				}
				if (mediaStream) {
					logger.info("Cleaning up media stream");
					mediaStream.getTracks().forEach(track => track.stop());
				}
			};
		}
	}, [videoRef, mediaStream]);

	return (
		<div style={style}>
			{mediaStream && (
				<video
					ref={videoRef}
					autoPlay
					style={{
						width: "100%",
					}}
				/>
			)}
		</div>
	);
};