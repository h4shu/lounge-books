<header>
	<a href="/">
		<h1>Lounge Books</h1>
	</a>
</header>

<style scoped>
	header {
		position: sticky;
		top: 0;
		background-color: white;
		display: flex;
		justify-content: center;
		height: 4rem;
		z-index: 10;
		a {
			text-decoration: none;
			display: flex;
			align-items: center;
			h1 {
				color: var(--color-primary);
				margin: 0;
				transition: all 0.3s;
			}
			&:hover {
				h1 {
					color: var(--color-secondary);
				}
			}
		}
	}
</style>
