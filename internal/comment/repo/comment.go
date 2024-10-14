package repo

import (
	"commentService/internal/comment/dao/query"
)

type Comment struct {
	*query.Query
}

func NewComment(query *query.Query) *Comment {
	return &Comment{Query: query}
}
