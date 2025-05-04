<script lang="ts" module>
	type ToastData = {
		title: string;
		description: string;
	};

	const toaster = new Toaster<ToastData>();

	export const addToast = toaster.addToast;
</script>

<script lang="ts">
	import { Toaster } from "melt/builders";
	import { fly } from "svelte/transition";
</script>

<div
	class="fixed top-[unset] right-4 bottom-4 left-[unset] bg-[unset] text-white"
	{...toaster.root}
>
	{#each toaster.toasts as toast (toast.id)}
		<div
			class="border border-blue-300"
			in:fly={{ y: 60, opacity: 0.9 }}
			out:fly={{ y: 20 }}
			{...toast.content}
		>
			<h3 {...toast.title}>{toast.data.title}</h3>
			<div {...toast.description}>{toast.data.description}</div>
			<button {...toast.close} aria-label="dismiss alert">X</button>
		</div>
	{/each}
</div>
