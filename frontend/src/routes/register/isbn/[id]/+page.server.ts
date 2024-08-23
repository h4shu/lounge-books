import type { RegisterBookRequest } from '$lib/types/api.js';
import { parseNumber } from '$lib/utils.js';
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
			published_year: parseNumber(data.get('published_year') as string),
			published_month: parseNumber(data.get('published_month') as string),
			published_day: parseNumber(data.get('published_day') as string),
			page_count: parseInt(data.get('page_count')?.toString() || '0', 10)
		};
		console.log(postData);

		const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/books`, {
			method: 'POST',
			body: JSON.stringify(postData)
		});

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
