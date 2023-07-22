import Retalk from "./retalk";
import { IS_CLIENT } from "./utils";

import "virtual:uno.css";

if (import.meta.env.DEV && IS_CLIENT) {
	window.Retalk = Retalk;
}

export default Retalk;
