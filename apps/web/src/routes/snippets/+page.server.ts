import type { PageServerLoad } from "./$types";
import { error } from "@sveltejs/kit";

interface Snippet {
	id: number;
	title: string;
	created_at: string;
	expires_at: string;
}

export const load: PageServerLoad = async ({ fetch }) => {
	const response = await fetch("http://localhost:4000/snippets");

	if (!response.ok) {
		error(response.status, response.statusText);
	}

	const snippets = (await response.json()) as Snippet[];

	return {
		snippets,
	};
};
