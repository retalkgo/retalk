import type { MountableElement } from "solid-js/web";

export type ElementOrSelector = string | MountableElement;

export interface Options {
	el: ElementOrSelector;
	server: string;
	logRetalkInfo?: boolean;
	gravatarProxy?: string;
	lang?: "zh-CN" | "wenyan" | "en" | "fr";
}

export type ResolvedOptions = Required<Options>;
