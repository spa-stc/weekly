<script lang="ts">
	import type { Announcement } from '$lib/types';
	import dayjs from 'dayjs';

	export let announcement: Announcement;

	let daystring = dayjs(announcement.date).format('YYYY/MM/DD');
</script>

<div class="card">
	<div class="card-header">
		<h3 class="card-title">{announcement.title ? announcement.title : 'Annoucement'}</h3>
	</div>
	<div class="card-body">
		<p>{announcement.content ? announcement.content : 'None'}</p>

		<div class="d-flex align-items-center pt-4 mt-auto">
			<div>
				<a href="/" class="text-body">From: {announcement.author ? announcement.author : 'None'}</a>
				<div class="text-secondary">{daystring}</div>
			</div>
		</div>
	</div>
</div>
