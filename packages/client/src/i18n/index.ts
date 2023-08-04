import { useOptions } from "../options";
import { translations } from "./languages";

export type { Language } from "./languages";

interface Translations {
	name: string;
	email: string;
	link: string;
	send: string;
	welcome: string;
	admin: string;
	reply: string;
}

export function useI18n(): Translations {
	const lang = useOptions().lang;

	return lang in translations ? translations[lang] : translations["zh-CN"];
}
