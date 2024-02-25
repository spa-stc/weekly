<script lang="ts">
	import type { Day } from '$lib/types';
	import Daycard from './Daycard.svelte';

	export let days: Day[];
</script>

<div>
	{#each days as day}
		<Daycard {day} />
	{/each}
</div>
