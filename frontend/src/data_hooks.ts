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