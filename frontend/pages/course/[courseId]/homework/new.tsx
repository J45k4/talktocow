import React from "react";
import { useNavigate } from "react-router-dom";
import { NewHomeworkForm } from "../../../../src/components/homework/new_homework_form";
import { PageContainer } from "../../../../src/components/page-container";
import { useParam } from "../../../../src/hokers";

export default function NewHomework() {
	const navigate = useNavigate();
	const courseId = useParam("courseId");
	
	return (
		<PageContainer>
			<h1>New Homework</h1>
			<NewHomeworkForm courseId={courseId} onHomeworkCreated={(homeworkId) => {
				navigate(`/course/${courseId}/homework/${homeworkId}`);
			}} />
		</PageContainer>
	);
}
