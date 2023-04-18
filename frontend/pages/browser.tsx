import React from "react";

export default function BrowserInfoPage() {
	const [devices, setDevices] = React.useState([]);
	const [notificationsAvailable, setNotificationsAvailable] = React.useState(false);
	const [pushManagerAvailable, setPushManagerAvailable] = React.useState(false);
	const [rtcPeerConnectionAvailable, setRtcPeerConnectionAvailable] = React.useState(false);
	const [userAgent, setUserAgent] = React.useState("");
	const [cameraPermission, setCameraPermission] = React.useState("");
	const [microphonePermission, setMicrophonePermission] = React.useState("");
	const [screenSharingSupported, setScreenSharingSupported] = React.useState(false);
	const [screenSharingPermission, setScreenSharingPermission] = React.useState("");

	React.useEffect(() => {
		if (!navigator.mediaDevices?.enumerateDevices) {
			console.log("enumerateDevices() not supported.");
		} else {
			// List cameras and microphones.
			navigator.mediaDevices
				.enumerateDevices()
				.then((devices) => {
					const deviceInfos = []

					devices.forEach((device) => {
						console.log(`${device.kind}: ${device.label} id = ${device.deviceId}`);
						
						deviceInfos.push({
							kind: device.kind,
							label: device.label,
							deviceId: device.deviceId
						})
					});

					setDevices(deviceInfos)
				})
				.catch((err) => {
					console.error(`${err.name}: ${err.message}`);
				});
		}

		// window.RTCPeerConnection

		if ("RTCPeerConnection" in window) {
			setRtcPeerConnectionAvailable(true)
		}

		if ("Notification" in window) {
			setNotificationsAvailable(true)
		}

		if ("PushManager" in window) {
			setPushManagerAvailable(true)
		}

		if ("permissions" in navigator) {
			navigator.permissions.query({ name: "camera" } as any).then((result) => {
				setCameraPermission(result.state)
			})

			navigator.permissions.query({ name: "microphone" } as any).then((result) => {
				setMicrophonePermission(result.state)
			})

			navigator.permissions.query({ name: "display-capture" } as any).then((result) => {
				setScreenSharingPermission(result.state)
			})
		}

		setUserAgent(navigator.userAgent)

		if ("mediaDevices" in navigator && "getDisplayMedia" in navigator.mediaDevices) {
			setScreenSharingSupported(true)
		}
	}, [setNotificationsAvailable]);

	return (
		<div>
			<h1>Browser info</h1>

			<h2>Devices</h2>
			<table>
				<thead style={{
				border: "1px solid black"
			}}>
					<tr>
						<th>Kind</th>
						<th>Label</th>
						<th>Device ID</th>
					</tr>
				</thead>
				<tbody>
					{devices.map((device, i) => {
						return (
							<tr key={i}>
								<td>{device.kind}</td>
								<td>{device.label}</td>
								<td>{device.deviceId}</td>
							</tr>
						)
					})}
				</tbody>
			</table>

			<h2>Other</h2>
			<table>
				<tbody>
					<tr>
						<td>
							RTC Peer Connection
						</td>
						<td>
							{rtcPeerConnectionAvailable ? "Yes" : "No"}
						</td>
					</tr>
					<tr>
						<td>
							Notifications
						</td>
						<td>
							{notificationsAvailable ? "Yes" : "No"}
						</td>
					</tr>
					<tr>
						<td>
							Push Manager
						</td>
						<td>
							{pushManagerAvailable ? "Yes" : "No"}
						</td>
					</tr>
					<tr>
						<td>
							User Agent
						</td>
						<td>
							{userAgent}
						</td>
					</tr>
					<tr>
						<td>
							Camera permission
						</td>
						<td>
							{cameraPermission}
						</td>
					</tr>
					<tr>
						<td>
							Microphone permission
						</td>
						<td>
							{microphonePermission}
						</td>
					</tr>
					<tr>
						<td>
							Screen sharing supported
						</td>
						<td>
							{screenSharingSupported ? "Yes" : "No"}
						</td>
					</tr>
					<tr>
						<td>
							Screen sharing permission
						</td>
						<td>
							{screenSharingPermission}
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	)
}