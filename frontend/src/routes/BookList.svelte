<script lang="ts">
	import type { Book } from '$lib/types/book';
	import BookCard from './BookCard.svelte';
	import Search from './Search.svelte';

	export let books: Book[];

	let display_type: 'card' | 'table' = 'table';
</script>

<div class="books_menu">
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
	<Search />
</div>
{#if display_type === 'card'}
	<div class="grid_list">
		{#each books as book}
			<BookCard {book} />
		{/each}
	</div>
{:else}
	<div class="table_wrapper">
		<table class="table_list">
			<thead>
				<tr>
					<th>タイトル</th>
					<th>著者</th>
					<th>出版社</th>
					<th>出版日</th>
					<th>ページ数</th>
				</tr>
			</thead>
			<tbody>
				{#each books as book}
					<BookCard {book} display_type="table" />
				{/each}
			</tbody>
		</table>
	</div>
{/if}

<style>
	.books_menu {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 1rem;

		.display_type {
			background: var(--color-tertiary);
			height: 40px;
			width: 300px;
			border-radius: 5px;
			position: relative;
			z-index: 1;
		}

		.display_type::after {
			position: absolute;
			content: '';
			width: 8px;
			height: 8px;
			right: 10px;
			top: 50%;
			transform: translateY(-50%) rotate(45deg);
			border-bottom: 2px solid black;
			border-right: 2px solid black;
			z-index: -1;
		}

		select {
			/* 初期化 */
			appearance: none;
			background: none;
			border: none;
			color: #333;
			font-size: 1rem;
			width: 100%;
			height: 100%;
			padding: 0 10px;
		}
	}

	.grid_list {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
		gap: 1rem;
	}
	.table_list {
		display: table;
		width: 100%;

		thead th {
			background-color: var(--color-primary);
			color: white;
			font-weight: bold;
			position: -webkit-sticky;
			position: sticky;
			top: 4rem;
			z-index: 10;
		}
	}
</style>
