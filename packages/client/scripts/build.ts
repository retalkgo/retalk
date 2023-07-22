import fs from "node:fs/promises";

import { execa } from "execa";

(async () => {
  await execa("pnpm", ["run", "build:real"]);
  // Cleanup extra files
  try {
    await fs.rm("dist/components", { recursive: true });
    await fs.rm("dist/index.d.ts");
    await fs.rm("dist/utils.d.ts");
  } catch {}
})();
