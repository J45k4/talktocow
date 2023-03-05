import React from "react"
import { BsSearch } from "react-icons/bs"

export const CowGPT = () => {
	return (
		<div style={{
			display: "flex",
			flexDirection: "row",
			flexGrow: 1,
		}}>
			<div>
				asd
			</div>
			<div style={{
				display: "flex",
				flexDirection: "column",
				flexGrow: 1,
			}}>
				<div>
					<div style={{
						border: "solid 1px #8E8E8E",
						flexGrow: 1,
						marginLeft: "1em",
						marginRight: "1em",
						marginBottom: "0.2em",
						fontSize: "1.5em",
						padding: "0.2em",
					}}>
						<BsSearch />
					</div>
					
				</div>
				<div style={{
					display: "flex",
					flexGrow: 1,
					padding: "1em",
				}}>
					<div>
						qwer
					</div>
				</div>
				<div style={{
					display: "flex",

				}}>
					<div style={{
						border: "solid 1px #8E8E8E",
						flexGrow: 1,
						margin: "2em",
						fontSize: "1.5em",
						padding: "0.2em",
					}}>
						<input style={{
							border: "none",
							width: "100%",
							height: "100%",
						}} onKeyDown={e => {
							if (e.key === "Enter") {
								console.log("Enter")
							}
						}} />
					</div>
				</div>
			</div>
		</div>
	)
}