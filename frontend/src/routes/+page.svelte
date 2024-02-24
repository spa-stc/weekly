<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

{#each data.week as day}
	<h1>{day.x_period}</h1>
{/each}
