<script lang="ts">
	export let title: string;
	export let body: string | null;
</script>

{#if body}
	<div class="datagrid-item">
		<div class="datagrid-title">{title}</div>
		<div class="datagrid-content">{body}</div>
	</div>
{/if}
