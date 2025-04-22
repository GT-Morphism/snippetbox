import type { PageServerLoad } from "./$types";
import { error } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ fetch, params }) => {
	const response = await fetch(`http://localhost:4000/snippets/${params.id}`);

	if (!response.ok) {
		error(404, `Brother, for id ${params.id} no snippet.`);
	}

	const text = await response.text();

	return {
		text,
	};
};
