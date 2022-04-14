import { useRouter } from "next/router";
import React from "react";

import { NewCourseForm } from "../../src/components/course/new_course_form";
import { PageContainer } from "../../src/components/page_container";

export default function NewCoursePage() {
	const router = useRouter();

	return (
		<PageContainer>
			<h1>New Course</h1>

			<NewCourseForm onCourseCreated={(courseId) => {
				router.push(`/course/${courseId}`);
			}} />
		</PageContainer>
	);
}