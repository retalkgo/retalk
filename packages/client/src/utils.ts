import { version } from "../package.json";
import type { ElementOrSelector, Options } from "./types";

export const resolveElement = (el: ElementOrSelector): Element | null =>
	typeof el === "string" ? document.querySelector(el) : el;

export function logRetalkInfo() {
	// eslint-disable-next-line no-console
	console.log(
		`%c Retalk %c ${version} `,
		"background: #006BB8; padding: 5px; border-radius: 3px 0 0 3px; color: #fff",
		"background: #006BB818; padding: 5px; border-radius: 0 3px 3px 0; color: #006BB8",
	);
}

export const IS_CLIENT =
	typeof document !== "undefined" &&
	typeof window !== "undefined" &&
	typeof navigator !== "undefined";

export const resolveOptions = (options: Options): Required<Options> => ({
	...options,
	logRetalkInfo: options.logRetalkInfo ?? true,
});
