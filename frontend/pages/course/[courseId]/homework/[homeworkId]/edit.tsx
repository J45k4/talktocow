import React from "react"
import { Button } from "../../../../../src/components/button";
import { HomeworkDescriptionEdit } from "../../../../../src/components/homework/homework_description_edit";
import { PageContainer } from "../../../../../src/components/page_container"
import { useAsync, useParam } from "../../../../../src/utility/hokers";
import { getJson, postJson } from "../../../../../src/utility/talktocow-api-helpers";


export default function EditHomewWorkPage() {
	const courseId = useParam("courseId");
	const homeworkId = useParam("homeworkId");

	const data = useAsync(async () => {
		if (!courseId || !homeworkId) {
			return {}
		}

		const res = await getJson<any>(`/api/course/${courseId}/homework/${homeworkId}`);

		return res.payload || {}
	}, [courseId, homeworkId], {});

	return (
		<PageContainer>
			<h1>{data.title}</h1>
			<HomeworkDescriptionEdit description={data.description} />
			<Button onClick={() => {
				console.log("save homework")

				postJson(`/api/course/${courseId}/homework/${homeworkId}`, {
					title: data.title,
					description: data.description	
				})
			}}>
				Save homework
			</Button>
		</PageContainer>
	)
}