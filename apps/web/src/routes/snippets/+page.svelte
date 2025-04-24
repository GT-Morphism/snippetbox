<script lang="ts">
	import type { PageProps } from "./$types";

	let { data }: PageProps = $props();
</script>

<h1>LIST OF ALL SNIPPETS</h1>

<div class="grid gap-y-2">
	{#each data.snippets as snippet (snippet.id)}
		<article class="rounded-md border border-blue-300 px-2 py-4">
			<header>
				<h2>{snippet.title}</h2>
			</header>

			<footer class="grid">
				<a href={`/snippets/${snippet.id}`}>To the snippet</a>
				<aside>
					<span>{snippet.id}</span>
					<time>{snippet.created_at}</time>
					<time>{snippet.expires_at}</time>
				</aside>
			</footer>
		</article>
	{/each}
</div>
