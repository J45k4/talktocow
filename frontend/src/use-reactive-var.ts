import { useEffect, useState } from "react";
import { ReactiveVar } from "./reactive-var";

export const useReactVar = <T>(rv: ReactiveVar<T>): T => {
	const [state, setState] = useState(rv.get());

	useEffect(() => {		
		const sub = rv.sub((newState) => {
			setState(newState);
		});

		return () => {
			sub.unsubscribe();
		};
	}, [setState])

	return state;
}