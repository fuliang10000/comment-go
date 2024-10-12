package comment

import (
	"commentService/app/config"
	"commentService/internal/comment/dao/model"
	"commentService/internal/comment/dao/query"
	"commentService/pkg/gormx"
)

//go:generate go run ../../app/tools/sql/main.go -sql ../../develop/sql/comment.sql

func QueryInit(c *config.Config) (*query.Query, error) {
	db, err := gormx.Init(&c.Mysql)
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(
		&model.CommentContent{},
		&model.CommentIndex{},
		&model.CommentTheme{},
	); err != nil {
		return nil, err
	}
	query.SetDefault(db)
	return query.Use(db), nil
}
