import { fail, redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";
import { safeParse } from "$lib/forms/validation.svelte";
import { SnippetSchema } from "$lib/forms/schemas/snippet";
import client from "$lib/api";

export const actions = {
	default: async ({ request }) => {
		const formData = Object.fromEntries(await request.formData());

		const { message, success, errors, data } = safeParse(SnippetSchema, formData);

		if (!success) {
			return fail(400, {
				message,
				clientErrors: errors,
			});
		}

		const { response, error } = await client.POST("/snippets", {
			body: {
				...data,
				expires_at: parseInt(data.expires_at) as 1 | 7 | 365,
			},
		});

		if (!response.ok) {
			return fail(response.status, {
				message: response.statusText,
				serverErrors: error,
			});
		}

		const locationHeader = response.headers.get("location");

		if (!locationHeader) {
			redirect(303, "/snippets");
		}

		redirect(303, locationHeader);
	},
} satisfies Actions;
