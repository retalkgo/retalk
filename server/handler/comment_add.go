package handler

import (
	"net/http"

	"github.com/retalkgo/retalk/internal/entity"
	"github.com/retalkgo/retalk/internal/query"
	"github.com/retalkgo/retalk/server/common"

	"github.com/gofiber/fiber/v2"
)

// @Summary		新增评论
// @Description	新增评论
// @Tags			评论
// @Param			path	formData	string		true	"评论路径"
// @Param			name	formData	string		true	"发送者昵称"
// @Param			email	formData	string		true	"发送者邮箱"
// @Param			link	formData	string		true	"发送者网站"
// @Param			body	string		formData	string	true	"正文"
// @Success		200		{object}	common.Resp{}
// @Failure		400		{object}	common.Resp{}
// @Failure		500		{object}	common.Resp{}
// @Router			/api/comment/add [post]
func CommentAdd(router fiber.Router) {
	router.Post("/add", func(c *fiber.Ctx) error {
		path := c.FormValue("path")
		name := c.FormValue("name")
		email := c.FormValue("email")
		link := c.FormValue("link")
		body := c.FormValue("body")
		var author *entity.Author
		author, _ = query.Author.Where(query.Author.Name.Eq(name)).First()
		if author == nil {
			query.Author.Create(&entity.Author{
				Name:  name,
				Email: email,
				Link:  link,
			})
			author, _ = query.Author.Where(query.Author.Name.Eq(name)).First()
		}
		if author.IsAdmin {
			if common.Auth(c) == false {
				return common.RespError(c, "Token错误", nil, http.StatusForbidden)
			}
		}
		comment := query.Comment.Create(&entity.Comment{
			Path:     path,
			AuthorID: author.ID,
			Body:     body,
		})
		return common.RespSuccess(c, "成功发布评论", comment)
	})
}
