package comment

import (
	"retalk/internal/entity"
	"retalk/internal/logger"
	"retalk/internal/md5"
	"retalk/internal/query"
)

func SuperCommentGet(comments []*entity.Comment) ([]entity.CookedComment, error) {
	rawData, err := query.Comment.Find()
		if err != nil {
			logger.Error("服务器内部错误: " + err.Error())
			return nil, err
		}
		authorMap := make(map[uint]*entity.Author)  // 定义map，用于缓存作者信息
		data := []entity.CookedComment{}
		for _, v := range rawData {
			if author, ok := authorMap[v.AuthorID]; !ok {  // 判断map中是否存在对应的作者信息
				author, err = query.Author.Where(query.Author.ID.Eq(v.AuthorID)).First()
				if err != nil {
					logger.Error("服务器内部错误: " + err.Error())
					return nil, err
				}
				author.Email = md5.MD5(author.Email)
				authorMap[v.AuthorID] = author  // 将作者信息缓存到map中
			}
			cookedComment := &entity.CookedComment{
				ID: v.ID,
				Path: v.Path,
				CreatedAt: v.CreatedAt,
				Body: v.Body,
				Author: *authorMap[v.AuthorID],
			}
			data = append(data, *cookedComment)
		}
		return data, nil
}