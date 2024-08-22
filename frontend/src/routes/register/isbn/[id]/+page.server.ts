import type { RegisterBookRequest } from '$lib/types/api.js';
import { fail } from '@sveltejs/kit';

/** @type {import('./$types').Actions} */
export const actions = {
	default: async ({ request }) => {
		const data = await request.formData();
		const postData: RegisterBookRequest = {
			isbn: data.get('isbn')?.toString() || '',
			title: data.get('title')?.toString() || '',
			description: data.get('description')?.toString() || '',
			author: data.get('author')?.toString() || '',
			cover_link: data.get('cover_link')?.toString() || '',
			publisher: data.get('publisher')?.toString() || '',
			published_year: data.get('published_year') as number | null,
			published_month: data.get('published_month') as number | null,
			published_day: data.get('published_day') as number | null,
			page_count: parseInt(data.get('page_count')?.toString() || '0', 10)
		};

		const res = await fetch(`${process.env.VITE_BACKEND_URL}/books`, {
			method: 'POST',
			body: JSON.stringify(postData)
		});

		console.log(res);

		if (res.status === 204) {
			return {
				status: 'success'
			};
		} else {
			return fail(400, {
				status: 'error'
			});
		}
	}
};
