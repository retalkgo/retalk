import type { $Fetch } from "ofetch";

import type { Fetcher } from "../types";

export class CommentsFetcher implements Fetcher {
	constructor(private fetch: $Fetch) {}

	create(fetch: $Fetch): Fetcher {
		return new CommentsFetcher(fetch);
	}

	getComments() {}
}
