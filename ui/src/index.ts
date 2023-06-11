import "uno.css";

import Retalk from "./retalk";
import { IS_CLIENT } from "./utils";

if (import.meta.env.DEV && IS_CLIENT) {
  window.Retalk = Retalk;
}

export default Retalk;
