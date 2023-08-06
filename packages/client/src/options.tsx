import type { ParentComponent } from "solid-js";
import { createContext, useContext } from "solid-js";

import type { ResolvedOptions } from "./types";

const OptionsContext = createContext<ResolvedOptions>();

export interface OptionsProviderProps {
	options: ResolvedOptions;
}

export const OptionsProvider: ParentComponent<OptionsProviderProps> = (props) => (
	// eslint-disable-next-line solid/reactivity
	<OptionsContext.Provider value={props.options}>
		{props.children}
	</OptionsContext.Provider>
);

export const useOptions = () => useContext(OptionsContext)!;
