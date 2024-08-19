<script lang="ts">
	import type { Book } from '$lib/types/book';

	export let book: Book;
</script>

<div class="book_card">
	<h2>{book.title}</h2>
	<p>{book.author}</p>
	<p>{book.publisher}</p>
	<p>{book.published_at}</p>
	<p>{book.isbn}</p>
	<p>{book.description}</p>
</div>

<style>
	.book_card {
		border: 1px solid #ccc;
		border-radius: 4px;
		padding: 8px;
		margin: 8px;

		h2 {
			font-size: 1.5rem;
			font-weight: 800;
		}
	}
</style>
