import type { $Fetch } from "ofetch";

import type { Fetcher } from "../types";

export class CommentsFetcher implements Fetcher {
	private constructor(private $fetch: $Fetch) {}

	static create($fetch: $Fetch): CommentsFetcher {
		return new CommentsFetcher($fetch);
	}

	getComments() {}
}
