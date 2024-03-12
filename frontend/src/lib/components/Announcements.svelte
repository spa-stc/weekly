<script lang="ts">
	import AnnouncementCard from './AnnouncementCard.svelte';
	import type { Announcement } from '$lib/types';

	export let announcements: Announcement[];
</script>

{#each announcements as announcement}
	<AnnouncementCard {announcement} />
{/each}
