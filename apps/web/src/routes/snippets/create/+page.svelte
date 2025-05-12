<script lang="ts">
	import type { PageProps } from "./$types";
	import { enhance } from "$app/forms";
	import { snippetForm, type IndexableSnippetErrors } from "$lib/forms/validation/snippet";

	let { form }: PageProps = $props();
	const { formData: snippetData, errors, updateErrors, validateField } = snippetForm();
</script>

<a
	class="underline decoration-blue-300 decoration-dashed decoration-2 hover:decoration-solid focus-visible:decoration-solid"
	href="/snippets">Go to snippets</a
>

<div class="grid">
	<span>title: {snippetData.title}</span>
	<span>content: {snippetData.content}</span>
	<span>expires_at: {snippetData.expires_at}</span>
</div>

<form
	class="mb-12 grid"
	use:enhance={() => {
		return async ({ result, update }) => {
			if (result.type === "failure" && result.status === 400) {
				updateErrors(result.data?.clientErrors as IndexableSnippetErrors);
			}
			await update();
		};
	}}
	method="POST"
>
	<label for="title">Title of snippet</label>
	<input
		class="border border-blue-300"
		id="title"
		name="title"
		type="text"
		bind:value={snippetData.title}
		placeholder="The missing story"
		required
		maxlength="100"
		aria-invalid={errors?.title ? "true" : undefined}
		onblur={() => validateField("title")}
	/>
	{#if errors?.title}
		<span class="text-red-300">{errors.title}</span>
	{/if}

	<label for="content">Content of snippet</label>
	<textarea
		class="border border-blue-300"
		id="content"
		name="content"
		bind:value={snippetData.content}
		placeholder="Once upon a time…"
		required
		minlength="2"
		aria-invalid={errors?.content ? "true" : undefined}
		onblur={() => validateField("content")}
	></textarea>
	{#if errors?.content}
		<span class="text-red-300">{errors.content}</span>
	{/if}

	<fieldset>
		<legend>Delete snippet in</legend>
		<div>
			<input
				id="365-days"
				name="expires_at"
				type="radio"
				value="365"
				bind:group={snippetData.expires_at}
				onchange={() => validateField("expires_at")}
			/>
			<label for="365-days">365 days (one year)</label>
		</div>
		<div>
			<input
				id="7-days"
				name="expires_at"
				type="radio"
				value="7"
				bind:group={snippetData.expires_at}
				onchange={() => validateField("expires_at")}
			/>
			<label for="7-days">7 days (one week)</label>
		</div>

		<div>
			<input
				id="1-day"
				name="expires_at"
				type="radio"
				value="1"
				bind:group={snippetData.expires_at}
				onchange={() => validateField("expires_at")}
			/>
			<label for="1-day">One day</label>
		</div>
	</fieldset>
	{#if errors?.expires_at}
		<span class="text-red-300">{errors.expires_at}</span>
	{/if}

	<button class="rounded-full bg-yellow-500/25 px-1 py-2">Let's go</button>
</form>

<div class="grid text-red-300">
	<span>{form?.serverErrors?.title}</span>
	<span>{form?.serverErrors?.detail}</span>
</div>
