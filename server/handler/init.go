package handler

import (
	"net/http"

	"github.com/retalkgo/retalk/internal/entity"
	"github.com/retalkgo/retalk/internal/md5"
	"github.com/retalkgo/retalk/internal/query"
	"github.com/retalkgo/retalk/server/common"

	"github.com/gofiber/fiber/v2"
)

type RespInit struct {
	IsInit bool
}

//	@Summary		初始化服务端
//	@Description	初始化服务端, 创建ApiKey
//	@Tags			服务端
//	@Param			apikey	formData	string	true	"ApiKey"
//	@Param			name	formData	string	true	"管理员昵称"
//	@Param			email	formData	string	true	"管理员邮箱"
//	@Param			link	formData	string	true	"管理员网站"
//	@Success		200		{object}	common.Resp{data=RespInit}
//	@Success		403		{object}	common.Resp{data=RespInit}
//	@Failure		500		{object}	common.Resp
//	@Router			/api/init [post]
func Init(router fiber.Router) {
	router.Post("/init", func(c *fiber.Ctx) error {
		AllServers, err := query.Server.Find()
		if len(AllServers) != 0 {
			return common.RespError(c, "已经被初始化", nil, http.StatusForbidden)
		}
		apikey := c.FormValue("apikey")
		if apikey == "" {
			return common.RespError(c, "请填写ApiKey", nil, http.StatusBadRequest)
		}
		apikey = md5.MD5(apikey)
		server := &entity.Server{
			IsInit: true,
			ApiKey: apikey,
		}
		err = query.Server.Save(server)
		if err != nil {
			return common.RespServerError(c)
		}
		admin := &entity.Author{
			Name:    c.FormValue("name"),
			Email:   c.FormValue("email"),
			Link:    c.FormValue("link"),
			IsAdmin: true,
		}
		err = query.Author.Save(admin)
		if err != nil {
			return common.RespServerError(c)
		}
		return common.RespSuccess(c, "成功初始化", &RespInit{
			IsInit: true,
		})
	})
}
