import React from "react";

export default function BrowserInfoPage() {
	const [devices, setDevices] = React.useState([]);
	const [notificationsAvailable, setNotificationsAvailable] = React.useState(false);
	const [rtcPeerConnectionAvailable, setRtcPeerConnectionAvailable] = React.useState(false);

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
	}, [setNotificationsAvailable]);

	return (
		<div>
			<h1>Browser info</h1>
			{devices.map((device) => {
				return (
					<div key={device.deviceId}>
						<p>{device.kind}</p>
						<p>{device.label}</p>
						<p>{device.deviceId}</p>
					</div>
				)	
			})}

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
							User Agent
						</td>
						<td>
							{navigator.userAgent}
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	)
}