export interface RegisterBookRequest {
	isbn: string;
	title: string;
	description: string;
	cover_link: string;
	published_year: number | null;
	published_month: number | null;
	published_day: number | null;
	author: string;
	publisher: string;
	page_count: number;
}
