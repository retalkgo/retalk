import type { ApiConfig } from "./swagger";
import { Api as BaseApi } from "./swagger";

export class Api<T> extends BaseApi<T> {
	constructor(apiConfig: ApiConfig<T> = {}) {
		super();
		Object.assign(
			this,
			Object.assign(apiConfig, {
				baseApiParams: {
					...apiConfig.baseApiParams,
					format: "json",
				},
			}),
		);
	}
}
export * from "./types";
