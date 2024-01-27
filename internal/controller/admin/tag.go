package admin

import (
	"athena-server/api/admin/tag"
	"context"
)

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) CreateTag(ctx context.Context, req *tag.CreateTagReq) (res *tag.CreateTagRes, err error) {

	// database operations
	return &tag.CreateTagRes{}, nil
}

// func (t *Tag) UpdateTag(ctx context.Context, req *tag.UpdateTagReq) (res *tag.UpdateTagRes, err error) {

// 	return
// }
