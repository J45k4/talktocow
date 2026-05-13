package routes

import (
	"errors"
	"testing"
	"time"
)

func TestValidateDiaryEntryDateEditAllowsEntriesLessThanSevenDaysOld(t *testing.T) {
	now := time.Date(2026, 5, 12, 12, 0, 0, 0, time.UTC)
	existingCreatedAt := now.Add(-(6*24*time.Hour + 23*time.Hour + 59*time.Minute))

	parsed, err := validateDiaryEntryDateEdit(existingCreatedAt, "2026-05-10", now)

	if err != nil {
		t.Fatalf("expected date edit to be valid, got %v", err)
	}

	if parsed.Format("2006-01-02") != "2026-05-10" {
		t.Fatalf("expected parsed date 2026-05-10, got %s", parsed.Format("2006-01-02"))
	}
}

func TestValidateDiaryEntryDateEditRejectsEntriesSevenDaysOld(t *testing.T) {
	now := time.Date(2026, 5, 12, 12, 0, 0, 0, time.UTC)
	existingCreatedAt := now.Add(-7 * 24 * time.Hour)

	_, err := validateDiaryEntryDateEdit(existingCreatedAt, "2026-05-10", now)

	if !errors.Is(err, errDiaryEntryDateEditExpired) {
		t.Fatalf("expected expired date edit error, got %v", err)
	}
}

func TestValidateDiaryEntryDateEditRejectsInvalidDates(t *testing.T) {
	now := time.Date(2026, 5, 12, 12, 0, 0, 0, time.UTC)
	existingCreatedAt := now.Add(-24 * time.Hour)

	_, err := validateDiaryEntryDateEdit(existingCreatedAt, "not-a-date", now)

	if !errors.Is(err, errInvalidDiaryEntryDate) {
		t.Fatalf("expected invalid date error, got %v", err)
	}
}
