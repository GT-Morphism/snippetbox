<script lang="ts">
	import { page } from "$app/state";
	import type { PageProps } from "./$types";

	const id = page.params.id;

	let { data }: PageProps = $props();
</script>

<h1>SNIPPETS {id}</h1>
<p>{data.text}</p>
