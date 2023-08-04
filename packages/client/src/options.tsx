import type { ParentComponent } from "solid-js";
import { createContext, useContext } from "solid-js";

import type { ResolvedOptions } from "./types";

const OptionsContext = createContext<ResolvedOptions>();

export interface OptionsProviderProps {
	options: any;
}

export const OptionsProvider: ParentComponent<{
	options: ResolvedOptions;
}> = (props) => (
	// eslint-disable-next-line solid/reactivity
	<OptionsContext.Provider value={props.options}>
		{props.children}
	</OptionsContext.Provider>
);

export const useOptions = () => useContext(OptionsContext)!;
