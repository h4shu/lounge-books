<script lang="ts">
	import type { Book } from '$lib/types/book';
	import BookCard from './BookCard.svelte';

	export let books: Book[];

	console.log(books);

	let display_type: 'card' | 'table' = 'table';
</script>

<div class="display_type">
	<select
		on:change={(event) => {
			const target = event.target as HTMLSelectElement;
			display_type = target.value as 'card' | 'table';
		}}
	>
		<option value="table">リスト表示</option>
		<option value="card">カード表示</option>
	</select>
</div>
{#if display_type === 'card'}
	<div class="grid_list">
		{#each books as book}
			<BookCard {book} />
		{/each}
	</div>
{:else}
	<table class="table_list">
		<tbody>
			<tr>
				<th>タイトル</th>
				<th>著者</th>
				<th>出版社</th>
				<th>出版日</th>
				<th>ページ数</th>
			</tr>
			{#each books as book}
				<BookCard {book} display_type="table" />
			{/each}
		</tbody>
	</table>
{/if}

<style>
	.grid_list {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
		gap: 1rem;
	}
	.table_list {
		display: table;
		width: 100%;

		th {
			background-color: #f0f0f0;
			font-weight: bold;
		}
	}
</style>
