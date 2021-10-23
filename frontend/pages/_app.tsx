import '../styles/global.css'

if (typeof window !== "undefined") {
  if ("serviceWorker" in navigator) {
    window.addEventListener("load", () => {
      navigator.serviceWorker.register("sw.js")
    })
  }
}


export default function App({ Component, pageProps }) {
  return <Component {...pageProps} />
}
