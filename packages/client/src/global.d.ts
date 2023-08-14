import Retalk from "./retalk";

export {};

declare global {
	interface Window {
		Retalk: Retalk;
	}
}
