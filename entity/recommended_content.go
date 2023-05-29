package entity

import (
	"github.com/uptrace/bun"
	"reflect"
)

type RecommendedContent struct {
	bun.BaseModel `bun:"table:recommended_content"`
	ID            int64 `bun:",pk" json:"id"`
	ContentID     int64 `json:"content_id"`
	RecommendedID int64 `json:"recommended_id"`
}

func (r RecommendedContent) IsEmpty() bool {
	return reflect.DeepEqual(r, RecommendedContent{})
}
