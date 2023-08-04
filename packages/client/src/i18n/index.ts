import { useOptions } from "../options";
import en from "./en.json";
import fr from "./fr.json";
import wenyan from "./wenyan.json";
import zhCN from "./zh-CN.json";

interface Translations {
	name: string;
	email: string;
	link: string;
	send: string;
	welcome: string;
	admin: string;
	reply: string;
}

const translations = {
	"zh-CN": zhCN,
	wenyan,
	en,
	fr,
};

export type Language = keyof typeof translations;

export default (): Translations => {
	const lang = useOptions().lang;

	return lang in translations ? translations[lang] : translations["zh-CN"];
};
