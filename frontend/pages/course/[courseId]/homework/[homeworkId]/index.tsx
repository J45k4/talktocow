import Link from "next/link";
import React from "react";
import { PageContainer } from "../../../../../src/components/page_container";
import { useAsync, useGetData, useParam } from "../../../../../src/hokers";
import { getJson } from "../../../../../src/api-methods";


export default function HomeworkPage() {
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
			<Link href={`/course/${courseId}`}>
				back
			</Link>
			<h1>{data.title}</h1>
			<Link href={`/course/${courseId}/homework/${homeworkId}/edit`}>
				<button style={{
					maxWidth: "70px",
				}}>
					Edit
				</button>
			</Link>
			
			<h3>Description</h3>
			<pre>
				{data.description}
			</pre>
			<Link href={`/course/${courseId}/homework/${homeworkId}/submit`}>
				Make submission
			</Link>
		</PageContainer>
	);
}