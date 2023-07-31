import presetRemToPx from "@unocss/preset-rem-to-px";
import type { Theme } from "@unocss/preset-uno";
import type { SourceCodeTransformer } from "unocss";
import {
	defineConfig,
	presetUno,
	transformerCompileClass,
	transformerVariantGroup,
} from "unocss";

const transformers: SourceCodeTransformer[] = [];
if (process.env.NODE_ENV === "production") {
	transformers.push(transformerCompileClass());
}
transformers.push(transformerVariantGroup());

// eslint-disable-next-line @typescript-eslint/no-unnecessary-type-arguments
export default defineConfig<Theme>({
	presets: [presetUno(), presetRemToPx()],
	transformers,
	shortcuts: {
		inputlike:
			"border-2 border-solid border-normal hover:(border-primary shadow-active) focus:(border-primary shadow-active) transition duration-animation rounded-4 outline-none resize-y",
	},
	theme: {
		colors: {
			normal: "#86868B",
			primary: "#006BB8",
			second: "#11293A",
		},
		boxShadow: {
			active: "0px 1px 8px rgba(0, 107, 184, 0.8);",
			avatar: "0px 2px 10px 0px rgba(179, 190, 198, 0.40);",
			comment: "0px 2px 9px -2px rgba(179, 190, 198, 0.80);",
		},
		duration: {
			animation: "400",
		},
	},
});
