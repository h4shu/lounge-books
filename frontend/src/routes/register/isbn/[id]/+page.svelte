<script lang="ts">
	import { parseNumber } from '$lib/utils';
	/** @type {import('./$types').PageData} */
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div>
	<h2>本を登録</h2>
	{#if data.status === 'success'}
		<form method="post">
			<fieldset>
				<legend>ISBN</legend>
				<div>
					<input name="isbn" type="text" value={data.book_info?.isbn} />
				</div>
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
			<button type="submit">登録</button>
		</form>
	{:else}
		<p>データの取得に失敗しました。</p>
	{/if}
</div>
