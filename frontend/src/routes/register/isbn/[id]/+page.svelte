<script lang="ts">
	/** @type {import('./$types').PageData} */
	import type { PageData, ActionData } from './$types';

	export let data: PageData;
	export let form: ActionData;
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
					<input name="published_at" type="date" value={data.book_info?.published_at} />
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
