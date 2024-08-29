<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { BrowserMultiFormatReader, type IScannerControls } from '@zxing/browser';
	import { goto } from '$app/navigation';

	let showScanner = true;
	let videoRef: HTMLVideoElement;
	let codeReader: BrowserMultiFormatReader;
	let controls: IScannerControls;

	let isbn: string = '';

	onMount(() => {
		showScanner = true;
		codeReader = new BrowserMultiFormatReader(undefined, {
			delayBetweenScanAttempts: 1
		});

		const debouncedScan = async () => {
			controls = await codeReader.decodeFromVideoDevice(undefined, videoRef, (result, error) => {
				if (result && result.getText().startsWith('97') && result.getText() !== isbn) {
					isbn = result.getText();
					showScanner = false;
					goto(`/register/${result.getText()}`);
				}
			});
		};

		debouncedScan();
	});

	onDestroy(() => {
		showScanner = false;
		controls.stop();
	});
</script>

{#if isbn !== ''}
	<div>
		<p>ISBN: {isbn}の書誌情報を取得中......</p>
	</div>
{/if}

{#if showScanner}
	<!-- svelte-ignore a11y_media_has_caption -->
	<video style="width: 100%;" bind:this={videoRef}></video>
	<p>"978"から始まるバーコードをスキャンしてください。</p>
{/if}

<style scoped>
	video {
		border-radius: 1rem;
	}
	p {
		text-align: center;
	}
</style>
