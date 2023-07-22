export type ElementOrSelector = string | Element;

export interface Options {
	el: ElementOrSelector;
	server: string;
	logRetalkInfo?: boolean;
}
