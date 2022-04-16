import Link from "next/link";
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
					<Link href={`/course/${props.courseId}/homework/${p.id}`}>
						{p.title}
					</Link>
				</div>
			))}
		</div>
	);
};