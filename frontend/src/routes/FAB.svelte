<script lang="ts">
	import Icon from '@iconify/svelte';
</script>

<a href="/register/new" class="floating_button">
	<Icon icon="tabler:circle-plus" width="1.5rem" />
	本を登録する
</a>

<style scoped>
	.floating_button {
		position: fixed;
		bottom: 1rem;
		right: 1.5rem;
		padding: 1rem 1.5rem;
		background-color: var(--color-primary);
		color: white;
		border: 1px solid white;
		border-radius: 0.5rem;
		text-decoration: none;
		font-size: 1.2rem;
		z-index: 1000;
		display: flex;
		align-items: center;
		gap: 0.5rem;
		transition: all 0.3s;

		&:hover {
			background-color: var(--color-secondary);
		}
	}
</style>
