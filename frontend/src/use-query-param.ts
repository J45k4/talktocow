import { useState, useEffect } from 'react';

export const useQueryParam = (name, debounceTime = 400) => {
	const [value, setValue] = useState(null);
	const [debouncedValue, setDebouncedValue] = useState(null);

	useEffect(() => {
		const queryParams = new URLSearchParams(window.location.search);
		const paramValue = queryParams.get(name);
		setValue(paramValue);
		setDebouncedValue(paramValue);
	}, [name]);

	const updateQueryParam = debounce((newValue) => {
		const queryParams = new URLSearchParams(window.location.search);
		queryParams.set(name, newValue);
		const newUrl = `${window.location.pathname}?${queryParams.toString()}`;
		window.history.pushState({ path: newUrl }, '', newUrl);
		setValue(newValue);
	}, debounceTime);

	const handleChange = (event) => {
		const newValue = event.target.value;
		setDebouncedValue(newValue);
		updateQueryParam(newValue);
	};

	return [value, debouncedValue, handleChange];
};

const debounce = (func, delay) => {
	let timer;
	return (...args) => {
		clearTimeout(timer);
		timer = setTimeout(() => {
			func(...args);
		}, delay);
	};
};
