import { render } from "solid-js/web";

import { Retalk as RetalkComponent } from "./components/Retalk";
import type { Options } from "./types";
import { logRetalkInfo, resolveElement } from "./utils";

export default class Retalk {
	constructor({ el, ...options }: Options) {
		const resolvedEl = resolveElement(el);
		if (!resolvedEl) {
			throw new Error(`Retalk: Element ${el as string} not found`);
		}
		render(() => <RetalkComponent />, resolvedEl);
		options.logRetalkInfo && logRetalkInfo();
	}
}
