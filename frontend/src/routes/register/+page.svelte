<script>
	import Scanner from './Scanner.svelte';
</script>

<div>
	<h2>本を登録</h2>
	<h3>手動で登録</h3>
	<a href="/register/new">手動で登録する</a>
	<h3>バーコードをスキャン</h3>
	<div>
		<Scanner />
	</div>
</div>
