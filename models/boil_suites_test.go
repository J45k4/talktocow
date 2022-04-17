// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEvents)
	t.Run("ChatroomUsers", testChatroomUsers)
	t.Run("Chatrooms", testChatrooms)
	t.Run("CourseUsers", testCourseUsers)
	t.Run("Courses", testCourses)
	t.Run("DiaryEntries", testDiaryEntries)
	t.Run("DiaryEntryComments", testDiaryEntryComments)
	t.Run("Events", testEvents)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionComments)
	t.Run("HomeworkSubmissions", testHomeworkSubmissions)
	t.Run("Homeworks", testHomeworks)
	t.Run("LoginSessions", testLoginSessions)
	t.Run("Messages", testMessages)
	t.Run("NotificationLogs", testNotificationLogs)
	t.Run("PushoverTokens", testPushoverTokens)
	t.Run("SharedDiaryEntries", testSharedDiaryEntries)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEvents)
	t.Run("UserReceivedMessages", testUserReceivedMessages)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsDelete)
	t.Run("ChatroomUsers", testChatroomUsersDelete)
	t.Run("Chatrooms", testChatroomsDelete)
	t.Run("CourseUsers", testCourseUsersDelete)
	t.Run("Courses", testCoursesDelete)
	t.Run("DiaryEntries", testDiaryEntriesDelete)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsDelete)
	t.Run("Events", testEventsDelete)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsDelete)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsDelete)
	t.Run("Homeworks", testHomeworksDelete)
	t.Run("LoginSessions", testLoginSessionsDelete)
	t.Run("Messages", testMessagesDelete)
	t.Run("NotificationLogs", testNotificationLogsDelete)
	t.Run("PushoverTokens", testPushoverTokensDelete)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesDelete)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsDelete)
	t.Run("UserReceivedMessages", testUserReceivedMessagesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsQueryDeleteAll)
	t.Run("ChatroomUsers", testChatroomUsersQueryDeleteAll)
	t.Run("Chatrooms", testChatroomsQueryDeleteAll)
	t.Run("CourseUsers", testCourseUsersQueryDeleteAll)
	t.Run("Courses", testCoursesQueryDeleteAll)
	t.Run("DiaryEntries", testDiaryEntriesQueryDeleteAll)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsQueryDeleteAll)
	t.Run("Events", testEventsQueryDeleteAll)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsQueryDeleteAll)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsQueryDeleteAll)
	t.Run("Homeworks", testHomeworksQueryDeleteAll)
	t.Run("LoginSessions", testLoginSessionsQueryDeleteAll)
	t.Run("Messages", testMessagesQueryDeleteAll)
	t.Run("NotificationLogs", testNotificationLogsQueryDeleteAll)
	t.Run("PushoverTokens", testPushoverTokensQueryDeleteAll)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesQueryDeleteAll)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsQueryDeleteAll)
	t.Run("UserReceivedMessages", testUserReceivedMessagesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsSliceDeleteAll)
	t.Run("ChatroomUsers", testChatroomUsersSliceDeleteAll)
	t.Run("Chatrooms", testChatroomsSliceDeleteAll)
	t.Run("CourseUsers", testCourseUsersSliceDeleteAll)
	t.Run("Courses", testCoursesSliceDeleteAll)
	t.Run("DiaryEntries", testDiaryEntriesSliceDeleteAll)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsSliceDeleteAll)
	t.Run("Events", testEventsSliceDeleteAll)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsSliceDeleteAll)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsSliceDeleteAll)
	t.Run("Homeworks", testHomeworksSliceDeleteAll)
	t.Run("LoginSessions", testLoginSessionsSliceDeleteAll)
	t.Run("Messages", testMessagesSliceDeleteAll)
	t.Run("NotificationLogs", testNotificationLogsSliceDeleteAll)
	t.Run("PushoverTokens", testPushoverTokensSliceDeleteAll)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesSliceDeleteAll)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsSliceDeleteAll)
	t.Run("UserReceivedMessages", testUserReceivedMessagesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsExists)
	t.Run("ChatroomUsers", testChatroomUsersExists)
	t.Run("Chatrooms", testChatroomsExists)
	t.Run("CourseUsers", testCourseUsersExists)
	t.Run("Courses", testCoursesExists)
	t.Run("DiaryEntries", testDiaryEntriesExists)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsExists)
	t.Run("Events", testEventsExists)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsExists)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsExists)
	t.Run("Homeworks", testHomeworksExists)
	t.Run("LoginSessions", testLoginSessionsExists)
	t.Run("Messages", testMessagesExists)
	t.Run("NotificationLogs", testNotificationLogsExists)
	t.Run("PushoverTokens", testPushoverTokensExists)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesExists)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsExists)
	t.Run("UserReceivedMessages", testUserReceivedMessagesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsFind)
	t.Run("ChatroomUsers", testChatroomUsersFind)
	t.Run("Chatrooms", testChatroomsFind)
	t.Run("CourseUsers", testCourseUsersFind)
	t.Run("Courses", testCoursesFind)
	t.Run("DiaryEntries", testDiaryEntriesFind)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsFind)
	t.Run("Events", testEventsFind)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsFind)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsFind)
	t.Run("Homeworks", testHomeworksFind)
	t.Run("LoginSessions", testLoginSessionsFind)
	t.Run("Messages", testMessagesFind)
	t.Run("NotificationLogs", testNotificationLogsFind)
	t.Run("PushoverTokens", testPushoverTokensFind)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesFind)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsFind)
	t.Run("UserReceivedMessages", testUserReceivedMessagesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsBind)
	t.Run("ChatroomUsers", testChatroomUsersBind)
	t.Run("Chatrooms", testChatroomsBind)
	t.Run("CourseUsers", testCourseUsersBind)
	t.Run("Courses", testCoursesBind)
	t.Run("DiaryEntries", testDiaryEntriesBind)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsBind)
	t.Run("Events", testEventsBind)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsBind)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsBind)
	t.Run("Homeworks", testHomeworksBind)
	t.Run("LoginSessions", testLoginSessionsBind)
	t.Run("Messages", testMessagesBind)
	t.Run("NotificationLogs", testNotificationLogsBind)
	t.Run("PushoverTokens", testPushoverTokensBind)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesBind)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsBind)
	t.Run("UserReceivedMessages", testUserReceivedMessagesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsOne)
	t.Run("ChatroomUsers", testChatroomUsersOne)
	t.Run("Chatrooms", testChatroomsOne)
	t.Run("CourseUsers", testCourseUsersOne)
	t.Run("Courses", testCoursesOne)
	t.Run("DiaryEntries", testDiaryEntriesOne)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsOne)
	t.Run("Events", testEventsOne)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsOne)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsOne)
	t.Run("Homeworks", testHomeworksOne)
	t.Run("LoginSessions", testLoginSessionsOne)
	t.Run("Messages", testMessagesOne)
	t.Run("NotificationLogs", testNotificationLogsOne)
	t.Run("PushoverTokens", testPushoverTokensOne)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesOne)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsOne)
	t.Run("UserReceivedMessages", testUserReceivedMessagesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsAll)
	t.Run("ChatroomUsers", testChatroomUsersAll)
	t.Run("Chatrooms", testChatroomsAll)
	t.Run("CourseUsers", testCourseUsersAll)
	t.Run("Courses", testCoursesAll)
	t.Run("DiaryEntries", testDiaryEntriesAll)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsAll)
	t.Run("Events", testEventsAll)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsAll)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsAll)
	t.Run("Homeworks", testHomeworksAll)
	t.Run("LoginSessions", testLoginSessionsAll)
	t.Run("Messages", testMessagesAll)
	t.Run("NotificationLogs", testNotificationLogsAll)
	t.Run("PushoverTokens", testPushoverTokensAll)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesAll)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsAll)
	t.Run("UserReceivedMessages", testUserReceivedMessagesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsCount)
	t.Run("ChatroomUsers", testChatroomUsersCount)
	t.Run("Chatrooms", testChatroomsCount)
	t.Run("CourseUsers", testCourseUsersCount)
	t.Run("Courses", testCoursesCount)
	t.Run("DiaryEntries", testDiaryEntriesCount)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsCount)
	t.Run("Events", testEventsCount)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsCount)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsCount)
	t.Run("Homeworks", testHomeworksCount)
	t.Run("LoginSessions", testLoginSessionsCount)
	t.Run("Messages", testMessagesCount)
	t.Run("NotificationLogs", testNotificationLogsCount)
	t.Run("PushoverTokens", testPushoverTokensCount)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesCount)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsCount)
	t.Run("UserReceivedMessages", testUserReceivedMessagesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsHooks)
	t.Run("ChatroomUsers", testChatroomUsersHooks)
	t.Run("Chatrooms", testChatroomsHooks)
	t.Run("CourseUsers", testCourseUsersHooks)
	t.Run("Courses", testCoursesHooks)
	t.Run("DiaryEntries", testDiaryEntriesHooks)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsHooks)
	t.Run("Events", testEventsHooks)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsHooks)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsHooks)
	t.Run("Homeworks", testHomeworksHooks)
	t.Run("LoginSessions", testLoginSessionsHooks)
	t.Run("Messages", testMessagesHooks)
	t.Run("NotificationLogs", testNotificationLogsHooks)
	t.Run("PushoverTokens", testPushoverTokensHooks)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesHooks)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsHooks)
	t.Run("UserReceivedMessages", testUserReceivedMessagesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsInsert)
	t.Run("ChatroomEvents", testChatroomEventsInsertWhitelist)
	t.Run("ChatroomUsers", testChatroomUsersInsert)
	t.Run("ChatroomUsers", testChatroomUsersInsertWhitelist)
	t.Run("Chatrooms", testChatroomsInsert)
	t.Run("Chatrooms", testChatroomsInsertWhitelist)
	t.Run("CourseUsers", testCourseUsersInsert)
	t.Run("CourseUsers", testCourseUsersInsertWhitelist)
	t.Run("Courses", testCoursesInsert)
	t.Run("Courses", testCoursesInsertWhitelist)
	t.Run("DiaryEntries", testDiaryEntriesInsert)
	t.Run("DiaryEntries", testDiaryEntriesInsertWhitelist)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsInsert)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsInsertWhitelist)
	t.Run("Events", testEventsInsert)
	t.Run("Events", testEventsInsertWhitelist)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsInsert)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsInsertWhitelist)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsInsert)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsInsertWhitelist)
	t.Run("Homeworks", testHomeworksInsert)
	t.Run("Homeworks", testHomeworksInsertWhitelist)
	t.Run("LoginSessions", testLoginSessionsInsert)
	t.Run("LoginSessions", testLoginSessionsInsertWhitelist)
	t.Run("Messages", testMessagesInsert)
	t.Run("Messages", testMessagesInsertWhitelist)
	t.Run("NotificationLogs", testNotificationLogsInsert)
	t.Run("NotificationLogs", testNotificationLogsInsertWhitelist)
	t.Run("PushoverTokens", testPushoverTokensInsert)
	t.Run("PushoverTokens", testPushoverTokensInsertWhitelist)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesInsert)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesInsertWhitelist)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsInsert)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsInsertWhitelist)
	t.Run("UserReceivedMessages", testUserReceivedMessagesInsert)
	t.Run("UserReceivedMessages", testUserReceivedMessagesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ChatroomEventToChatroomUsingChatroom", testChatroomEventToOneChatroomUsingChatroom)
	t.Run("ChatroomEventToMessageUsingMessage", testChatroomEventToOneMessageUsingMessage)
	t.Run("ChatroomUserToChatroomUsingChatroom", testChatroomUserToOneChatroomUsingChatroom)
	t.Run("ChatroomUserToUserUsingUser", testChatroomUserToOneUserUsingUser)
	t.Run("CourseUserToCourseUsingCourse", testCourseUserToOneCourseUsingCourse)
	t.Run("CourseUserToUserUsingUser", testCourseUserToOneUserUsingUser)
	t.Run("DiaryEntryToUserUsingWhoPostedUser", testDiaryEntryToOneUserUsingWhoPostedUser)
	t.Run("DiaryEntryCommentToDiaryEntryUsingDiaryEntry", testDiaryEntryCommentToOneDiaryEntryUsingDiaryEntry)
	t.Run("DiaryEntryCommentToUserUsingUser", testDiaryEntryCommentToOneUserUsingUser)
	t.Run("HomeworkSubmissionCommentToHomeworkSubmissionUsingHomeworkSubmission", testHomeworkSubmissionCommentToOneHomeworkSubmissionUsingHomeworkSubmission)
	t.Run("HomeworkSubmissionCommentToUserUsingUser", testHomeworkSubmissionCommentToOneUserUsingUser)
	t.Run("HomeworkSubmissionToHomeworkUsingHomework", testHomeworkSubmissionToOneHomeworkUsingHomework)
	t.Run("HomeworkSubmissionToUserUsingUser", testHomeworkSubmissionToOneUserUsingUser)
	t.Run("HomeworkToCourseUsingCourse", testHomeworkToOneCourseUsingCourse)
	t.Run("LoginSessionToUserUsingUser", testLoginSessionToOneUserUsingUser)
	t.Run("MessageToChatroomUsingChatroom", testMessageToOneChatroomUsingChatroom)
	t.Run("MessageToUserUsingUser", testMessageToOneUserUsingUser)
	t.Run("SharedDiaryEntryToDiaryEntryUsingDiaryEntry", testSharedDiaryEntryToOneDiaryEntryUsingDiaryEntry)
	t.Run("SharedDiaryEntryToUserUsingUser", testSharedDiaryEntryToOneUserUsingUser)
	t.Run("UserReceivedChatroomEventToChatroomEventUsingChatroomEvent", testUserReceivedChatroomEventToOneChatroomEventUsingChatroomEvent)
	t.Run("UserReceivedChatroomEventToUserUsingUser", testUserReceivedChatroomEventToOneUserUsingUser)
	t.Run("UserReceivedMessageToMessageUsingMessage", testUserReceivedMessageToOneMessageUsingMessage)
	t.Run("UserReceivedMessageToUserUsingUser", testUserReceivedMessageToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ChatroomEventToUserReceivedChatroomEvents", testChatroomEventToManyUserReceivedChatroomEvents)
	t.Run("ChatroomToChatroomEvents", testChatroomToManyChatroomEvents)
	t.Run("ChatroomToChatroomUsers", testChatroomToManyChatroomUsers)
	t.Run("ChatroomToMessages", testChatroomToManyMessages)
	t.Run("CourseToCourseUsers", testCourseToManyCourseUsers)
	t.Run("CourseToHomeworks", testCourseToManyHomeworks)
	t.Run("DiaryEntryToDiaryEntryComments", testDiaryEntryToManyDiaryEntryComments)
	t.Run("DiaryEntryToSharedDiaryEntries", testDiaryEntryToManySharedDiaryEntries)
	t.Run("HomeworkSubmissionToHomeworkSubmissionComments", testHomeworkSubmissionToManyHomeworkSubmissionComments)
	t.Run("HomeworkToHomeworkSubmissions", testHomeworkToManyHomeworkSubmissions)
	t.Run("MessageToChatroomEvents", testMessageToManyChatroomEvents)
	t.Run("MessageToUserReceivedMessages", testMessageToManyUserReceivedMessages)
	t.Run("UserToChatroomUsers", testUserToManyChatroomUsers)
	t.Run("UserToCourseUsers", testUserToManyCourseUsers)
	t.Run("UserToWhoPostedUserDiaryEntries", testUserToManyWhoPostedUserDiaryEntries)
	t.Run("UserToDiaryEntryComments", testUserToManyDiaryEntryComments)
	t.Run("UserToHomeworkSubmissionComments", testUserToManyHomeworkSubmissionComments)
	t.Run("UserToHomeworkSubmissions", testUserToManyHomeworkSubmissions)
	t.Run("UserToLoginSessions", testUserToManyLoginSessions)
	t.Run("UserToMessages", testUserToManyMessages)
	t.Run("UserToSharedDiaryEntries", testUserToManySharedDiaryEntries)
	t.Run("UserToUserReceivedChatroomEvents", testUserToManyUserReceivedChatroomEvents)
	t.Run("UserToUserReceivedMessages", testUserToManyUserReceivedMessages)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ChatroomEventToChatroomUsingChatroomEvents", testChatroomEventToOneSetOpChatroomUsingChatroom)
	t.Run("ChatroomEventToMessageUsingChatroomEvents", testChatroomEventToOneSetOpMessageUsingMessage)
	t.Run("ChatroomUserToChatroomUsingChatroomUsers", testChatroomUserToOneSetOpChatroomUsingChatroom)
	t.Run("ChatroomUserToUserUsingChatroomUsers", testChatroomUserToOneSetOpUserUsingUser)
	t.Run("CourseUserToCourseUsingCourseUsers", testCourseUserToOneSetOpCourseUsingCourse)
	t.Run("CourseUserToUserUsingCourseUsers", testCourseUserToOneSetOpUserUsingUser)
	t.Run("DiaryEntryToUserUsingWhoPostedUserDiaryEntries", testDiaryEntryToOneSetOpUserUsingWhoPostedUser)
	t.Run("DiaryEntryCommentToDiaryEntryUsingDiaryEntryComments", testDiaryEntryCommentToOneSetOpDiaryEntryUsingDiaryEntry)
	t.Run("DiaryEntryCommentToUserUsingDiaryEntryComments", testDiaryEntryCommentToOneSetOpUserUsingUser)
	t.Run("HomeworkSubmissionCommentToHomeworkSubmissionUsingHomeworkSubmissionComments", testHomeworkSubmissionCommentToOneSetOpHomeworkSubmissionUsingHomeworkSubmission)
	t.Run("HomeworkSubmissionCommentToUserUsingHomeworkSubmissionComments", testHomeworkSubmissionCommentToOneSetOpUserUsingUser)
	t.Run("HomeworkSubmissionToHomeworkUsingHomeworkSubmissions", testHomeworkSubmissionToOneSetOpHomeworkUsingHomework)
	t.Run("HomeworkSubmissionToUserUsingHomeworkSubmissions", testHomeworkSubmissionToOneSetOpUserUsingUser)
	t.Run("HomeworkToCourseUsingHomeworks", testHomeworkToOneSetOpCourseUsingCourse)
	t.Run("LoginSessionToUserUsingLoginSessions", testLoginSessionToOneSetOpUserUsingUser)
	t.Run("MessageToChatroomUsingMessages", testMessageToOneSetOpChatroomUsingChatroom)
	t.Run("MessageToUserUsingMessages", testMessageToOneSetOpUserUsingUser)
	t.Run("SharedDiaryEntryToDiaryEntryUsingSharedDiaryEntries", testSharedDiaryEntryToOneSetOpDiaryEntryUsingDiaryEntry)
	t.Run("SharedDiaryEntryToUserUsingSharedDiaryEntries", testSharedDiaryEntryToOneSetOpUserUsingUser)
	t.Run("UserReceivedChatroomEventToChatroomEventUsingUserReceivedChatroomEvents", testUserReceivedChatroomEventToOneSetOpChatroomEventUsingChatroomEvent)
	t.Run("UserReceivedChatroomEventToUserUsingUserReceivedChatroomEvents", testUserReceivedChatroomEventToOneSetOpUserUsingUser)
	t.Run("UserReceivedMessageToMessageUsingUserReceivedMessages", testUserReceivedMessageToOneSetOpMessageUsingMessage)
	t.Run("UserReceivedMessageToUserUsingUserReceivedMessages", testUserReceivedMessageToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("ChatroomEventToMessageUsingChatroomEvents", testChatroomEventToOneRemoveOpMessageUsingMessage)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ChatroomEventToUserReceivedChatroomEvents", testChatroomEventToManyAddOpUserReceivedChatroomEvents)
	t.Run("ChatroomToChatroomEvents", testChatroomToManyAddOpChatroomEvents)
	t.Run("ChatroomToChatroomUsers", testChatroomToManyAddOpChatroomUsers)
	t.Run("ChatroomToMessages", testChatroomToManyAddOpMessages)
	t.Run("CourseToCourseUsers", testCourseToManyAddOpCourseUsers)
	t.Run("CourseToHomeworks", testCourseToManyAddOpHomeworks)
	t.Run("DiaryEntryToDiaryEntryComments", testDiaryEntryToManyAddOpDiaryEntryComments)
	t.Run("DiaryEntryToSharedDiaryEntries", testDiaryEntryToManyAddOpSharedDiaryEntries)
	t.Run("HomeworkSubmissionToHomeworkSubmissionComments", testHomeworkSubmissionToManyAddOpHomeworkSubmissionComments)
	t.Run("HomeworkToHomeworkSubmissions", testHomeworkToManyAddOpHomeworkSubmissions)
	t.Run("MessageToChatroomEvents", testMessageToManyAddOpChatroomEvents)
	t.Run("MessageToUserReceivedMessages", testMessageToManyAddOpUserReceivedMessages)
	t.Run("UserToChatroomUsers", testUserToManyAddOpChatroomUsers)
	t.Run("UserToCourseUsers", testUserToManyAddOpCourseUsers)
	t.Run("UserToWhoPostedUserDiaryEntries", testUserToManyAddOpWhoPostedUserDiaryEntries)
	t.Run("UserToDiaryEntryComments", testUserToManyAddOpDiaryEntryComments)
	t.Run("UserToHomeworkSubmissionComments", testUserToManyAddOpHomeworkSubmissionComments)
	t.Run("UserToHomeworkSubmissions", testUserToManyAddOpHomeworkSubmissions)
	t.Run("UserToLoginSessions", testUserToManyAddOpLoginSessions)
	t.Run("UserToMessages", testUserToManyAddOpMessages)
	t.Run("UserToSharedDiaryEntries", testUserToManyAddOpSharedDiaryEntries)
	t.Run("UserToUserReceivedChatroomEvents", testUserToManyAddOpUserReceivedChatroomEvents)
	t.Run("UserToUserReceivedMessages", testUserToManyAddOpUserReceivedMessages)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("MessageToChatroomEvents", testMessageToManySetOpChatroomEvents)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("MessageToChatroomEvents", testMessageToManyRemoveOpChatroomEvents)
}

func TestReload(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsReload)
	t.Run("ChatroomUsers", testChatroomUsersReload)
	t.Run("Chatrooms", testChatroomsReload)
	t.Run("CourseUsers", testCourseUsersReload)
	t.Run("Courses", testCoursesReload)
	t.Run("DiaryEntries", testDiaryEntriesReload)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsReload)
	t.Run("Events", testEventsReload)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsReload)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsReload)
	t.Run("Homeworks", testHomeworksReload)
	t.Run("LoginSessions", testLoginSessionsReload)
	t.Run("Messages", testMessagesReload)
	t.Run("NotificationLogs", testNotificationLogsReload)
	t.Run("PushoverTokens", testPushoverTokensReload)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesReload)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsReload)
	t.Run("UserReceivedMessages", testUserReceivedMessagesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsReloadAll)
	t.Run("ChatroomUsers", testChatroomUsersReloadAll)
	t.Run("Chatrooms", testChatroomsReloadAll)
	t.Run("CourseUsers", testCourseUsersReloadAll)
	t.Run("Courses", testCoursesReloadAll)
	t.Run("DiaryEntries", testDiaryEntriesReloadAll)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsReloadAll)
	t.Run("Events", testEventsReloadAll)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsReloadAll)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsReloadAll)
	t.Run("Homeworks", testHomeworksReloadAll)
	t.Run("LoginSessions", testLoginSessionsReloadAll)
	t.Run("Messages", testMessagesReloadAll)
	t.Run("NotificationLogs", testNotificationLogsReloadAll)
	t.Run("PushoverTokens", testPushoverTokensReloadAll)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesReloadAll)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsReloadAll)
	t.Run("UserReceivedMessages", testUserReceivedMessagesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsSelect)
	t.Run("ChatroomUsers", testChatroomUsersSelect)
	t.Run("Chatrooms", testChatroomsSelect)
	t.Run("CourseUsers", testCourseUsersSelect)
	t.Run("Courses", testCoursesSelect)
	t.Run("DiaryEntries", testDiaryEntriesSelect)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsSelect)
	t.Run("Events", testEventsSelect)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsSelect)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsSelect)
	t.Run("Homeworks", testHomeworksSelect)
	t.Run("LoginSessions", testLoginSessionsSelect)
	t.Run("Messages", testMessagesSelect)
	t.Run("NotificationLogs", testNotificationLogsSelect)
	t.Run("PushoverTokens", testPushoverTokensSelect)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesSelect)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsSelect)
	t.Run("UserReceivedMessages", testUserReceivedMessagesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsUpdate)
	t.Run("ChatroomUsers", testChatroomUsersUpdate)
	t.Run("Chatrooms", testChatroomsUpdate)
	t.Run("CourseUsers", testCourseUsersUpdate)
	t.Run("Courses", testCoursesUpdate)
	t.Run("DiaryEntries", testDiaryEntriesUpdate)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsUpdate)
	t.Run("Events", testEventsUpdate)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsUpdate)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsUpdate)
	t.Run("Homeworks", testHomeworksUpdate)
	t.Run("LoginSessions", testLoginSessionsUpdate)
	t.Run("Messages", testMessagesUpdate)
	t.Run("NotificationLogs", testNotificationLogsUpdate)
	t.Run("PushoverTokens", testPushoverTokensUpdate)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesUpdate)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsUpdate)
	t.Run("UserReceivedMessages", testUserReceivedMessagesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("ChatroomEvents", testChatroomEventsSliceUpdateAll)
	t.Run("ChatroomUsers", testChatroomUsersSliceUpdateAll)
	t.Run("Chatrooms", testChatroomsSliceUpdateAll)
	t.Run("CourseUsers", testCourseUsersSliceUpdateAll)
	t.Run("Courses", testCoursesSliceUpdateAll)
	t.Run("DiaryEntries", testDiaryEntriesSliceUpdateAll)
	t.Run("DiaryEntryComments", testDiaryEntryCommentsSliceUpdateAll)
	t.Run("Events", testEventsSliceUpdateAll)
	t.Run("HomeworkSubmissionComments", testHomeworkSubmissionCommentsSliceUpdateAll)
	t.Run("HomeworkSubmissions", testHomeworkSubmissionsSliceUpdateAll)
	t.Run("Homeworks", testHomeworksSliceUpdateAll)
	t.Run("LoginSessions", testLoginSessionsSliceUpdateAll)
	t.Run("Messages", testMessagesSliceUpdateAll)
	t.Run("NotificationLogs", testNotificationLogsSliceUpdateAll)
	t.Run("PushoverTokens", testPushoverTokensSliceUpdateAll)
	t.Run("SharedDiaryEntries", testSharedDiaryEntriesSliceUpdateAll)
	t.Run("UserReceivedChatroomEvents", testUserReceivedChatroomEventsSliceUpdateAll)
	t.Run("UserReceivedMessages", testUserReceivedMessagesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
