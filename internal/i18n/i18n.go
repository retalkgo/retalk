package i18n

import "github.com/retalkgo/retalk/internal/i18n/languages"

var Translation languages.Translation

func InitI18n(lang string) {
	switch lang {
	case "zh-CN":
		Translation = languages.ZhCN
	case "en":
		Translation = languages.En
	case "zh-wenyan":
		Translation = languages.ZhWenyan
	}
}

func I18n(key string) string {
	switch key {
	case "welcome":
		return Translation.Welcome
	case "processingTime":
		return Translation.ProcessingTime
	case "microsecond":
		return Translation.Microsecond
	case "tokenError":
		return Translation.TokenError
	case "successPostComment":
		return Translation.SuccessPostComment
	case "needCommentID":
		return Translation.NeedCommentID
	case "successDelete":
		return Translation.SuccessDelete
	case "successGetAllComments":
		return Translation.SuccessGetAllComments
	case "successGetComments":
		return Translation.SuccessGetComments
	case "serverInited":
		return Translation.ServerInited
	case "needApiKey":
		return Translation.NeedApiKey
	case "successInit":
		return Translation.SuccessInit
	case "notFound":
		return Translation.NotFound
	}
	return "Text Not Found"
}
