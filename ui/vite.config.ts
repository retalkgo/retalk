import { defineConfig } from "vite"
import path from "path"

export default defineConfig({
    build: {
        target: 'es2015',
        outDir: path.resolve(__dirname, "dist"),
        minify: 'esbuild',
        lib: {
            entry: path.resolve(__dirname, "src/main.ts"),
            name: 'retalk',
            fileName: ((format) => (format == 'umd' ? 'retalk.js' : `retalk.${format}.js`)),
            formats: ['umd', 'es', 'iife']
        },
        rollupOptions: {
            output: {
                assetFileNames: (assetInfo) => (/\.css$/.test(assetInfo.name || '') ? "retalk.css" : "[name].[ext]"),
            }
        }
    }
})