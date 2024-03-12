<script lang="ts">
	import type { Day } from '$lib/types';
	import dayjs from 'dayjs';
	import TextItem from './datagrid/TextItem.svelte';
	import OptionalTextItem from './datagrid/OptionalTextItem.svelte';

	export let day: Day;

	let datestring = dayjs(day.date).add(1, 'day').format('dddd - MM/DD/YY');
	let lunch = day.lunch?.replaceAll('\\n', ', ') || 'TBD';
</script>

<div class="card mb-1">
	<div class="card-header">
		<h3 class="card-title">{datestring} - Day {day.rotation_day ? day.rotation_day : 'N/A'}:</h3>
	</div>
	<div class="card-body">
		<div class="datagrid">
			<TextItem title="Lunch" body={lunch} />
			<TextItem title="XPeriod" body={day.x_period} />
			<TextItem title="XPeriod Location" body={day.location} />
			<TextItem title="Notes" body={day.notes} />

			<!-- Optional Text Items -->
			<OptionalTextItem title="Grade 9" body={day.grade9} />
			<OptionalTextItem title="Grade 10" body={day.grade10} />
			<OptionalTextItem title="Grade 11" body={day.grade11} />
			<OptionalTextItem title="Grade 12" body={day.grade12} />

			<OptionalTextItem title="College Counseling Info" body={day.cc_info} />
			<OptionalTextItem title="AP Info" body={day.ap_info} />
		</div>
	</div>
</div>
