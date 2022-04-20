import React from "react"
import { async } from "rxjs";
import { Button } from "../../../../../src/components/button";
import { PageContainer } from "../../../../../src/components/page_container";
import { useCurrHomework } from "../../../../../src/data_hooks";
import { useAsync, useParam } from "../../../../../src/utility/hokers";
import { getJson, postJson } from "../../../../../src/utility/talktocow-api-helpers";

export default function SubmitHomeworkPage() {
	const { homework, courseId, homeworkId } = useCurrHomework()
	const [submission, setSubmission] = React.useState<string>("")

	return (
		<PageContainer>
			<h2>{homework.title}</h2>
			<div>
				{homework.description}
			</div>
			<textarea value={submission} onChange={e => {
				setSubmission(e.target.value)
			}} />
			<Button onClick={async () => {
				await postJson(`/api/course/${courseId}/homework/${homeworkId}/submit`, {
					submission: submission
				})
			}}>
				Submit
			</Button>
		</PageContainer>
	);
}