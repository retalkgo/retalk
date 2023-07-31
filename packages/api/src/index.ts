import type { $Fetch } from "ofetch";
import { ofetch } from "ofetch";

import { CommentsFetcher } from "./fetchers/comments";

export class Api {
	private fetch: $Fetch;
	private comments: CommentsFetcher;

	constructor(apiUrl: string) {
		this.fetch = ofetch.create({
			baseURL: apiUrl,
		});
		this.comments = new CommentsFetcher(this.fetch);
	}
}
