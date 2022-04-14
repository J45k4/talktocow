import Link from "next/link"
import React from "react"
import { CoursesList } from "../src/components/course/courses_list"
import { PageContainer } from "../src/components/page_container"

export default function CoursesPage() {
	return (
		<PageContainer>
			<h1>Courses</h1>
			<div>
				<Link href={"/course/new"}>
					<button>
						Create course
					</button>
				</Link>
			</div>
			
			<CoursesList />
		</PageContainer>
	)
}