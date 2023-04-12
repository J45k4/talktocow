import { useRouter } from "next/router";
import React from "react";
import { NewHomeworkForm } from "../../../../src/components/homework/new_homework_form";
import { PageContainer } from "../../../../src/components/page_container";
import { useParam } from "../../../../src/hokers";

export default function NewHomework() {
	const router = useRouter();
	const courseId = useParam("courseId");
	
	return (
		<PageContainer>
			<h1>New Homework</h1>
			<NewHomeworkForm courseId={courseId} onHomeworkCreated={(homeworkId) => {
				router.push(`/course/${courseId}/homework/${homeworkId}`);
			}} />
		</PageContainer>
	);
}