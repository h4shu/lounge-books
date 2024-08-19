<script lang="ts">
	import type { Book } from '$lib/types/book';
	import BookList from './BookList.svelte';

	let books: Book[] = [
		{
			id: 1,
			isbn: '978-4-7741-9230-3',
			title: 'JavaScript: The Good Parts',
			description:
				'本書は、JavaScriptのエキスパートであるDouglas Crockfordが、JavaScriptの良い部分を紹介する。',
			cover_link:
				'https://images-na.ssl-images-amazon.com/images/I/51ZU4b8ZGTL._SX258_BO1,204,203,200_.jpg',
			published_at: '2008-05-08',
			author: 'Douglas Crockford',
			publisher: 'オライリージャパン',
			page_count: 176
		}
	];
</script>

<svelte:head>
	<title>Lounge Books</title>
	<meta name="description" content="CA Tech Loungeの蔵書管理アプリ" />
</svelte:head>

<section>
	<h1>Lounge Books</h1>
	<input type="text" placeholder="検索" />
	<BookList {books} />
</section>

<style>
</style>
