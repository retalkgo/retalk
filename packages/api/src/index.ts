import type { ApiConfig } from "./swagger";
import { Api as BaseApi } from "./swagger";

export class Api {
	baseApi: BaseApi<unknown>;

	constructor(apiConfig: ApiConfig = {}) {
		this.baseApi = new BaseApi({
			...apiConfig,
			baseApiParams: {
				...apiConfig.baseApiParams,
				format: "json",
			},
		});
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

	// TODO
	deleteComment() {}
}
export * from "./types";
