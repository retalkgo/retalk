import type { $Fetch } from "ofetch";

export interface Fetcher {
	create(fetch: $Fetch): Fetcher;
}
