import path from "node:path";

import { defineConfig } from "vite";
import Solid from "vite-plugin-solid";

export default defineConfig({
  plugins: [Solid()],
  build: {
    target: "es2015",
    outDir: path.resolve(__dirname, "dist"),
    minify: "esbuild",
    lib: {
      entry: path.resolve(__dirname, "src/index.tsx"),
      name: "retalk",
      fileName: (format) =>
        format === "umd" ? "retalk.js" : `retalk.${format}.js`,
      formats: ["umd", "es", "iife"],
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
});
