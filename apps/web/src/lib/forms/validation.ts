import * as v from "valibot";

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

type FormDataFromEntries = Record<string, FormDataEntryValue>;

type SafeParseSuccess = {
	success: true;
	message: string;
	errors: Record<string, never>;
	data: v.InferOutput<typeof SnippetSchema>;
};

type SafeParseError = {
	success: false;
	message: string;
	errors: v.FlatErrors<typeof SnippetSchema>["nested"];
	data: Record<string, never>;
};

type SafeParseResult = SafeParseSuccess | SafeParseError;

export function safeParse(data: FormDataFromEntries): SafeParseResult {
	const safeParsedData = v.safeParse(SnippetSchema, data);

	if (!safeParsedData.success) {
		return {
			success: false,
			message: "Validation error",
			errors: v.flatten<typeof SnippetSchema>(safeParsedData.issues).nested,
			data: {},
		} as SafeParseError;
	}

	return {
		success: true,
		message: "Validation success",
		errors: {},
		data: safeParsedData.output,
	} as SafeParseSuccess;
}
