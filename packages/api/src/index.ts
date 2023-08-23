import defu from "defu";

import type { ApiConfig } from "./swagger";
import { Api as BaseApi } from "./swagger";

export class Api {
	baseApi: BaseApi<unknown>;

	constructor(apiConfig: ApiConfig = {}) {
		this.baseApi = new BaseApi(
			defu(apiConfig, {
				baseApiParams: {
					format: "json",
				},
			} satisfies ApiConfig),
		);
	}

	getRoot() {
		return this.baseApi.getRoot();
	}

	getComments(path?: string) {
		return path
			? this.baseApi.api.commentGetByPathList({ path })
			: this.baseApi.api.commentGetAllList();
	}

	createComment(data: Parameters<typeof this.baseApi.api.commentAddCreate>[0]) {
		return this.baseApi.api.commentAddCreate(data);
	}

	deleteComment(
		query: Parameters<typeof this.baseApi.api.commentDeleteDelete>[0],
	) {
		return this.baseApi.api.commentDeleteDelete(query);
	}
}

export * from "./types";
