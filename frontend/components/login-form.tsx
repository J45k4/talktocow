import { useState } from "react";
import { setSession } from "../logic/session-manager";
import { postJson } from "../utility/talktocow-api-helpers";

export const LoginForm = () => {
    const [username, setUsername] = useState();
    const [password, setPassword] = useState();

    const [currentlyLogining, setCurrentlyLogining] = useState(false);
    const [loginError, setLoginError] = useState("");

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
                    <button onClick={async () => {
                        setCurrentlyLogining(true);

                        const res = await postJson("/api/login", {
                            username: username,
                            password: password
                        })

                        if (res.error) {
                            setCurrentlyLogining(false)
                            setLoginError(res.error.message)

                            return
                        }
                        
                        setSession({
                            token: res.payload.token,
                            userId: res.payload.userId,
                            username: res.payload.username
                        })

                    }}>
                        Login
					</button>}
            </div>
        </div>
    )
}