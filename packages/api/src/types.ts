import type { RequiredDeep } from "type-fest";

import type {
	CommonResp,
	EntityAuthor,
	EntityCookedComment,
	HttpResponse,
} from "./swagger";

export interface ApiResult<T> {
	data: T;
	msg: string;
	success: boolean;
}
export type ApiResultSwagger<T> = HttpResponse<ApiResult<T>, CommonResp>;
export type Comment = RequiredDeep<EntityCookedComment>;
export type Author = RequiredDeep<EntityAuthor>;
