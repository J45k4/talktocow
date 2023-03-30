import { Fragment, useEffect, useState } from "react"
import { getSession, SessionChangeNotify, subscribeToSessionEvents, unsubscribeToSessionEvents } from "../logic/session-manager"

export const useIsLoggedIn = () => {
    const [isLoggedin, setIsloggedin] = useState(true)
    
    useEffect(() => {
		const a = getSession().token != null
		setIsloggedin(a)

        function handle(s: SessionChangeNotify) {
            if (s.token != null) {
                setIsloggedin(true)
            } else {
                setIsloggedin(false)
            }
        }

        subscribeToSessionEvents(handle)

        return () => {
            unsubscribeToSessionEvents(handle)
        }
    }, [, setIsloggedin])

    return isLoggedin
}

export const IsLoggedIn = (props: any) => {
    const isLoggedin = useIsLoggedIn()

    if (isLoggedin) {
        return (
            <Fragment>
                {props.children}
            </Fragment>
        )
    }

    return <Fragment />
}

export const IsNotLoggedIn = (props: {
	children: any
}) => {
    const isLoggedin = useIsLoggedIn()

    if (isLoggedin == false) {
        return <Fragment>
            {props.children}
        </Fragment>
    }

    return <Fragment />
}