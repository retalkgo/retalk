import en from "./en.json";
import fr from "./fr.json";
import wenyan from "./wenyan.json";
import zhCN from "./zh-CN.json";

export const translations = {
	"zh-CN": zhCN,
	wenyan,
	en,
	fr,
};

export type Language = keyof typeof translations;
