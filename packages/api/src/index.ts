import defu from "defu";

import type { ApiConfig } from "./swagger";
import { Api as BaseApi } from "./swagger";
import type { ApiResultSwagger, Comment } from "./types";

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

	getRoot(): Promise<ApiResultSwagger<null>> {
		return this.baseApi.getRoot() as any;
	}

	getComments(path?: string): Promise<ApiResultSwagger<Comment[]>> {
		return (
			path
				? this.baseApi.api.commentGetByPathList({ path })
				: this.baseApi.api.commentGetAllList()
		) as any;
	}

	createComment(
		data: Parameters<typeof this.baseApi.api.commentAddCreate>[0],
	): /* TODO */ Promise<any> {
		return this.baseApi.api.commentAddCreate(data);
	}

	deleteComment(
		query: Parameters<typeof this.baseApi.api.commentDeleteDelete>[0],
	): /* TODO */ Promise<any> {
		return this.baseApi.api.commentDeleteDelete(query);
	}
}

export * from "./types";
