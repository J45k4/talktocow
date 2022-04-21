import React from "react"
import { useCourseStudents } from "../../data_hooks"
import { AddStudentForm } from "./add_student_form"

export const CourseStudentsList = (props: {
	courseId: string
}) => {
	const students = useCourseStudents(props)

	return (
		<div>
			<h2>Students</h2>
			<div>
				{students.map(p => (
					<div key={p.id}>
						{p.name}
					</div>
				))}
			</div>
			
			<AddStudentForm courseId={props.courseId} />
		</div>
	)
}