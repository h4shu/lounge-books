<script>
	import Header from './Header.svelte';
	import '../app.css';
	import Fab from './FAB.svelte';
</script>

<div class="app">
	<Header />

	<main>
		<slot />
	</main>

	<footer>
		<p>2024 - CA Tech Lounge</p>
	</footer>

	<Fab />
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}

	main {
		flex: 1;
		display: flex;
		flex-direction: column;
		width: 100%;
		max-width: 64rem;
		margin: 0 auto;
		box-sizing: border-box;
	}

	footer {
		background-color: var(--color-primary);
		color: white;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding: 2rem;
		margin-top: 3rem;
	}

	@media (min-width: 480px) {
		footer {
			padding: 1rem 0;
		}
	}
</style>
