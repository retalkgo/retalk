import { defineConfig } from "vite"
import path from "path"
import solidPlugin from "vite-plugin-solid"

export default defineConfig({
    plugins: [solidPlugin()],
    build: {
        target: 'es2015',
        outDir: path.resolve(__dirname, "dist"),
        minify: 'esbuild',
        lib: {
            entry: path.resolve(__dirname, "src/index.tsx"),
            name: 'retalk',
            fileName: ((format) => (format == 'umd' ? 'retalk.js' : `retalk.${format}.js`)),
            formats: ['umd', 'es', 'iife']
        },
        rollupOptions: {
            output: {
                assetFileNames: (assetInfo) => (/\.css$/.test(assetInfo.name || '') ? "retalk.css" : "[name].[ext]"),
            }
        }
    },
    css: {
        modules: {
            generateScopedName: "[hash:6]"
        }
    }
})