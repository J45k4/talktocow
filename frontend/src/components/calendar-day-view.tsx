import React, { useCallback, useEffect, useMemo, useState } from "react"
import { Link, useSearchParams } from "react-router-dom"
import { deleteJson, getJson, postJson } from "../api-methods"
import { Modal } from "./modal"
import styles from "./calendar-day-view.module.css"

type DiaryCalendarEntry = {
    id: number
    title: string
    body: string
    createdAt: string
    label?: string
    startsAt?: string
    endsAt?: string
    allDay: boolean
}

const sameDay = (a: Date, b: Date) => (
    a.getFullYear() === b.getFullYear() &&
    a.getMonth() === b.getMonth() &&
    a.getDate() === b.getDate()
)

const dayKey = (date: Date) => [date.getFullYear(), date.getMonth(), date.getDate()].join("-")

const toDateInputValue = (date: Date) => {
    const month = `${date.getMonth() + 1}`.padStart(2, "0")
    const day = `${date.getDate()}`.padStart(2, "0")
    return `${date.getFullYear()}-${month}-${day}`
}

const monthName = (date: Date) => date.toLocaleDateString("en-US", { month: "long", year: "numeric" })

const getCalendarDays = (selectedDate: Date) => {
    const firstDay = new Date(selectedDate.getFullYear(), selectedDate.getMonth(), 1)
    const start = new Date(firstDay)
    start.setDate(firstDay.getDate() - firstDay.getDay())

    return Array.from({ length: 42 }, (_, index) => {
        const date = new Date(start)
        date.setDate(start.getDate() + index)
        return date
    })
}

const entryHappensOnDay = (entry: DiaryCalendarEntry, day: Date) => {
    const start = new Date(entry.startsAt ?? entry.createdAt)
    const end = entry.endsAt ? new Date(entry.endsAt) : start
    const dayStart = new Date(day.getFullYear(), day.getMonth(), day.getDate())
    const dayEnd = new Date(dayStart)
    dayEnd.setDate(dayEnd.getDate() + 1)

    return start < dayEnd && end >= dayStart
}

const labelColorClasses = [
    styles.labelColorBlue,
    styles.labelColorGreen,
    styles.labelColorOrange,
    styles.labelColorPurple,
    styles.labelColorPink,
    styles.labelColorCyan,
    styles.labelColorYellow,
    styles.labelColorRed,
    styles.labelColorBrown,
    styles.labelColorGray
]

const getLabelColorClass = (label: string) => {
    let hash = 5381

    for (const char of label.toLowerCase()) {
        hash = ((hash << 5) + hash) ^ char.charCodeAt(0)
    }

    return labelColorClasses[Math.abs(hash) % labelColorClasses.length]
}

const getUniqueLabels = (entries: DiaryCalendarEntry[]) => {
    const labels: string[] = []

    for (const entry of entries) {
        if (entry.label && !labels.includes(entry.label)) {
            labels.push(entry.label)
        }
    }

    return labels
}

const getPreview = (body: string) => {
    if (!body) {
        return "No details yet."
    }

    return body.length > 120 ? `${body.slice(0, 120)}…` : body
}

const formatEntryTime = (entry: DiaryCalendarEntry) => {
    if (entry.allDay || !entry.startsAt) {
        return "All day"
    }

    const start = new Date(entry.startsAt).toLocaleTimeString("en-US", { hour: "2-digit", minute: "2-digit" })

    if (!entry.endsAt) {
        return start
    }

    const end = new Date(entry.endsAt).toLocaleTimeString("en-US", { hour: "2-digit", minute: "2-digit" })
    return `${start}–${end}`
}

