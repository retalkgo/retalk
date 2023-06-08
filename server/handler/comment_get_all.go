package handler

import (
	"retalk/internal/entity"
	"retalk/internal/logger"
	"retalk/internal/query"
	"retalk/server/common"

	"github.com/gofiber/fiber/v2"
)

type RespGetComment struct {
	Total int `json:"total"`
	Comment *[]entity.Comment `json:"comment"`
}

//	@Summary		获取所有评论
//	@Description	获取所有评论
//	@Tags			评论
//	@Success		200	{object}	common.Resp{data=RespGetComment}
//	@Failure		500	{object}	common.Resp
//	@Router			/api/comment/getAll [get]
func CommentGetAll(router fiber.Router) {
	router.Get("/getAll", func(c *fiber.Ctx) error {
		data, err := query.Comment.Find()
		if err != nil {
			logger.Error("服务器内部错误: " + err.Error())
			return common.RespServerError(c)
		}
		return common.RespSuccess(c, "成功获取所有评论", data)
	})
}