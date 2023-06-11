import path from "node:path";
import { rmdir } from "node:fs/promises";

import { defineConfig } from "vite";
import Solid from "vite-plugin-solid";
import Dts from "vite-plugin-dts";

export default defineConfig(async () => {
  await rmdir(path.resolve(__dirname, "dist"), { recursive: true });

  return {
    plugins: [Solid(), Dts()],
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
    css: {
      modules: {
        generateScopedName: "[hash:6]",
      },
    },
  };
});
