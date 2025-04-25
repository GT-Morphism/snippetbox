<script lang="ts">
	import { page } from "$app/state";
	import type { PageProps } from "./$types";

	const id = page.params.id;

	let { data }: PageProps = $props();
</script>

<a href="/snippets">Back to Snippets overview</a>

<h1>SNIPPET {id}</h1>
<aside>
	<time>Created at: {data.snippet.created_at}</time>
	<time>Expires at: {data.snippet.expires_at}</time>
</aside>
<h2>{data.snippet.title}</h2>
<p>{data.snippet.content}</p>