export const CalendarDayView = () => {
    const [searchParams] = useSearchParams()
    const selectedDateParam = searchParams.get("date")
    const [selectedDate, setSelectedDate] = useState(selectedDateParam ? new Date(`${selectedDateParam}T12:00`) : new Date())
    const [entries, setEntries] = useState<DiaryCalendarEntry[]>([])
    const [labels, setLabels] = useState<string[]>([])
    const [quickLabel, setQuickLabel] = useState("")
    const [isLoading, setIsLoading] = useState(true)
    const [isAddingLabel, setIsAddingLabel] = useState(false)
    const [entryToDelete, setEntryToDelete] = useState<DiaryCalendarEntry | null>(null)
    const [deleteError, setDeleteError] = useState("")

    const fetchCalendar = useCallback(() => {
        setIsLoading(true)
        Promise.all([
            getJson<DiaryCalendarEntry[]>("/api/diary/entries?offset=0&limit=300"),
            getJson<string[]>("/api/diary/labels")
        ]).then(([entriesResponse, labelsResponse]) => {
            setEntries(entriesResponse.payload ?? [])
            setLabels(labelsResponse.payload ?? [])
        }).finally(() => {
            setIsLoading(false)
        })
    }, [])

    useEffect(() => {
        fetchCalendar()
    }, [fetchCalendar])

    const calendarDays = useMemo(() => getCalendarDays(selectedDate), [selectedDate])

    const entriesByDate = useMemo(() => {
        const map = new Map<string, DiaryCalendarEntry[]>()

        for (const day of calendarDays) {
            const dayEntries = entries.filter(entry => entryHappensOnDay(entry, day))

            if (dayEntries.length > 0) {
                map.set(dayKey(day), dayEntries)
            }
        }

        return map
    }, [calendarDays, entries])

    const selectedEntries = useMemo(() => {
        return entries.filter(entry => entryHappensOnDay(entry, selectedDate))
    }, [entries, selectedDate])

    const monthlyLabelCounts = useMemo(() => {
        const counts = new Map<string, number>()

        for (const entry of entries) {
            if (!entry.label) {
                continue
            }

            const entryDate = new Date(entry.createdAt)

            if (entryDate.getFullYear() !== selectedDate.getFullYear() || entryDate.getMonth() !== selectedDate.getMonth()) {
                continue
            }

            counts.set(entry.label, (counts.get(entry.label) ?? 0) + 1)
        }

        return [...counts.entries()].sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0]))
    }, [entries, selectedDate])

    const moveMonth = (amount: number) => {
        setSelectedDate(current => new Date(current.getFullYear(), current.getMonth() + amount, Math.min(current.getDate(), 28)))
    }

    const deleteEntry = (entry: DiaryCalendarEntry) => {
        deleteJson(`/api/diary/entry/${entry.id}`).then(r => {
            if (r.error && r.error.code !== 404 && r.error.message.toLowerCase() !== "not found") {
                setDeleteError(r.error.message)
                return
            }

            setEntries(currentEntries => currentEntries.filter(currentEntry => currentEntry.id !== entry.id))
            setEntryToDelete(null)
        })
    }

    const addLabel = (label: string) => {
        const cleanLabel = label.trim()

        if (!cleanLabel) {
            return
        }

        setIsAddingLabel(true)
        postJson<DiaryCalendarEntry>("/api/diary/entry", {
            title: cleanLabel,
            body: "",
            label: cleanLabel,
            createdAt: `${toDateInputValue(selectedDate)}T12:00:00Z`,
            allDay: true
        }).then(r => {
            if (r.payload) {
                setQuickLabel("")
                fetchCalendar()
            }
        }).finally(() => {
            setIsAddingLabel(false)
        })
    }

    return (
        <section className={styles.calendarCard}>
            <div className={styles.calendarHeader}>
                <div>
                    <div className={styles.eyebrow}>Diary calendar</div>
                    <h2>{selectedDate.toLocaleDateString("en-US", { weekday: "long", month: "long", day: "numeric" })}</h2>
                </div>
                <div className={styles.headerActions}>
                    <Link className={styles.newEventButton} to={`/diary/new?date=${toDateInputValue(selectedDate)}`}>+ Full entry</Link>
                    <div className={styles.monthControls}>
                        <button onClick={() => moveMonth(-1)} aria-label="Previous month">‹</button>
                        <span>{monthName(selectedDate)}</span>
                        <button onClick={() => moveMonth(1)} aria-label="Next month">›</button>
                    </div>
                </div>
            </div>

            <div className={styles.contentGrid}>
                <div className={styles.monthPanel}>
                    <div className={styles.overviewBox}>
                        <div className={styles.quickAddTitle}>Label overview this month</div>
                        {monthlyLabelCounts.length === 0 ? (
                            <div className={styles.overviewEmpty}>No labels this month yet.</div>
                        ) : (
                            <div className={styles.overviewList}>
                                {monthlyLabelCounts.map(([label, count]) => (
                                    <div className={styles.overviewRow} key={label}>
                                        <span className={`${styles.overviewLabel} ${getLabelColorClass(label)}`}>{label}</span>
                                        <span className={styles.overviewCount}>{count}</span>
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>
                    <div className={styles.weekdays}>
                        {['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'].map(day => (
                            <div key={day}>{day}</div>
                        ))}
                    </div>
                    <div className={styles.daysGrid}>
                        {calendarDays.map(date => {
                            const dateEntries = entriesByDate.get(dayKey(date)) ?? []
                            const isSelected = sameDay(date, selectedDate)
                            const isToday = sameDay(date, new Date())
                            const isOtherMonth = date.getMonth() !== selectedDate.getMonth()

                            return (
                                <button
                                    key={date.toISOString()}
                                    className={`${styles.dayButton} ${isSelected ? styles.selectedDay : ""} ${isToday ? styles.today : ""} ${isOtherMonth ? styles.otherMonth : ""}`}
                                    onClick={() => setSelectedDate(date)}
                                >
                                    <span className={styles.dayNumber}>{date.getDate()}</span>
                                    {getUniqueLabels(dateEntries).length > 0 && (
                                        <span className={styles.dayLabels}>
                                            {getUniqueLabels(dateEntries).slice(0, 3).map(label => (
                                                <span className={`${styles.dayLabelChip} ${getLabelColorClass(label)}`} key={label}>{label}</span>
                                            ))}
                                            {getUniqueLabels(dateEntries).length > 3 && (
                                                <span className={styles.moreLabels}>+{getUniqueLabels(dateEntries).length - 3}</span>
                                            )}
                                        </span>
                                    )}
                                </button>
                            )
                        })}
                    </div>
                </div>

                <div className={styles.dayPanel}>
                    <div className={styles.dayPanelHeader}>
                        <h3>What happened this day</h3>
                        <span>{selectedEntries.length} item{selectedEntries.length === 1 ? "" : "s"}</span>
                    </div>

                    <div className={styles.quickAddBox}>
                        <div className={styles.quickAddTitle}>Quick labels</div>
                        <div className={styles.labelChips}>
                            {labels.map(label => (
                                <button className={getLabelColorClass(label)} key={label} onClick={() => addLabel(label)} disabled={isAddingLabel}>{label}</button>
                            ))}
                        </div>
                        <div className={styles.quickAddInputRow}>
                            <input
                                value={quickLabel}
                                placeholder="New label, e.g. Sauna day"
                                onChange={e => setQuickLabel(e.target.value)}
                                onKeyDown={e => {
                                    if (e.key === "Enter") {
                                        addLabel(quickLabel)
                                    }
                                }}
                            />
                            <button onClick={() => addLabel(quickLabel)} disabled={isAddingLabel}>Add</button>
                        </div>
                    </div>

                    {isLoading && <div className={styles.emptyState}>Loading calendar…</div>}

                    {!isLoading && selectedEntries.length === 0 && (
                        <div className={styles.emptyState}>Nothing tracked for this day yet.</div>
                    )}

                    {!isLoading && selectedEntries.map(entry => (
                        <div className={styles.eventItem} key={entry.id}>
                            <Link className={styles.eventItemLink} to={`/diary/entry/${entry.id}`}>
                                <div className={styles.eventTime}>
                                    {entry.label ? <span className={`${styles.entryLabel} ${getLabelColorClass(entry.label)}`}>{entry.label}</span> : formatEntryTime(entry)}
                                </div>
                                <div>
                                    <div className={styles.eventTitle}>{entry.title}</div>
                                    <div className={styles.eventBody}>{getPreview(entry.body)}</div>
                                </div>
                            </Link>
                            <button className={styles.deleteEntryButton} onClick={() => setEntryToDelete(entry)} aria-label={`Remove ${entry.title}`}>
                                ×
                            </button>
                        </div>
                    ))}
                </div>
            </div>
            <Modal isOpen={entryToDelete != null} title="Remove diary entry?" onClose={() => setEntryToDelete(null)}>
                <div className={styles.confirmModalContent}>
                    <p>Remove “{entryToDelete?.title}” from this day?</p>
                    <div className={styles.modalActions}>
                        <button className={styles.cancelButton} onClick={() => setEntryToDelete(null)}>Cancel</button>
                        <button className={styles.dangerButton} onClick={() => entryToDelete && deleteEntry(entryToDelete)}>Remove</button>
                    </div>
                </div>
            </Modal>
            <Modal isOpen={deleteError !== ""} title="Could not remove entry" onClose={() => setDeleteError("")}>
                <div className={styles.confirmModalContent}>
                    <p>{deleteError}</p>
                    <div className={styles.modalActions}>
                        <button className={styles.cancelButton} onClick={() => setDeleteError("")}>OK</button>
                    </div>
                </div>
            </Modal>
        </section>
    )
}
