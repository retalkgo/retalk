import en from "./en";
import fr from "./fr";
import zhCN from "./zh-CN";
import zhWenyan from "./zh-wenyan";

export const translations = {
	"zh-CN": zhCN,
	"zh-wenyan": zhWenyan,
	en,
	fr,
};

export type Language = keyof typeof translations;
