package handler

import (
	"retalk/internal/comment"
	"retalk/internal/query"
	"retalk/server/common"

	"github.com/gofiber/fiber/v2"
)

//	@Summary		获取所有评论
//	@Description	获取所有评论
//	@Tags			评论
//	@Success		200	{object}	common.Resp{data=[]entity.CookedComment}
//	@Failure		500	{object}	common.Resp
//	@Router			/api/comment/getAll [get]
func CommentGetAll(router fiber.Router) {
	router.Get("/getAll", func(c *fiber.Ctx) error {
		rawData, _ := query.Comment.Find()
		data, err := comment.SuperCommentGet(rawData)
		if err != nil {
			return common.RespServerError(c)
		}
		return common.RespSuccess(c, "成功获取所有评论", data)
	})
}