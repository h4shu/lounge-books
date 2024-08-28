<script lang="ts">
	import { page } from '$app/stores';
	import { parseNumber } from '$lib/utils';
	/** @type {import('./$types').PageData} */
	import type { ActionData, PageData } from './$types';

	export let data: PageData;
	export let form: ActionData;
</script>

<div>
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
			<img src={data.book_info?.cover_link} alt="表紙" />
		</fieldset>
		<button type="submit">登録</button>
	</form>
</div>
