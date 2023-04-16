import { Fragment } from 'react'
import Head from 'next/head'

import '../styles/globals.css'
import { startNotificationHandler } from '../src/logic/notification-handler'

if (typeof window !== "undefined") {
    // if ("serviceWorker" in navigator) {
    //     window.addEventListener("load", () => {
    //         navigator.serviceWorker.register("sw.js")
    //     })
    // }

    startNotificationHandler()
}

const isServer = () => typeof window === 'undefined';

export default function App({ Component, pageProps }: { Component: any, pageProps: any }) {
    return (
        <Fragment>
            <Head>
                <title>Talktocow</title>
				<meta name="viewport" content="initial-scale=1.0, maximum-scale=1.0, width=device-width, user-scalable=no" />
            </Head>
            <Component {...pageProps} />
        </Fragment>
    )
}
