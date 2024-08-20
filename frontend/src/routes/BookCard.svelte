<script lang="ts">
	import type { Book } from '$lib/types/book';

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
			{book.published_at}
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

	td {
		border-bottom: 1px solid #ccc;
		padding: 8px;
	}
</style>
