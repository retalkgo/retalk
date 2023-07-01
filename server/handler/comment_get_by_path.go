package handler

import (
	"retalk/internal/comment"
	"retalk/internal/query"
	"retalk/server/common"

	"github.com/gofiber/fiber/v2"
)

//	@Summary		根据路径获取评论
//	@Description	根据路径获取评论
//	@Tags			评论
//	@Params			path query string true "路径"
//	@Success		200	{object}	common.Resp{data=[]entity.CookedComment}
//	@Failure		500	{object}	common.Resp
//	@Router			/api/comment/getByPath [get]
func CommentGetByPath(router fiber.Router) {
	router.Get("/getByPath", func(c *fiber.Ctx) error {
		rawData, _ := query.Comment.Where(query.Comment.Path.Eq(c.Query("path"))).Find()
		data, err := comment.SuperCommentGet(rawData)
		if err != nil {
			return common.RespServerError(c)
		}
		return common.RespSuccess(c, "成功获取评论", data)
	})
}