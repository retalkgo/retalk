import { rm } from "node:fs/promises";

import { execa } from "execa";

(async () => {
	try {
		await rm("dist", { recursive: true });
		await execa("pnpm", ["run", "build:real"]);
		// Cleanup extra files
		await rm("dist/components", { recursive: true });
		await rm("dist/index.d.ts");
		await rm("dist/utils.d.ts");
	} catch {}
})();
