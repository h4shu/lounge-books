<script lang="ts">
	import Icon from '@iconify/svelte';
</script>

<form method="post">
	<div class="searchbar">
		<div class="searchbar-wrapper">
			<div class="searchbar-center">
				<div class="searchbar-input-spacer"></div>

				<input name="query" type="text" placeholder="キーワードで検索" class="searchbar-input" />
			</div>

			<div class="searchbar-right">
				<button type="submit" class="send-button">
					<Icon icon="bi:arrow-right-circle-fill" />
				</button>
			</div>
		</div>
	</div>
</form>

<style>
	.searchbar {
		font-size: 14px;
		font-family: arial, sans-serif;
		color: #202124;
		display: flex;
		z-index: 3;
		height: 44px;
		background: white;
		border: 1px solid #dfe1e5;
		box-shadow: none;
		border-radius: 24px;
		margin: 0 auto;
		width: auto;
		max-width: 224px;
		transition: all 0.3s;
	}

	.searchbar:hover {
		box-shadow: 0 1px 6px rgb(32 33 36 / 28%);
		border-color: rgba(223, 225, 229, 0);
	}

	.searchbar-wrapper {
		flex: 1;
		display: flex;
		padding: 5px 3px 0 14px;
	}

	.searchbar-center {
		display: flex;
		flex: 1;
		flex-wrap: wrap;
	}

	.searchbar-input-spacer {
		color: transparent;
		flex: 100%;
		white-space: pre;
		height: 34px;
		font-size: 16px;
	}

	.searchbar-input {
		background-color: transparent;
		border: none;
		margin: 0;
		padding: 0;
		color: rgba(0, 0, 0, 0.87);
		word-wrap: break-word;
		outline: none;
		display: flex;
		flex: 100%;
		margin-top: -37px;
		height: 34px;
		font-size: 16px;
		max-width: 100%;
		width: 100%;
	}

	.searchbar-right {
		display: flex;
		flex: 0 0 auto;
		margin-top: -5px;
		align-items: stretch;
		flex-direction: row;
	}
	.send-button {
		border: none;
		background-color: transparent;
		font-size: 2rem;
		color: var(--color-primary);
		cursor: pointer;
		display: flex;
		align-items: center;
	}
</style>
