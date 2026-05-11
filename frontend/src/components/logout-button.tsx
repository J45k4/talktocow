import { logout } from "../logic/session-manager"

export const LogoutButton = () => {
    return (
        <button onClick={() => {
            logout()
        }}>
            Log out
        </button>
    )
}
