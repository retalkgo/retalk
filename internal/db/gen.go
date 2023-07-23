package db

import (
	"github.com/retalkgo/retalk/internal/entity"

	"gorm.io/gen"
)

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
	g.ApplyInterface(func(AllQuerier) {}, &entity.Server{}, &entity.Comment{}, &entity.Author{}, &entity.Reply{})
	g.Execute()
}
