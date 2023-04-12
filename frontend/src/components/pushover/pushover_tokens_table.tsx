import React from "react"
import { useGetData } from "../../hokers"

export const PushoverTokensTable = () => {
	const [tokens, error] = useGetData<any>("/api/pushovertokens", [])

	if (error) {
		return <div>
			{error.message}
		</div>
	}
	
	return (
		<table>
			<thead>
				<tr>
					<th>
						Token
					</th>
					<th>
						User token
					</th>
					<th>
						User name
					</th>
				</tr>
			</thead>
			<tbody>
				
			</tbody>
		</table>
	)
}