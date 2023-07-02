package common

import (
	"retalk/internal/md5"
	"retalk/internal/query"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) bool {
	userToken := c.Get("Authorization")
	userToken = strings.TrimPrefix(userToken, "Bearer ")
	server, err := query.Server.First()
	if err != nil {
		return false
	}
	apikey := server.ApiKey
	if md5.MD5(userToken) != apikey {
		return false
	}
	return true
}
