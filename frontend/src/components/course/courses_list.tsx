import React from "react"
import { Link } from "react-router-dom"
import { useGetData, useParam } from "../../hokers"

export const CoursesList = () => {
	// const courseId = useParam("courseId")

	const [courses] = useGetData<any[]>(`/api/courses`, [])
	
	return (
		<div>
			{courses.map(p => (
				<div key={p.id}>
					<Link to={`/course/${p.id}`}>
						{p.name}
					</Link>
				</div>
			))}
		</div>
	)
}
