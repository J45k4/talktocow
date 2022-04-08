import { Fragment } from 'react'
import Head from 'next/head'

import '../styles/global.css'
import { startNotificationHandler } from '../src/logic/notification-handler'

if (typeof window !== "undefined") {
    if ("serviceWorker" in navigator) {
        window.addEventListener("load", () => {
            navigator.serviceWorker.register("sw.js")
        })
    }

    startNotificationHandler()
}


export default function App({ Component, pageProps }) {
    return (
        <Fragment>
            <Head>
                <link rel="manifest" href="/manifest.json" />
                <title>Talktocow 🥰 </title>

                <link rel="preconnect" href="https://fonts.googleapis.com" />
                <link rel="preconnect" href="https://fonts.gstatic.com" />
                <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@700&display=swap" rel="stylesheet"></link>
                <meta name="viewport" content="initial-scale=1.0, maximum-scale=1.0, width=device-width, user-scalable=no" />
            </Head>
            <Component {...pageProps} />
        </Fragment>
    )
}
