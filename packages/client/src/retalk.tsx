import { ErrorBoundary } from "solid-js";
import { render } from "solid-js/web";

import { ApiProvider } from "./api";
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
				<ErrorBoundary
					fallback={(err) => <div>Critical error: {err.toString()}</div>}
				>
					<OptionsProvider options={resolvedOptions}>
						<ApiProvider>
							<RetalkComponent />
						</ApiProvider>
					</OptionsProvider>
				</ErrorBoundary>
			),
			resolvedEl,
		);
		options.logRetalkInfo && logRetalkInfo();
	}

	destroy() {
		this.#destroy();
	}
}
