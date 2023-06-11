import { render } from "solid-js/web";

import packageJSON from "../package.json";

import styles from "./styles/index.module.css";

const App = () => <h1 class={styles.demo}>Retalk</h1>;

export default class Retalk {
  constructor(config: { el: string; server: string }) {
    render(() => <App />, document.querySelector(config.el)!);
    // eslint-disable-next-line no-console
    console.log(
      `%c Retalk %c ${packageJSON.version} `,
      "background: #006BB8; padding: 5px; border-radius: 3px 0 0 3px; color: #fff",
      "background: #006BB818; padding: 5px; border-radius: 0 3px 3px 0; color: #006BB8",
    );
  }
}
