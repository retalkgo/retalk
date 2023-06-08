import { render } from "solid-js/web";

import styles from "./styles/index.module.css";

const App = () => <h1 class={styles.demo}>Retalk</h1>;

render(() => <App />, document.querySelector("#retalk")!);
