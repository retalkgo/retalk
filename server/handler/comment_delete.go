package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/retalkgo/retalk/internal/entity"
	"github.com/retalkgo/retalk/internal/i18n"
	"github.com/retalkgo/retalk/internal/query"
	"github.com/retalkgo/retalk/server/common"
)

// @Summary		根据ID删除评论
// @Description	根据ID删除评论
// @Tags			评论
// @Success		200	{object}	common.Resp
// @Failure		403	{object}	common.Resp
// @Failure		500	{object}	common.Resp
// @Param			id	query		string	true	"评论ID"
// @Security		ApiKeyAuth
// @Router			/api/comment/delete [delete]
func CommentDelete(router fiber.Router) {
	router.Delete("/delete", func(c *fiber.Ctx) error {
		if !common.Auth(c) {
			return common.RespError(c, i18n.I18n("tokenError"), nil, http.StatusForbidden)
		}
		raw_id := c.Query("id")
		if raw_id == "" {
			return common.RespError(c, i18n.I18n("needCommentID"), nil, http.StatusBadRequest)
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
		return common.RespSuccess(c, i18n.I18n("successDelete"), nil)
	})
}
