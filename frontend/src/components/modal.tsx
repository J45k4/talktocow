import React, { ReactNode, useEffect } from "react"
import styles from "./modal.module.css"

type ModalProps = {
    isOpen: boolean
    title: string
    children: ReactNode
    onClose: () => void
}

export const Modal = (props: ModalProps) => {
    useEffect(() => {
        if (!props.isOpen) {
            return
        }

        const onKeyDown = (event: KeyboardEvent) => {
            if (event.key === "Escape") {
                props.onClose()
            }
        }

        document.addEventListener("keydown", onKeyDown)

        return () => {
            document.removeEventListener("keydown", onKeyDown)
        }
    }, [props.isOpen, props.onClose])

    if (!props.isOpen) {
        return null
    }

    return (
        <div className={styles.backdrop} onClick={props.onClose}>
            <div className={styles.modal} role="dialog" aria-modal="true" aria-label={props.title} onClick={e => e.stopPropagation()}>
                <div className={styles.header}>
                    <h2 className={styles.title}>{props.title}</h2>
                    <button className={styles.closeButton} type="button" aria-label="Close" onClick={props.onClose}>
                        ×
                    </button>
                </div>
                <div className={styles.content}>
                    {props.children}
                </div>
            </div>
        </div>
    )
}
