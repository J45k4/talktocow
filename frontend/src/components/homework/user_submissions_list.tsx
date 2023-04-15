import React from "react"
import { useCourseUserSubmissions } from "../../data_hooks"
import { PageContainer } from "../page-container"

export const UserSubmissionsList = (props: {
	courseId: string
	userId: string
}) => {
	const submissions = useCourseUserSubmissions(props)

	return (
		<div>
			{submissions.map(submission => {
				return (
					<div key={submission.id}>
						{submission.submission}
					</div>
				)
			})}
		</div>
	)
}