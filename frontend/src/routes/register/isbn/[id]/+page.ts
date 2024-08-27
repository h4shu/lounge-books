import type { OpenBDResponse } from '$lib/types/openbd.js';

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	const res = await fetch(`https://api.openbd.jp/v1/get?isbn=${params.id}`);
	const data: OpenBDResponse = await res.json();

	if (data[0]) {
		const book = data[0];

		return {
			status: 'success',
			book_info: {
				isbn: book.summary.isbn,
				title: book.summary.title,
				description: book.onix.CollateralDetail?.TextContent?.[0]?.Text,
				cover_link:
					book.onix.CollateralDetail?.SupportingResource?.[0]?.ResourceVersion?.[0]?.ResourceLink,
				published_at: book.summary.pubdate,
				author: book.summary.author,
				publisher: book.summary.publisher,
				page_count: book.onix.DescriptiveDetail?.Extent?.[0]?.ExtentValue
			}
		};
	} else {
		return {
			status: 'error'
		};
	}
}
