import Link from "next/link";
import React from "react";
import { CourseStudentsList } from "../../../src/components/course/course_students_list";
import { HomeworkList } from "../../../src/components/homework/homework_list";
import { PageContainer } from "../../../src/components/page_container";
import { useCourseMyMeta, useParam } from "../../../src/utility/hokers";

export default function CoursePage() {
	const courseId = useParam("courseId");

	const meta = useCourseMyMeta()

	return (
		<PageContainer>
			<Link href={`/courses`}>
				Go back
			</Link>
			<h1>Course</h1>
			{meta.role === 2 && (
			<Link href={`/course/${courseId}/homework/new`}>
				<button>
					Create homework
				</button>
			</Link>)}
			{courseId &&
			<HomeworkList courseId={courseId} />}
			<CourseStudentsList courseId={courseId} />
		</PageContainer>
	)
}