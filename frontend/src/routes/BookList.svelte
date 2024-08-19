<script lang="ts">
	import type { Book } from '$lib/types/book';
	import BookCard from './BookCard.svelte';

	export let books: Book[];
</script>

<div class="book_card_list">
	{#each books as book}
		<BookCard {book} />
	{/each}
</div>

<style>
</style>
