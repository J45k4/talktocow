package models

type UserCourseRolde int

const (
	Student UserCourseRolde = iota
	Teacher
)

type HomeworkSubmissionStatus int

const (
	NotSubmitted HomeworkSubmissionStatus = iota
	Submitted
	Graded
)
