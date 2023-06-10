import { render } from "solid-js/web";
import packageJSON from "../package.json"
import versionImg from "../assets/version.svg"

import styles from "./styles/index.module.css";


const App = () => <h1 class={styles.demo}>Retalk</h1>;


console.log(`%c Retalk %c ${packageJSON.version} `, 'background: #006BB8; padding: 5px; border-radius: 3px 0 0 3px; color: #fff', 'background: #006BB818; padding: 5px; border-radius: 0 3px 3px 0; color: #006BB8')
render(() => <App />, document.querySelector("#retalk")!);