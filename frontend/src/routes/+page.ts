import type { Book } from '$lib/types/book';

/** @type {import('./$types').PageLoad} */
export async function load() {
	const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/books`);
	const data: Book[] = await res.json();

	if (data) {
		return {
			status: 'success',
			books: data
		};
	} else {
		return {
			status: 'error'
		};
	}
}
