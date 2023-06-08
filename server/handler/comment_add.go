package handler

import (
	"net/http"
	"retalk/internal/entity"
	"retalk/internal/query"
	"retalk/server/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//	@Summary		新增评论
//	@Description	新增评论
//	@Tags			评论
//	@Param			path		formData	string		true	"评论路径"
//	@Param			author_id	formData	uint		true	"发送者ID"
//	@Param			body		string		formData	string	true	"正文"
//	@Success		200			{object}	common.Resp{data=entity.Comment}
//	@Failure		400			{object}	common.Resp{}
//	@Failure		500			{object}	common.Resp{}
//	@Router			/api/comment/add [post]
func CommentAdd(router fiber.Router) {
	router.Post("/add", func(c *fiber.Ctx) error {
		path := c.FormValue("path")
		authorID, err := strconv.ParseUint(c.FormValue("author_id"), 10, 0)
		body := c.FormValue("body")
		author, err := query.Author.Where(query.Author.ID.Eq(uint(authorID))).First()
		if author == nil {
			return common.RespError(c, "无法找到对应的Author", nil, http.StatusBadRequest)
		}
		comment := query.Comment.Create(&entity.Comment{
			Path: path,
			AuthorID: uint(authorID),
			Body: body,
		})
		if err != nil {
			return common.RespServerError(c)
		}
		return common.RespSuccess(c, "成功发布评论", comment)
	})
}