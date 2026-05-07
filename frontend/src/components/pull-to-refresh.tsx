import React, { ReactNode, TouchEvent, useState } from "react"
import styles from "./pull-to-refresh.module.css"

type PullToRefreshProps = {
    children: ReactNode
    onRefresh: () => Promise<void> | void
}

const refreshThreshold = 80
const maxPullDistance = 110

export const PullToRefresh = (props: PullToRefreshProps) => {
    const [startY, setStartY] = useState<number | null>(null)
    const [pullDistance, setPullDistance] = useState(0)
    const [isRefreshing, setIsRefreshing] = useState(false)

    const canStartPull = () => window.scrollY <= 0 && !isRefreshing

    const resetPull = () => {
        setStartY(null)
        setPullDistance(0)
    }

    const onTouchStart = (event: TouchEvent<HTMLDivElement>) => {
        if (!canStartPull()) {
            return
        }

        setStartY(event.touches[0].clientY)
    }

    const onTouchMove = (event: TouchEvent<HTMLDivElement>) => {
        if (startY === null || !canStartPull()) {
            return
        }

        const distance = event.touches[0].clientY - startY

        if (distance <= 0) {
            setPullDistance(0)
            return
        }

        setPullDistance(Math.min(distance, maxPullDistance))
    }

    const onTouchEnd = async () => {
        if (startY === null) {
            return
        }

        const shouldRefresh = pullDistance >= refreshThreshold

        resetPull()

        if (!shouldRefresh) {
            return
        }

        setIsRefreshing(true)

        try {
            await props.onRefresh()
        } finally {
            setIsRefreshing(false)
        }
    }

    const indicatorVisible = pullDistance > 10 || isRefreshing
    const readyToRefresh = pullDistance >= refreshThreshold

    return (
        <div
            className={styles.container}
            onTouchStart={onTouchStart}
            onTouchMove={onTouchMove}
            onTouchEnd={onTouchEnd}
            onTouchCancel={resetPull}
        >
            <div className={styles.indicator}>
                <div className={`${styles.indicatorContent} ${indicatorVisible ? styles.visible : ""}`}>
                    {isRefreshing ? "Refreshing..." : readyToRefresh ? "Release to refresh" : "Pull to refresh"}
                </div>
            </div>
            <div style={{ transform: pullDistance ? `translateY(${pullDistance / 2}px)` : undefined }}>
                {props.children}
            </div>
        </div>
    )
}
