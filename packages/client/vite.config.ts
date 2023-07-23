import path from "node:path";

import PurgeCss from "@mojojoejo/vite-plugin-purgecss";
import Unocss from "unocss/vite";
import { defineConfig } from "vite";
import Dts from "vite-plugin-dts";
import Solid from "vite-plugin-solid";

export default defineConfig({
	plugins: [
		Solid(),
		Dts({
			tsconfigPath: "../../tsconfig.json",
			entryRoot: "src",
			exclude: [
				"node_module/**",
				"**/uno.config.ts",
				"vite.config.ts",
				"scripts/**",
			],
		}),
		Unocss(),
		PurgeCss({
			variables: true,
		}),
	],
	build: {
		target: "es2015",
		outDir: path.resolve(__dirname, "dist"),
		minify: "esbuild",
		lib: {
			entry: path.resolve(__dirname, "src/index.ts"),
			name: "Retalk",
			fileName: (format) =>
				format === "umd"
					? "retalk.umd.js"
					: `retalk.${format === "cjs" ? "c" : "m"}js`,
			formats: ["umd", "es", "cjs"],
		},
		rollupOptions: {
			output: {
				assetFileNames: (assetInfo) =>
					(assetInfo.name ?? "").endsWith(".css")
						? "retalk.css"
						: "[name].[ext]",
			},
		},
	},
});
