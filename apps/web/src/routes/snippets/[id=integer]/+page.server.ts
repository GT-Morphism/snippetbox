import type { PageServerLoad } from "./$types";
import { error } from "@sveltejs/kit";
import client from "$lib/api";

export const load: PageServerLoad = async ({ fetch, params }) => {
	const { response, data } = await client.GET("/snippets/{id}", {
		params: {
			path: {
				id: parseInt(params.id),
			},
		},
		fetch,
	});

	if (!response.ok) {
		error(response.status, response.statusText);
	}

	if (!data) {
		error(404, `Brother, for id ${params.id} no snippet.`);
	}

	return {
		snippet: data,
	};
};
