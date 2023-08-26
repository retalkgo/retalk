import { Api } from "@retalkgo/api";
import { ErrorBoundary } from "solid-js";
import { render } from "solid-js/web";

import { version } from "../package.json";
import { Retalk as RetalkComponent } from "./components/Retalk";
import { ApiProvider } from "./contexts/api";
import { OptionsProvider } from "./contexts/options";
import type { Options } from "./types";
import { logRetalkInfo, resolveElement, resolveOptions } from "./utils";

export default class Retalk {
	#destroy: () => void;
	api: Api;
	
	

	constructor(_options: Options) {
		const resolvedOptions = resolveOptions(_options);
		const { el, ...options } = resolvedOptions;
		const resolvedEl = resolveElement(el);
		if (!resolvedEl) {
			throw new Error(`Retalk: Element ${el as string} not found`);
		}
		this.api = new Api({ baseUrl: options.server });
		this.#destroy = render(
			() => (
				<ErrorBoundary
					fallback={(err) => <div>Critical error: {err.toString()}</div>}
				>
					<OptionsProvider options={resolvedOptions}>
						<ApiProvider api={this.api}>
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
static get version (){
	return version
}
}
