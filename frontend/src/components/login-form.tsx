import { useCallback, useEffect } from "react";
import { useState } from "react";
import { browserSupportsWebAuthn, startAuthentication } from "@simplewebauthn/browser";
import { setSession } from "../logic/session-manager";
import { postJson } from "../api-methods";
import { ws } from "../ws";

export const LoginForm = () => {
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();

    const [currentlyLogining, setCurrentlyLogining] = useState(false);
    const [currentlyPasskeyLogining, setCurrentlyPasskeyLogining] = useState(false);
    const [loginError, setLoginError] = useState("");
    const [passkeysSupported, setPasskeysSupported] = useState(true);

	useEffect(() => {
		setPasskeysSupported(browserSupportsWebAuthn())
	}, [])

	const login = useCallback(async () => {
		setCurrentlyLogining(true);
		setLoginError("");

		try {
			const res = await postJson<any>("/api/login", {
				username: username,
				password: password
			})

			if (res.error) {
				setCurrentlyLogining(false)
				setLoginError(res.error.message)

				return
			}

			ws.openConn(res.payload.token)
			
			setSession({
				token: res.payload.token,
				userId: res.payload.userId,
				username: res.payload.username,
				authMethod: "password"
			})
		} catch (e) {
			setLoginError("Unknown login error")
			setCurrentlyLogining(false)
		}
	}, [password, username, setCurrentlyLogining, setLoginError])

	const loginWithPasskey = useCallback(async () => {
		if (!passkeysSupported) {
			setLoginError("This browser does not support passkeys")
			return
		}

		setCurrentlyPasskeyLogining(true)
		setLoginError("")

		try {
			const begin = await postJson<any>("/api/passkeys/login/begin", {})

			if (begin.error) {
				setLoginError(begin.error.message)
				setCurrentlyPasskeyLogining(false)
				return
			}

			const assertion = await startAuthentication({
				optionsJSON: begin.payload.options
			})

			const finish = await postJson<any>("/api/passkeys/login/finish", {
				ceremonyId: begin.payload.ceremonyId,
				response: assertion
			})

			if (finish.error) {
				setLoginError(finish.error.message)
				setCurrentlyPasskeyLogining(false)
				return
			}

			ws.openConn(finish.payload.token)

			setSession({
				token: finish.payload.token,
				userId: finish.payload.userId,
				username: finish.payload.username,
				authMethod: "passkey"
			})
		} catch (e) {
			setLoginError(e instanceof Error ? e.message : "Unknown passkey login error")
			setCurrentlyPasskeyLogining(false)
		}
	}, [passkeysSupported])

    return (
        <div>
            <div>
                <div>
                    <label>
                        Username
					</label>
                </div>
                <input type="text" value={username} onChange={(e: any) => {
                    setUsername(e.target.value)
                }} onKeyDown={async (e) => {
					if (e.key === "Enter") {
						await login()
					}
				}} />
            </div>
            <div>
                <div>
                    <label>
                        Password
					</label>
                </div>
                <input type="password" value={password} onChange={(e: any) => {
                    setPassword(e.target.value)
                }} onKeyDown={async (e) => {
					if (e.key === "Enter") {
						await login()
					}
				}} />
            </div>

            <div>
                <div style={{
                    color: "red"
                }}>
                    {loginError}
                </div>
                {currentlyLogining ?
                    <div>
                        Logining...
								</div>
                    :
                    <button onClick={login}>
                        Login
					</button>}
				<button onClick={loginWithPasskey} disabled={currentlyPasskeyLogining || !passkeysSupported}>
					{currentlyPasskeyLogining ? "Waiting for passkey..." : "Sign in with passkey"}
				</button>
				{passkeysSupported ? null : (
					<div>
						This browser does not support passkeys.
					</div>
				)}
            </div>
        </div>
    )
}
