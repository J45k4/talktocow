import { useAsync, useParam } from "./utility/hokers";
import { getJson } from "./utility/talktocow-api-helpers";

export const useCurrHomework = () => {
	const courseId = useParam("courseId");
	const homeworkId = useParam("homeworkId");

	const homework = useAsync(async () => {
		if (!courseId || !homeworkId) {
			return {}
		}

		const res = await getJson<any>(`/api/course/${courseId}/homework/${homeworkId}`);

		return res.payload || {}
	}, [courseId, homeworkId], {});

	return {
		homework,
		courseId,
		homeworkId
	}
}

export const useCourseUserSubmissions = (args: {
	courseId: string
	userId: string
}): {
	id: string
	submission: string
	status: string
}[] => {
	const submissions = useAsync(async () => {
		if (!args.courseId || !args.userId) {
			return []
		}

		const res = await getJson<any>(`/api/course/${args.courseId}/student/${args.userId}/submissions`);

		return res.payload || []
	}, [args.courseId, args.userId], []);

	return submissions
}

export const useCourseStudents = (args: {
	courseId: string
}): {
	id: string
	name: string
}[] => {
	const courseId = useParam("courseId");

	const students = useAsync(async () => {
		if (!courseId) {
			return []
		}

		const res = await getJson<any>(`/api/course/${courseId}/students`);

		return res.payload || []
	}, [courseId], []);

	return students
}

export const useSearchUsers = (searchword?: string): {
	id: string
	name: string
}[] => {
	const users = useAsync(async () => {
		const res = await getJson<any>(`/api/users${searchword ? "?searchword=" + searchword : ""}`);

		return res.payload || []
	}, [searchword], []);

	return users
}