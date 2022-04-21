import React from "react"
import { UserSubmissionsList } from "../../../../src/components/homework/user_submissions_list"
import { PageContainer } from "../../../../src/components/page_container"
import { useParam } from "../../../../src/utility/hokers"

export default function UserHomeworkSubmission() {
	const courseId = useParam("courseId")
	const userId = useParam("userId")

	return (
		<PageContainer>
			<UserSubmissionsList courseId={courseId} userId={userId} />
		</PageContainer>
	)
}