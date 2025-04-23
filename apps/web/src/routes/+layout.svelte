<script lang="ts">
	import "../app.css";

	let { children } = $props();
</script>

<main class="wrapper">{@render children()}</main>
