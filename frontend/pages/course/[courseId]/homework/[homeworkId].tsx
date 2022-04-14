import React from "react";
import { PageContainer } from "../../../../src/components/page_container";
import { useGetData, useParam } from "../../../../src/utility/hokers";


export default function HomeworksPage() {
	const courseId = useParam("courseId");
	const homeworkId = useParam("homeworkId");
	
	const [data] = useGetData(`/api/course/${courseId}/homework/${homeworkId}`, {});

	return (
		<PageContainer>
			<h1>Homework</h1>


		</PageContainer>
	);
}