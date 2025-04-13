import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch, params }) => {
	const data = await fetch(`http://localhost:4000/snippets/${params.id}`);

	const text = await data.text();

	return {
		text,
	};
};
