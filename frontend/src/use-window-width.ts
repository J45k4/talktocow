import { useEffect, useState } from "react";

// export const useWindowWidth = () => {
// 	const [screenWidth, setScreenWidth] = useState(window.innerWidth);

// 	useEffect(() => {
// 		const handleResize = () => {
// 			setScreenWidth(window.innerWidth);
// 		};

// 		window.addEventListener("resize", handleResize);

// 		return () => {
// 			window.removeEventListener("resize", handleResize);
// 		};
// 	}, []);

// 	return screenWidth;
// }

export const useWindowWiderThan = (threshold: number) => {
	const [cond, setCond] = useState(false);

	useEffect(() => {
		const handleResize = () => {
			setCond(window.innerWidth > threshold);
		};

		window.addEventListener("resize", handleResize);

		return () => {
			window.removeEventListener("resize", handleResize);
		};
	}, [threshold]);

	return cond;
}