import tailwindcss from "@tailwindcss/vite";
import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";
import fs from "fs";

export default defineConfig(({ command }) => {
	const SHARED_CONFIG = {
		plugins: [sveltekit(), tailwindcss()],
	};

	// Only apply server settings in development
	// reference: https://vite.dev/config/#conditional-config
	if (command === "serve") {
		return {
			...SHARED_CONFIG,
			server: {
				https: {
					cert: fs.readFileSync("./tls/certs-chain.pem"),
					key: fs.readFileSync("./tls/key.pem"),
				},
			},
		};
	}

	return SHARED_CONFIG;
});
