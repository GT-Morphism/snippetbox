<script lang="ts">
	import type { PageProps } from "./$types";
	import { enhance } from "$app/forms";

	let { form }: PageProps = $props();
</script>

<a
	class="underline decoration-blue-300 decoration-dashed decoration-2 hover:decoration-solid focus-visible:decoration-solid"
	href="/snippets">Go to snippets</a
>
<form class="mb-12 grid" use:enhance method="POST">
	<label for="title">Title of snippet</label>
	<input
		class="border border-blue-300"
		id="title"
		name="title"
		type="text"
		placeholder="The missing story"
		required
	/>
	{#if form?.clientErrors?.title}
		<span class="text-red-300">{form.clientErrors.title}</span>
	{/if}

	<label for="content">Content of snippet</label>
	<textarea
		class="border border-blue-300"
		id="content"
		name="content"
		placeholder="Once upon a time…"
		required
	></textarea>
	{#if form?.clientErrors?.content}
		<span class="text-red-300">{form.clientErrors.content}</span>
	{/if}

	<fieldset>
		<legend>Delete snippet in</legend>
		<div>
			<input id="365-days" name="expires_at" type="radio" value="365" checked />
			<label for="365-days">365 days (one year)</label>
		</div>
		<div>
			<input id="7-days" name="expires_at" type="radio" value="7" />
			<label for="7-days">7 days (one week)</label>
		</div>

		<div>
			<input id="1-day" name="expires_at" type="radio" value="1" />
			<label for="1-day">One day</label>
		</div>
	</fieldset>
	{#if form?.clientErrors?.expires_at}
		<span class="text-red-300">{form.clientErrors.expires_at}</span>
	{/if}

	<button class="rounded-full bg-yellow-500/25 px-1 py-2">Let's go</button>
</form>
