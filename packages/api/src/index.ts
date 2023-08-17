import type { ApiConfig } from "./swagger";
import { Api as BaseApi } from "./swagger";

export class Api<SecurityDataType> extends BaseApi<SecurityDataType> {
	constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
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
