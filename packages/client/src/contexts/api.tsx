import { Api } from "@retalkgo/api";
import type { ParentComponent } from "solid-js";
import { createContext, useContext } from "solid-js";

import { useOptions } from "./options";

const ApiContext = createContext<Api>();

export const ApiProvider: ParentComponent = (props) => {
	const options = useOptions();
	const api = new Api({ baseUrl: options.server });

	return (
		<ApiContext.Provider value={api}>{props.children}</ApiContext.Provider>
	);
};

export const useApi = () => useContext(ApiContext)!;
