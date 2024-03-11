<script lang="ts">
	import AnnouncementCard from './AnnouncementCard.svelte';
	import type { Announcement } from '$lib/types';

	let announcements: Announcement[] = [
		{
			author: 'Lorenzo Good',
			title: 'Announcements Are Coming Soon',
			date: '2024-03-11',
			content: 'The announcements feature will be arriving soon.'
		}
	];
</script>

{#each announcements as announcement}
	<AnnouncementCard {announcement} />
{/each}
