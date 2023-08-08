import { useOptions } from "../options";
import { translations } from "./languages";
import type { Translations } from "./types";

export type { Language } from "./languages";

export function useI18n(): Translations {
	const lang = useOptions().lang;

	return lang in translations ? translations[lang] : translations["zh-CN"];
}
