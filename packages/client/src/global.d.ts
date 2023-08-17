import type Retalk from "./retalk";

export {};

declare global {
	interface Window {
		Retalk: typeof Retalk;
	}
}
