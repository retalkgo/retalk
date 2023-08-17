import type { RequiredDeep } from "type-fest";

import type { EntityAuthor, EntityCookedComment } from "./swagger";

export interface ApiResult<T> {
	data: T;
	msg: string;
	success: boolean;
}
export type Comment = RequiredDeep<EntityCookedComment>;
export type Author = RequiredDeep<EntityAuthor>;
