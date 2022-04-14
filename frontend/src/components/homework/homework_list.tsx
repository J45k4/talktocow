import React from "react"
import { useGetData } from "../../utility/hokers";

export const HomeworkList = (props: {
	courseId: string
}) => {
	const [homeworks] = useGetData(`/api/course/${props.courseId}/homeworks`, []);
	
	return (
		<div>
			{homeworks.map(p => (
				<div key={p.id}>
					{p.name}
				</div>
			))}
		</div>
	);
};