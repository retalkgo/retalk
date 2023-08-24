import type { Api } from "@retalkgo/api";
import type { ParentComponent } from "solid-js";
import { createContext, useContext } from "solid-js";

const ApiContext = createContext<Api>();

export interface ApiProviderProps {
	api: Api;
}

export const ApiProvider: ParentComponent<ApiProviderProps> = (props) => {
	return (
		// eslint-disable-next-line solid/reactivity
		<ApiContext.Provider value={props.api}>
			{props.children}
		</ApiContext.Provider>
	);
};

export const useApi = () => useContext(ApiContext)!;
