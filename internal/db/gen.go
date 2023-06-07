package db

import (
	"gorm.io/gen"
	"retalk/internal/entity"
)

// 公用Querier定义
type AllQuerier interface {
}

func Gen() {
	InitDB()
	g := gen.NewGenerator(gen.Config{
		ModelPkgPath: "./internal/query",
		OutPath:      "./internal/query",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(dbInterface)
	g.ApplyBasic(&entity.Comment{})
	g.GenerateAllTable()
	g.ApplyInterface(func(AllQuerier) {}, &entity.Comment{}, &entity.Author{}, &entity.Reply{})
	g.Execute()
}
