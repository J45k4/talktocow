import Link from "next/link";
import React from "react";
import { HomeworkList } from "../../../src/components/homework/homework_list";
import { useParam } from "../../../src/utility/hokers";

export default function CoursePage() {
	const courseId = useParam("courseId");

	return (
		<div>
			<h1>Course</h1>
			<Link href={`/course/${courseId}/homework/new`}>
				<button>
					Create homework
				</button>
			</Link>
			<HomeworkList courseId={courseId} />
		</div>
	)
}