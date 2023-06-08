package handler

import (
	"retalk/internal/entity"
	"retalk/internal/logger"
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
		rawData, err := query.Comment.Find()
		if err != nil {
			logger.Error("服务器内部错误: " + err.Error())
			return common.RespServerError(c)
		}
		authorMap := make(map[uint]*entity.Author)  // 定义map，用于缓存作者信息
		data := []entity.CookedComment{}
		for _, v := range rawData {
			if author, ok := authorMap[v.AuthorID]; !ok {  // 判断map中是否存在对应的作者信息
				author, err = query.Author.Where(query.Author.ID.Eq(v.AuthorID)).First()
				if err != nil {
					logger.Error("服务器内部错误: " + err.Error())
					return common.RespServerError(c)
				}
				authorMap[v.AuthorID] = author  // 将作者信息缓存到map中
			}
			cookedComment := &entity.CookedComment{
				Path: v.Path,
				CreatedAt: v.CreatedAt,
				Body: v.Body,
				Author: *authorMap[v.AuthorID],
			}
			data = append(data, *cookedComment)
		}
		return common.RespSuccess(c, "成功获取所有评论", data)
	})
}