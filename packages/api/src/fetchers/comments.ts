import type { $Fetch } from "ofetch";

import type { Fetcher } from "../types";

export class CommentsFetcher implements Fetcher {
	#fetch: $Fetch;

	private constructor($fetch: $Fetch) {
		this.#fetch = $fetch;
	}

	static create($fetch: $Fetch): CommentsFetcher {
		return new CommentsFetcher($fetch);
	}

	getComments() {}
}
