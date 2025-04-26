<script lang="ts">
	import type { PageProps } from "./$types";

	let { data }: PageProps = $props();
</script>

<h1>LIST OF ALL SNIPPETS</h1>

<ul class="grid gap-y-2">
	{#each data.snippets as snippet (snippet.id)}
		<li>
			<article class="rounded-md border border-blue-300 px-2 py-4">
				<header>
					<h2 class="text-xl/tight">{snippet.title}</h2>
				</header>

				<footer class="grid">
					<a
						class="underline decoration-blue-300 decoration-dashed decoration-2 hover:decoration-solid focus-visible:decoration-solid"
						href={`/snippets/${snippet.id}`}>To the snippet</a
					>
					<aside class="text-sm">
						<span>{snippet.id}</span>
						<time>{snippet.created_at}</time>
						<time>{snippet.expires_at}</time>
					</aside>
				</footer>
			</article>
		</li>
	{/each}
</ul>
