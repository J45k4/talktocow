import React from "react";
import { useNavigate } from "react-router-dom";

import { NewCourseForm } from "../../src/components/course/new_course_form";
import { PageContainer } from "../../src/components/page-container";

export default function NewCoursePage() {
	const navigate = useNavigate();

	return (
		<PageContainer>
			<h1>New Course</h1>

			<NewCourseForm onCourseCreated={(courseId) => {
				navigate(`/course/${courseId}`);
			}} />
		</PageContainer>
	);
}
