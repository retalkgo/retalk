package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/retalkgo/retalk/internal/entity"
	"github.com/retalkgo/retalk/internal/query"
	"github.com/retalkgo/retalk/server/common"
)

// @Summary		根据ID删除评论
// @Description	根据ID删除评论
// @Tags			评论
// @Success		200	{object}	common.Resp
// @Failure		403	{object}	common.Resp
// @Failure		500	{object}	common.Resp
// @Security		ApiKeyAuth
// @Router			/api/comment/delete [delete]
func CommentDelete(router fiber.Router) {
	router.Delete("/delete", func(c *fiber.Ctx) error {
		if !common.Auth(c) {
			return common.RespError(c, "Token错误", nil, 403)
		}
		raw_id := c.Query("id")
		if raw_id == "" {
			return common.RespError(c, "请传递评论ID", nil, 400)
		}
		int_id, err := strconv.Atoi(raw_id)
		if err != nil {
			return common.RespServerError(c)
		}
		id := uint(int_id)
		_, err = query.Comment.Delete(&entity.Comment{Base: entity.Base{ID: id}})
		if err != nil {
			return common.RespServerError(c)
		}
		return common.RespSuccess(c, "成功删除", nil)
	})
}
