import { fail, redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";
import * as v from "valibot";
import client from "$lib/api";

const SnippetSchema = v.object({
	title: v.pipe(
		v.string(),
		v.nonEmpty("Title is required"),
		v.maxLength(100, "Brother, too long; not more than 100"),
	),
	content: v.pipe(
		v.string(),
		v.nonEmpty("Content is required"),
		v.minLength(2, "Brother, too short; at least 2"),
	),
	expires_at: v.picklist(
		["1", "7", "365"],
		"Brother, only one of the following is allowed: '1', '7' or '365'",
	),
});

export const actions = {
	default: async ({ request }) => {
		const formData = Object.fromEntries(await request.formData());

		const safeParsedData = v.safeParse(SnippetSchema, formData);

		if (!safeParsedData.success) {
			return fail(400, {
				message: "Validation error",
				clientErrors: v.flatten<typeof SnippetSchema>(safeParsedData.issues).nested,
			});
		}

		const { response, error } = await client.POST("/snippets", {
			body: {
				...safeParsedData.output,
				expires_at: parseInt(safeParsedData.output.expires_at),
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
