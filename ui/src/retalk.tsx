import { render } from "solid-js/web";

import styles from "./styles/index.module.css";
import { logRetalkInfo, resolveElement } from "./utils";
import type { ElementOrSelector, Options } from "./types";

const App = () => <h1 class={styles.demo}>Retalk</h1>;

export default class Retalk {
  constructor(el: ElementOrSelector, _server: string, options: Options = {}) {
    const resolvedEl = resolveElement(el);
    if (!resolvedEl) {
      throw new Error(`Retalk: Element ${el as string} not found`);
    }
    render(() => <App />, resolvedEl);
    options.logRetalkInfo && logRetalkInfo();
  }
}
