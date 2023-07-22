import Retalk from "@retalkgo/client";

import "@retalkgo/client/retalk.css";

const appEl = document.querySelector("#app")!;

// eslint-disable-next-line no-new
new Retalk({
  el: appEl,
  server: "",
});
