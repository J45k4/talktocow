import { useEffect, useState } from "react";
import { ReactVar } from "./react-var";

export const useReactVar = <T>(rv: ReactVar<T>): T => {
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