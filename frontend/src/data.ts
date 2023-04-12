import { ReactiveVar } from "./reactive-var";
import { User } from "./types";


export const chatroomMembers = new ReactiveVar<User[]>({
	initialState: [],
	fetch: async () => {
		return []
	},
	resolve: (a) => {

	}
})