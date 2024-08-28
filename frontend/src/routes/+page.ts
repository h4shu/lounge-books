import type { Book } from '$lib/types/book';

/** @type {import('./$types').PageLoad} */
export async function load() {
	const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/books`);
	const data: { books: Book[] } = await res.json();

	console.log(data);

	if (data) {
		return {
			status: 'success',
			books: data.books
		};
	} else {
		return {
			status: 'error'
		};
	}
}
