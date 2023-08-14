package handler

import (
	"github.com/retalkgo/retalk/internal/comment"
	"github.com/retalkgo/retalk/internal/i18n"
	"github.com/retalkgo/retalk/internal/query"
	"github.com/retalkgo/retalk/server/common"

	"github.com/gofiber/fiber/v2"
)

// @Summary		获取所有评论
// @Description	获取所有评论
// @Tags			评论
// @Success		200	{object}	common.Resp{data=[]entity.CookedComment}
// @Failure		500	{object}	common.Resp
// @Router			/api/comment/getAll [get]
func CommentGetAll(router fiber.Router) {
	router.Get("/getAll", func(c *fiber.Ctx) error {
		rawData, _ := query.Comment.Find()
		data, err := comment.SuperCommentGet(rawData)
		if err != nil {
			return common.RespServerError(c)
		}
		return common.RespSuccess(c, i18n.I18n("successGetAllComments"), data)
	})
}
