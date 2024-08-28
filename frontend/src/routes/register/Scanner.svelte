<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { BrowserMultiFormatReader, type IScannerControls } from '@zxing/browser';
	import { goto } from '$app/navigation';

	let showScanner = false;
	let videoRef: HTMLVideoElement;
	let codeReader: BrowserMultiFormatReader;
	let controls: IScannerControls;

	onMount(() => {
		showScanner = true;
		codeReader = new BrowserMultiFormatReader();

		const debouncedScan = async () => {
			console.log('Scanning...');
			controls = await codeReader.decodeFromVideoDevice(undefined, videoRef, (result, error) => {
				if (result && result.getText().startsWith('97')) {
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

<!-- svelte-ignore a11y_media_has_caption -->
<video style="width: 100%;" bind:this={videoRef}></video>
