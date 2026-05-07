import React, { useEffect, useState } from "react"
import { PageContainer } from "../src/components/page-container"
import { deleteJson, getJson, postJson } from "../src/api-methods"
import { getSession } from "../src/logic/session-manager"

type ApiKey = {
	id: number
	name: string
	prefix: string
	createdAt: string
	lastUsedAt?: string | null
	revokedAt?: string | null
}

type CreatedApiKey = ApiKey & {
	token: string
}

const formatDate = (value?: string | null) => {
	if (!value) {
		return "-"
	}

	return new Date(value).toLocaleString()
}

export default function ProfilePage() {
	const session = getSession()
	const [apiKeys, setApiKeys] = useState<ApiKey[]>([])
	const [name, setName] = useState("")
	const [createdToken, setCreatedToken] = useState("")
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState("")

	const loadApiKeys = async () => {
		const response = await getJson<ApiKey[]>("/api/api-keys")
		if (response.error) {
			setError(response.error.message)
			return
		}

		setApiKeys(response.payload || [])
	}

	useEffect(() => {
		loadApiKeys()
	}, [])

	const createApiKey = async (event: React.FormEvent) => {
		event.preventDefault()
		setLoading(true)
		setError("")
		setCreatedToken("")

		const response = await postJson<CreatedApiKey>("/api/api-keys", { name })
		setLoading(false)

		if (response.error) {
			setError(response.error.message)
			return
		}

		if (response.payload) {
			setCreatedToken(response.payload.token)
			setName("")
			await loadApiKeys()
		}
	}

	const revokeApiKey = async (apiKey: ApiKey) => {
		if (!window.confirm(`Revoke API key '${apiKey.name}'? This cannot be undone.`)) {
			return
		}

		const response = await deleteJson(`/api/api-keys/${apiKey.id}`)
		if (response.error) {
			setError(response.error.message)
			return
		}

		await loadApiKeys()
	}

	const copyCreatedToken = async () => {
		await navigator.clipboard.writeText(createdToken)
	}

	return (
		<PageContainer>
			<h1>Profile</h1>
			<p>
				Logged in as <strong>{session.username}</strong>
			</p>

			<h2>API keys</h2>
			<p>
				Use API keys for CLI tools and scripts. The full key is shown only once when you create it.
			</p>

			{error ? (
				<div style={{ color: "red", marginBottom: "12px" }}>{error}</div>
			) : null}

			{createdToken ? (
				<div style={{ border: "1px solid #444", padding: "12px", marginBottom: "16px" }}>
					<strong>Copy your new API key now. It will not be shown again.</strong>
					<pre style={{ whiteSpace: "pre-wrap", wordBreak: "break-all" }}>{createdToken}</pre>
					<button onClick={copyCreatedToken}>Copy</button>
				</div>
			) : null}

			<form onSubmit={createApiKey} style={{ display: "flex", gap: "8px", marginBottom: "20px" }}>
				<input
					value={name}
					onChange={(event) => setName(event.target.value)}
					placeholder="API key name"
					maxLength={100}
				/>
				<button disabled={loading || name.trim() === ""} type="submit">
					{loading ? "Creating..." : "Create API key"}
				</button>
			</form>

			<table>
				<thead>
					<tr>
						<th>Name</th>
						<th>Prefix</th>
						<th>Created</th>
						<th>Last used</th>
						<th>Status</th>
						<th></th>
					</tr>
				</thead>
				<tbody>
					{apiKeys.map((apiKey) => (
						<tr key={apiKey.id}>
							<td>{apiKey.name}</td>
							<td>{apiKey.prefix}</td>
							<td>{formatDate(apiKey.createdAt)}</td>
							<td>{formatDate(apiKey.lastUsedAt)}</td>
							<td>{apiKey.revokedAt ? `Revoked ${formatDate(apiKey.revokedAt)}` : "Active"}</td>
							<td>
								{apiKey.revokedAt ? null : (
									<button onClick={() => revokeApiKey(apiKey)}>Revoke</button>
								)}
							</td>
						</tr>
					))}
				</tbody>
			</table>
		</PageContainer>
	)
}
