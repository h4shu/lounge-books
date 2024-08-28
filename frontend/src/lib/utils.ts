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

export const parseDate = (year: string | null, month: string | null, day: string | null) => {
	if (!year) {
		return null;
	}
	if (!day) {
		return `${year}/${month}`;
	}
	return `${year}/${month}/${day}`;
};
