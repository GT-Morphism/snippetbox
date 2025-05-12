import * as v from "valibot";

export type IndexableErrors<TData> = Partial<Record<keyof TData, [string, ...string[]]>>;

type SafeParseSuccess<TData> = {
	success: true;
	message: string;
	errors: Record<string, never>;
	data: TData;
};

type SafeParseError<TData> = {
	success: false;
	message: string;
	errors: IndexableErrors<TData>;
	data: Record<string, never>;
};

type SafeParseResult<TData> = SafeParseSuccess<TData> | SafeParseError<TData>;

export function safeParse<TSchema extends v.BaseSchema<unknown, unknown, v.BaseIssue<unknown>>>(
	schema: TSchema,
	data: Record<string, FormDataEntryValue>,
): SafeParseResult<v.InferOutput<TSchema>> {
	type TData = v.InferOutput<TSchema>;
	const safeParsedData = v.safeParse(schema, data);

	if (!safeParsedData.success) {
		return {
			success: false,
			message: "Validation error",
			errors: v.flatten<TSchema>(safeParsedData.issues).nested,
			data: {},
		} as SafeParseError<TData>;
	}

	return {
		success: true,
		message: "Validation success",
		errors: {},
		data: safeParsedData.output,
	} as SafeParseSuccess<TData>;
}

export class FormValidator<
	TSchema extends v.ObjectSchema<v.ObjectEntries, v.ErrorMessage<v.ObjectIssue> | undefined>,
> {
	formData = $state<v.InferOutput<TSchema>>({});
	errors = $state<IndexableErrors<v.InferOutput<TSchema>>>({});
	schema: TSchema;

	constructor(schema: TSchema, initialData: Partial<v.InferOutput<TSchema>> = {}) {
		this.schema = schema;
		this.formData = { ...initialData };
	}

	updateErrors = (errors: IndexableErrors<v.InferOutput<TSchema>>) => {
		const entries = Object.entries(errors) as Array<
			[keyof v.InferOutput<TSchema>, [string, ...string[]]]
		>;
		for (const [key, value] of entries) {
			this.errors[key] = value;
		}
	};

	validateField = (field: keyof v.InferOutput<TSchema>) => {
		const fieldSchema = v.pick(this.schema, [field]);
		const parsedField = v.safeParse(fieldSchema, { [field]: this.formData[field] });

		if (!parsedField.success) {
			const flattenErrors = v.flatten<typeof fieldSchema>(parsedField.issues)
				.nested as IndexableErrors<v.InferOutput<TSchema>>;
			this.errors[field] = flattenErrors?.[field];
		} else {
			this.errors[field] = undefined;
		}
	};
}
