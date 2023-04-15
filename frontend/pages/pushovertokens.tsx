import React from "react"
import { PageContainer } from "../src/components/page-container"
import { PushoverTokensTable } from "../src/components/pushover/pushover_tokens_table"

export default function PushoverTokensPage() {
	return (
		<PageContainer>
			<h1>Pushover Tokens</h1>
			<PushoverTokensTable />
		</PageContainer>
	)
}