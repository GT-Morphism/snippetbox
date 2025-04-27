import { fail, redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";
import { safeParse } from "$lib/forms/validation";
import client from "$lib/api";

export const actions = {
	default: async ({ request }) => {
		const formData = Object.fromEntries(await request.formData());

		const { message, success, errors, data } = safeParse(formData);

		if (!success) {
			return fail(400, {
				message,
				clientErrors: errors,
			});
		}

		const { response, error } = await client.POST("/snippets", {
			body: {
				...data,
				expires_at: parseInt(data.expires_at),
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
