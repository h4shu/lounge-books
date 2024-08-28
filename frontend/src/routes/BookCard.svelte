<script lang="ts">
	import type { Book } from '$lib/types/book';
	import { parseDate } from '$lib/utils';

	export let book: Book;

	export let display_type: 'card' | 'table' = 'card';
</script>

{#if display_type === 'card'}
	<div class="book_card">
		<img src={book.cover_link} alt={book.title} />
	</div>
{:else}
	<tr>
		<td>
			{book.title}
		</td>
		<td>
			{book.author}
		</td>
		<td>
			{book.publisher}
		</td>
		<td>
			{parseDate(book.published_year, book.published_month, book.published_day)}
		</td>
		<td>
			{book.page_count}
		</td>
	</tr>
{/if}

<style>
	.book_card {
		border: 1px solid #ccc;
		border-radius: 4px;
		padding: 8px;
		margin: 8px;
	}

	img {
		width: 100%;
	}

	td {
		border-bottom: 1px solid #ccc;
		padding: 8px;
	}
</style>
