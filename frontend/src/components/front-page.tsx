import { CalendarDayView } from "./calendar-day-view"
import { IsLoggedIn, IsNotLoggedIn, useIsLoggedIn } from "./isloggedin"
import { LoginForm } from "./login-form"
import { NavigationBar } from "./navigation-bar"
import { ValentinesDayGift } from "./valentines-day-gift"

export const FrontPage = () => {
    return (
        <div style={{
            position: "absolute",
            top: "0px",
            right: "0px",
            bottom: "0px",
            left: "0px"
        }}>
            <IsLoggedIn>
                <NavigationBar />
                <CalendarDayView />
            </IsLoggedIn>
            <IsNotLoggedIn>
                <LoginForm />
            </IsNotLoggedIn>
        </div>
    )
}
