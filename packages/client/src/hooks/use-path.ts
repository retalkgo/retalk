import { createSignal, onCleanup } from "solid-js";

// TODO: SSR friendly
export function usePath() {
	const [path, setPath] = createSignal(location.pathname);

	function updatePath() {
		setPath(location.pathname);
	}

	window.addEventListener("popstate", updatePath);
	onCleanup(() => window.removeEventListener("popstate", updatePath));

	return path;
}
