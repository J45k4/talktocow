import React, { useState } from "react"
import { Link, useNavigate } from "react-router-dom";
import { Button } from "../../../../../src/components/button";
import { HomeworkDescriptionEdit } from "../../../../../src/components/homework/homework_description_edit";
import { PageContainer } from "../../../../../src/components/page-container"
import { useAsync, useParam } from "../../../../../src/hokers";
import { getJson, postJson, putJson } from "../../../../../src/api-methods";


export default function EditHomewWorkPage() {
	const [title, setTitle] = useState("")
	const [ description, setDescription ] = useState<string>("")

	const courseId = useParam("courseId");
	const homeworkId = useParam("homeworkId");

	const navigate = useNavigate()

	useAsync(async () => {
		if (!courseId || !homeworkId) {
			return {}
		}

		const res = await getJson<any>(`/api/course/${courseId}/homework/${homeworkId}`);

		setTitle(res.payload.title)
		setDescription(res.payload.description)
	}, [courseId, homeworkId], {});

	return (
		<PageContainer>
			<Link to={`/course/${courseId}/homework/${homeworkId}`}>
				back
			</Link>
			<h1>{title}</h1>
			<HomeworkDescriptionEdit description={description} onChange={setDescription} />
			<Button onClick={async () => {
				console.log("save homework")

				await putJson(`/api/course/${courseId}/homework/${homeworkId}`, {
					title: title,
					description: description	
				})

				navigate(`/course/${courseId}/homework/${homeworkId}`)
			}}>
				Save homework
			</Button>
		</PageContainer>
	)
}
