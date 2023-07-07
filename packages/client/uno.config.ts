import presetRemToPx from "@unocss/preset-rem-to-px";
import type { Theme } from "@unocss/preset-uno";
import {
  defineConfig,
  presetUno,
  transformerCompileClass,
  transformerVariantGroup,
} from "unocss";

const transformers = [transformerVariantGroup()];
if (process.env.NODE_ENV === "production") {
  transformers.push(transformerCompileClass());
}

export default defineConfig<Theme>({
  presets: [presetUno(), presetRemToPx()],
  transformers,
  shortcuts: {
    inputlike:
      "border-2.5 border-solid border-normal hover:(border-primary shadow-active) focus:(border-primary shadow-active) transition duration-animation rounded-4 outline-none",
  },
  theme: {
    colors: {
      normal: "#86868B",
      primary: "#006BB8",
    },
    boxShadow: {
      active: "0px 4px 16px rgba(0, 107, 184, 0.8);",
    },
    duration: {
      animation: "400",
    },
  },
});
