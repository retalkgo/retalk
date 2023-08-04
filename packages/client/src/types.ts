import type { MountableElement } from "solid-js/web";

import type { Language } from "./i18n";

export type ElementOrSelector = string | MountableElement;

export interface Options {
	el: ElementOrSelector;
	server: string;
	logRetalkInfo?: boolean;
	gravatarProxy?: string;
	lang?: Language;
}

export type ResolvedOptions = Required<Options>;
