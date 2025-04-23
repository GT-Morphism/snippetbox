import type { PageServerLoad } from "./$types";
import { error } from "@sveltejs/kit";
import client from "$lib/api";

export const load: PageServerLoad = async ({ fetch }) => {
	const { response, data } = await client.GET("/snippets", {
		fetch,
	});

	if (!response.ok) {
		error(response.status, response.statusText);
	}

	if (!data) {
		return {
			snippets: [],
		};
	}

	return {
		snippets: data,
	};
};
