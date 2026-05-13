const diaryDateEditWindowMs = 7 * 24 * 60 * 60 * 1000

export const dateInputValueFromDiaryDate = (value?: string | null) => {
    if (!value) {
        return ""
    }

    const date = new Date(value)

    if (Number.isNaN(date.getTime())) {
        return ""
    }

    return date.toISOString().slice(0, 10)
}

export const createdAtFromDiaryDateInputValue = (value: string) => {
    return value ? `${value}T12:00:00Z` : undefined
}

export const isDiaryDateEditable = (createdAt?: string | null, now: Date = new Date()) => {
    if (!createdAt) {
        return false
    }

    const date = new Date(createdAt)

    if (Number.isNaN(date.getTime())) {
        return false
    }

    return now.getTime() - date.getTime() < diaryDateEditWindowMs
}
