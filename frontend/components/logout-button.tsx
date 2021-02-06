import { clearSession } from "../logic/session-manager"

export const LogoutButton = () => {
    return (
        <button onClick={() => {
            clearSession()
        }}>
            Log out
        </button>
    )
}