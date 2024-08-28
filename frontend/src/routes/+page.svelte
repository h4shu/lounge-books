<script lang="ts">
	import type { ActionData, PageData } from './$types';
	import BookList from './BookList.svelte';
	import Fab from './FAB.svelte';
	import Search from './Search.svelte';

	export let data: PageData;
	export let form: ActionData;

	$: {
		if (form && form.status === 'success') {
			data = form;
		}
	}
</script>

<svelte:head>
	<title>Lounge Books</title>
	<meta name="description" content="CA Tech Loungeの蔵書管理アプリ" />
</svelte:head>

<section>
	{#if data.status === 'success' && data.books}
		<BookList books={data.books} />
	{/if}
	<Fab />
</section>

<style>
</style>
