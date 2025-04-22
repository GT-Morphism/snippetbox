<script lang="ts">
	import type { PageProps } from "./$types";

	let { data }: PageProps = $props();
</script>

<h1>LIST OF ALL SNIPPETS</h1>
<p>
	{data.text}
</p>
