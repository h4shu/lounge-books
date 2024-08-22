export const parseNumber = (str: string | undefined | null) => {
	if (!str) {
		return null;
	}
	const x = Number(str);
	if (isNaN(x)) {
		return null;
	} else {
		return x;
	}
};
