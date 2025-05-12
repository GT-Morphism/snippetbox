import * as v from "valibot";

export const SnippetSchema = v.object({
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
