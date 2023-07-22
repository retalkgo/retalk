import Retalk from "@retalkgo/client";

import "@retalkgo/client/retalk.css";

// eslint-disable-next-line no-new
new Retalk({
  el: document.querySelector("#app")!,
  // TODO
  server: "",
});
