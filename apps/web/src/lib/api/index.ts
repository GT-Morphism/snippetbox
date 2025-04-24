import createClient from "openapi-fetch";
import type { paths } from "./schema.d.ts";
import { PUBLIC_API_BASE_URL } from "$env/static/public";

const client = createClient<paths>({ baseUrl: PUBLIC_API_BASE_URL });
export default client;
