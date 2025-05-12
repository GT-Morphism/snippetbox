import { type InferOutput } from "valibot";
import { FormValidator } from "../validation.svelte";
import { SnippetSchema } from "../schemas/snippet";
import type { IndexableErrors } from "../validation.svelte";

type SnippetData = InferOutput<typeof SnippetSchema>;

export type IndexableSnippetErrors = IndexableErrors<SnippetData>;

export const snippetForm = () =>
	new FormValidator(SnippetSchema, {
		title: "",
		content: "",
		expires_at: "365",
	});
