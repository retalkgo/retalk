import type { $Fetch } from "ofetch";
import { ofetch } from "ofetch";

import { CommentsFetcher } from "./fetchers/comments";

export class Api {
	#fetch: $Fetch;
	comments: CommentsFetcher;

	constructor(apiUrl: string) {
		this.#fetch = ofetch.create({
			baseURL: apiUrl,
		});
		this.comments = CommentsFetcher.create(this.#fetch);
	}
}
