import type { $Fetch } from "ofetch";

export abstract class Fetcher {
	static create: (fetch: $Fetch) => Fetcher;
}
