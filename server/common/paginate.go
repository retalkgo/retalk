package common

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// 分页查询
func Paginate(c *fiber.Ctx, data []*interface{}) []*interface{} {
	pageSize := 10
	page, err := strconv.Atoi(c.Params("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	// 计算总页数
	totalPages := (len(data) + pageSize - 1) / pageSize

	// 判断当前页是否超过总页数
	if page > totalPages {
		page = totalPages
	}

	// 计算起始索引和结束索引
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > len(data) {
		endIndex = len(data)
	}

	return data[startIndex:endIndex]

}
