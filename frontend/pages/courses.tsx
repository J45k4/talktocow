import React from "react"
import { Link } from "react-router-dom"
import { CoursesList } from "../src/components/course/courses_list"
import { PageContainer } from "../src/components/page-container"

export default function CoursesPage() {
	return (
		<PageContainer>
			<h1>Courses</h1>
			<div>
				<Link to={"/course/new"}>
					<button>
						Create course
					</button>
				</Link>
			</div>
			
			<CoursesList />
		</PageContainer>
	)
}
