<script lang="ts">
	import type { Book } from '$lib/types/book';
	import BookCard from './BookCard.svelte';

	export let books: Book[];
</script>

<div class="display_type">
	<select>
		<option>リスト表示</option>
		<option>カード表示</option>
	</select>
</div>
<div class="book_card_list">
	{#each books as book}
		<BookCard {book} />
	{/each}
</div>

<style>
	.book_card_list {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
		gap: 1rem;
	}
</style>
