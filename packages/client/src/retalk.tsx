import { render } from "solid-js/web";

import { Retalk as RetalkComponent } from "./components/Retalk";
import { OptionsProvider } from "./options";
import type { Options } from "./types";
import { logRetalkInfo, resolveElement, resolveOptions } from "./utils";

export default class Retalk {
	#destroy: () => void;

	constructor(_options: Options) {
		const resolvedOptions = resolveOptions(_options);
		const { el, ...options } = resolvedOptions;
		const resolvedEl = resolveElement(el);
		if (!resolvedEl) {
			throw new Error(`Retalk: Element ${el as string} not found`);
		}
		this.#destroy = render(
			() => (
				<OptionsProvider options={resolvedOptions}>
					<RetalkComponent />
				</OptionsProvider>
			),
			resolvedEl,
		);
		options.logRetalkInfo && logRetalkInfo();
	}

	destroy() {
		this.#destroy();
	}
}
