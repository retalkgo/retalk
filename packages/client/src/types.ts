export type ElementOrSelector = string | HTMLElement;

export interface Options {
  el: ElementOrSelector;
  server: string;
  logRetalkInfo?: boolean;
}
