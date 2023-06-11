import { render } from "solid-js/web";

import RetalkComponent from "./components/Retalk";
import { logRetalkInfo, resolveElement } from "./utils";
import type { Options } from "./types";

export default class Retalk {
  constructor({ el, ...options }: Options) {
    const resolvedEl = resolveElement(el);
    if (!resolvedEl) {
      throw new Error(`Retalk: Element ${el as string} not found`);
    }
    render(() => <RetalkComponent />, resolvedEl);
    options.logRetalkInfo && logRetalkInfo();
  }
}
