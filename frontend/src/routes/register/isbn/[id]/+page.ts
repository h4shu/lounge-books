import type { OpenBDResponse } from '$lib/types/openbd.js';

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	const res = await fetch(`https://api.openbd.jp/v1/get?isbn=${params.id}`);
	const data: OpenBDResponse = await res.json();

	if (data[0]) {
		return {
			status: 'success',
			book_info: {
				isbn: data[0].onix.ProductIdentifier.IDValue,
				title: data[0].onix.DescriptiveDetail.TitleDetail.TitleElement.TitleText.content,
				description: data[0].onix.CollateralDetail.TextContent[0].Text,
				cover_link: `https://cover.openbd.jp/${data[0].onix.ProductIdentifier.IDValue}.jpg`,
				published_at: data[0].onix.PublishingDetail.PublishingDate[0].Date,
				author: data[0].onix.DescriptiveDetail.Contributor.map(
					(contributor) => contributor.PersonName.content
				).join(', '),
				publisher: data[0].onix.PublishingDetail.Publisher.PublisherName,
				page_count: data[0].onix.DescriptiveDetail.Extent[0].ExtentValue
			}
		};
	} else {
		return {
			status: 'error'
		};
	}
}
