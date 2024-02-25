<script lang="ts">
	import SignInButton from 'clerk-sveltekit/client/SignInButton.svelte';
	import SignOutButton from 'clerk-sveltekit/client/SignOutButton.svelte';
	import SignUpButton from 'clerk-sveltekit/client/SignUpButton.svelte';
	import SignedIn from 'clerk-sveltekit/client/SignedIn.svelte';
	import SignedOut from 'clerk-sveltekit/client/SignedOut.svelte';
</script>

<SignedIn let:user>
	<a href="/" class="btn btn-sm">Submit Announcement</a>
	<SignOutButton class="btn btn-sm" />
</SignedIn>
<SignedOut>
	<SignInButton class="btn btn-sm" mode="modal" />
	<SignUpButton class="btn btn-sm" mode="modal" />
</SignedOut>
