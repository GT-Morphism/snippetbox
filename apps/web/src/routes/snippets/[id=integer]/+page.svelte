<script lang="ts">
	import { page } from "$app/state";
	import { onMount } from "svelte";
	import type { PageProps } from "./$types";
	import { addToast } from "$lib/components/Toaster.svelte";

	const id = page.params.id;

	let { data }: PageProps = $props();

	onMount(() => {
		if (data.showToast) {
			addToast({
				data: {
					title: "Amazing job",
					description: "Here is your fresh snippet",
				},
			});
		}
	});
</script>

<a
	class="underline decoration-blue-300 decoration-dashed decoration-2 hover:decoration-solid focus-visible:decoration-solid"
	href="/snippets">Back to Snippets overview</a
>

<h1 class="mb-4 text-2xl/tight">{data.snippet.title} (ID = {id})</h1>

<aside class="mb-8 text-sm">
	<time>Created at: {data.snippet.created_at}</time>
	<time>Expires at: {data.snippet.expires_at}</time>
</aside>

<pre
	class="relative ps-4 before:absolute before:top-0 before:bottom-0 before:left-0 before:w-1 before:bg-blue-300 before:content-['']">{data
		.snippet.content}</pre>
