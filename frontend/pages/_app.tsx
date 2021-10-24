import { Fragment } from 'react'
import Head from 'next/head'

// import '../styles/global.css'

if (typeof window !== "undefined") {
    if ("serviceWorker" in navigator) {
        window.addEventListener("load", () => {
            navigator.serviceWorker.register("sw.js")
        })
    }
}


export default function App({ Component, pageProps }) {
    return (
        <Fragment>
            <Head>
                <link rel="manifest" href="/manifest.json" />
                <title>Talktocow 🥰 </title>
                <meta name="viewport" content="initial-scale=1.0, maximum-scale=1.0, width=device-width, user-scalable=no" />
            </Head>
            <Component {...pageProps} />
        </Fragment>
    )
}
