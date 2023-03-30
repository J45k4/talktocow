import { useCallback } from "react";
import { useState } from "react";
import { setSession } from "../logic/session-manager";
import { postJson } from "../utility/talktocow-api-helpers";
import { openConn } from "../ws";

export const LoginForm = () => {
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();

    const [currentlyLogining, setCurrentlyLogining] = useState(false);
    const [loginError, setLoginError] = useState("");

	const login = useCallback(async () => {
		setCurrentlyLogining(true);

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

			openConn(res.payload.token)
			
			setSession({
				token: res.payload.token,
				userId: res.payload.userId,
				username: res.payload.username
			})
		} catch (e) {
			setLoginError("Unknown login error")
			setCurrentlyLogining(false)
		}
	}, [password, username, setCurrentlyLogining, setLoginError])

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
            </div>
        </div>
    )
}