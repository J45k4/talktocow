import { createdAtFromDiaryDateInputValue, dateInputValueFromDiaryDate, isDiaryDateEditable } from "../src/components/diary/diary-entry-date"

it("formats diary timestamps for date inputs", () => {
    expect(dateInputValueFromDiaryDate("2026-05-10T12:00:00Z")).toBe("2026-05-10")
})

it("creates diary timestamps from date input values", () => {
    expect(createdAtFromDiaryDateInputValue("2026-05-10")).toBe("2026-05-10T12:00:00Z")
})

it("allows editing dates for diary entries less than seven days old", () => {
    const now = new Date("2026-05-12T12:00:00Z")

    expect(isDiaryDateEditable("2026-05-05T12:01:00Z", now)).toBe(true)
    expect(isDiaryDateEditable("2026-05-05T12:00:00Z", now)).toBe(false)
})
