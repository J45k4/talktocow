import Link from "next/link"
import React from "react"
import { useGetData, useParam } from "../../utility/hokers"

export const CoursesList = () => {
	// const courseId = useParam("courseId")

	const [courses] = useGetData<any[]>(`/api/courses`, [])
	
	return (
		<div>
			{courses.map(p => (
				<div key={p.id}>
					<Link href={`/course/${p.id}`}>
						{p.name}
					</Link>
				</div>
			))}
		</div>
	)
}