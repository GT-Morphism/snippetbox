import type { PageServerLoad } from "./$types";
import { error } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ fetch }) => {
	const response = await fetch("http://localhost:4000/snippets");

	if (!response.ok) {
		error(response.status, response.statusText);
	}

	const text = await response.text();

	return {
		text,
	};
};
