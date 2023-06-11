import path from "node:path";

import { defineConfig } from "vite";
import Solid from "vite-plugin-solid";
import Dts from "vite-plugin-dts";
import Unocss from "unocss/vite";
import PurgeCss from "@mojojoejo/vite-plugin-purgecss";

export default defineConfig({
  plugins: [
    Solid(),
    Dts(),
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
          /\.css$/.test(assetInfo.name || "") ? "retalk.css" : "[name].[ext]",
      },
    },
  },
});
