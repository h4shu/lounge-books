<script lang="ts">
	import { deserialize } from '$app/forms';
	import { page } from '$app/stores';
	import { parseNumber } from '$lib/utils';
	import Icon from '@iconify/svelte';
	/** @type {import('./$types').PageData} */
	import type { ActionData, PageData } from './$types';

	export let data: PageData;
	export let form: ActionData;
</script>

<div class="register_page">
	<h2>本を登録</h2>
	{#if form?.status === 'success'}
		<p>
			登録しました。
			<a href="/register">別の本を登録する</a>
		</p>
	{:else if form?.status === 'error'}
		<p>登録に失敗しました。</p>
	{/if}
	<form method="post">
		<fieldset>
			<legend>ISBN</legend>
			<div>
				<input
					name="isbn"
					type="text"
					value={data.book_info?.isbn === 'new' ? '' : data.book_info?.isbn}
				/>
			</div>
			{#if $page.params.id === 'new'}
				<a href="/register">ISBNコードをスキャンする</a>
			{/if}
		</fieldset>
		<fieldset>
			<legend>タイトル</legend>
			<div>
				<input name="title" type="text" value={data.book_info?.title} />
			</div>
		</fieldset>
		<fieldset>
			<legend>著者</legend>
			<div>
				<input name="author" type="text" value={data.book_info?.author} />
			</div>
		</fieldset>
		<fieldset>
			<legend>説明文</legend>
			<div>
				<textarea name="genre">{data.book_info?.description}</textarea>
			</div>
		</fieldset>
		<fieldset>
			<legend>出版社</legend>
			<div>
				<input name="publisher" type="text" value={data.book_info?.publisher} />
			</div>
		</fieldset>
		<fieldset>
			<legend>出版日</legend>
			<div>
				<input
					name="published_year"
					type="number"
					value={parseNumber(data.book_info?.published_at.slice(0, 4))}
				/>
				<input
					name="published_month"
					type="number"
					value={parseNumber(data.book_info?.published_at.slice(4, 6))}
				/>
				<input
					name="published_day"
					type="number"
					value={parseNumber(data.book_info?.published_at.slice(6, 8))}
				/>
			</div>
		</fieldset>
		<fieldset>
			<legend>ページ数</legend>
			<div>
				<input name="page_count" type="number" value={data.book_info?.page_count} />
			</div>
		</fieldset>
		<fieldset>
			<legend>画像URL</legend>
			<div>
				<input name="cover_link" type="text" value={data.book_info?.cover_link} />
			</div>
			{#if data.book_info?.cover_link}
				<img src={data.book_info?.cover_link} alt="表紙" />
			{/if}
		</fieldset>
		<button type="submit">
			<Icon icon="tabler:circle-plus" width="1.5rem" />
			登録
		</button>
	</form>
</div>

<style scoped>
	.register_page {
		margin: 0 1rem;
		font-size: 1rem;
	}
	form {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	button {
		background-color: var(--color-primary);
		color: white;
		border: none;
		border-radius: 0.5rem;
		font-size: 1rem;
		padding: 0.5rem 1rem;
		cursor: pointer;
		transition: all 0.3s;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 0.5rem;

		&:hover {
			background-color: var(--color-secondary);
		}
	}
</style>
