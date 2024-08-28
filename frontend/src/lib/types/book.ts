export type Book = {
	id: number;
	isbn: string | null;
	title: string;
	description: string;
	cover_link: string;
	published_year: string;
	published_month: string;
	published_day: string;
	author: string;
	publisher: string;
	page_count: number;
};
