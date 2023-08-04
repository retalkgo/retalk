import zhCN from './zh-CN.json'
import wenyan from './wenyan.json'
import en from './en.json'
import fr from './fr.json'

import { Options } from '../types.ts'

interface Translations {
    name: string;
    email: string;
    link: string;
    send: string;
    welcome: string;
    admin: string;
    reply: string;
}
  

export default function useI18n(options: Options) {
    if (options.lang == "zh-CN") {
        const translations: Translations = zhCN;
        return translations;
    }
    if (options.lang == "wenyan") {
        const translations: Translations = wenyan;
        return translations;
    }
    if (options.lang == "en") {
        const translations: Translations = en;
        return translations;
    }
    if (options.lang == "fr") {
        const translations: Translations = fr;
        return translations;
    }
    const translations: Translations = zhCN;
    return translations;
}