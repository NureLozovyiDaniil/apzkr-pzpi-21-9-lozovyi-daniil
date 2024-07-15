<script>
	import '../app.css';
	import Header from '../components/Header.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	import { ModeWatcher } from 'mode-watcher';
</script>

<Header />
<Toaster />
<ModeWatcher defaultMode={'light'} />

<slot></slot>
