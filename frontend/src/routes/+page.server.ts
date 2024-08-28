import type { Book } from '$lib/types/book.js';

/** @type {import('./$types').Actions} */
export const actions = {
	default: async ({ request }) => {
		const formData = await request.formData();
		const q = formData.get('query');
		const res = await fetch(
			q === ''
				? `${import.meta.env.VITE_BACKEND_URL}/books`
				: `${import.meta.env.VITE_BACKEND_URL}/books?q=${q}`
		);
		const data: { books: Book[] } = await res.json();

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
};
