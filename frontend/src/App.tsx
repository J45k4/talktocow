import { Route, Routes } from "react-router-dom"

import BrowserInfoPage from "../pages/browser"
import CallPage from "../pages/call"
import ChatroomPage from "../pages/chatroom/[chatroomId]"
import NewChatroomPage from "../pages/chatroom/new"
import ChatroomsPage from "../pages/chatrooms"
import CowGPTChatroomPage from "../pages/chats/[chatroomId]"
import ChatsPage from "../pages/chats"
import NewChatPage from "../pages/chats/new"
import CoursePage from "../pages/course/[courseId]"
import HomeworksPage from "../pages/course/[courseId]/homeworks"
import NewHomework from "../pages/course/[courseId]/homework/new"
import HomeworkPage from "../pages/course/[courseId]/homework/[homeworkId]"
import EditHomewWorkPage from "../pages/course/[courseId]/homework/[homeworkId]/edit"
import SubmitHomeworkPage from "../pages/course/[courseId]/homework/[homeworkId]/submit"
import UserHomeworkSubmission from "../pages/course/[courseId]/student/[userId]"
import NewCoursePage from "../pages/course/new"
import CoursesPage from "../pages/courses"
import DiaryEntryPage from "../pages/diary/entry/[diaryEntryId]"
import DiaryPage from "../pages/diary"
import ExperimentsPage from "../pages/experiments"
import Index from "../pages"
import PushoverTokensPage from "../pages/pushovertokens"

export default function App() {
	return (
		<Routes>
			<Route path="/" element={<Index />} />
			<Route path="/browser" element={<BrowserInfoPage />} />
			<Route path="/call" element={<CallPage />} />
			<Route path="/chatroom/new" element={<NewChatroomPage />} />
			<Route path="/chatroom/:chatroomId" element={<ChatroomPage />} />
			<Route path="/chatrooms" element={<ChatroomsPage />} />
			<Route path="/chats" element={<ChatsPage />} />
			<Route path="/chats/new" element={<NewChatPage />} />
			<Route path="/chats/:chatroomId" element={<CowGPTChatroomPage />} />
			<Route path="/course/new" element={<NewCoursePage />} />
			<Route path="/course/:courseId" element={<CoursePage />} />
			<Route path="/course/:courseId/homeworks" element={<HomeworksPage />} />
			<Route path="/course/:courseId/homework/new" element={<NewHomework />} />
			<Route path="/course/:courseId/homework/:homeworkId" element={<HomeworkPage />} />
			<Route path="/course/:courseId/homework/:homeworkId/edit" element={<EditHomewWorkPage />} />
			<Route path="/course/:courseId/homework/:homeworkId/submit" element={<SubmitHomeworkPage />} />
			<Route path="/course/:courseId/student/:userId" element={<UserHomeworkSubmission />} />
			<Route path="/courses" element={<CoursesPage />} />
			<Route path="/diary" element={<DiaryPage />} />
			<Route path="/diary/entry/:diaryEntryId" element={<DiaryEntryPage />} />
			<Route path="/experiments" element={<ExperimentsPage />} />
			<Route path="/pushovertokens" element={<PushoverTokensPage />} />
		</Routes>
	)
}
