<script lang="ts">
	import Announcements from '$lib/components/Announcements.svelte';
	import Days from '$lib/components/Days.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>Student Newsletter - Week Overview</title>
</svelte:head>

<div class="card">
	<div class="card-header">
		<ul class="nav nav-tabs card-header-tabs" data-bs-toggle="tabs">
			<li class="nav-item">
				<a href="#tabs-home-1" class="nav-link active" data-bs-toggle="tab">🏠</a>
			</li>
			<li class="nav-item">
				<a href="#tabs-home-2" class="nav-link" data-bs-toggle="tab">📣</a>
			</li>
		</ul>

		<div class="card-actions">
			<a href="/" class="btn btn-sm">Submit Announcement</a>
		</div>
	</div>
	<div class="card-body">
		<div class="tab-content">
			<div class="tab-pane active show" id="tabs-home-1">
				<a href="/" class="btn btn-sm btn-ghost-rss mb-2" style="font-size: large">
					Week Overview:
				</a>
				<Days days={data.week} />
			</div>
			<div class="tab-pane" id="tabs-home-2">
				<Announcements />
			</div>
		</div>
	</div>
</div>
