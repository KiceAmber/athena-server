// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"athena-server/internal/model"
	"athena-server/internal/model/do"
	"context"
)

type (
	ITag interface {
		// GetTagList Get Tag List from Database
		// Now request don not contain page info
		GetTagList(ctx context.Context, in *model.GetTagListInput) (out *model.GetTagListOutput, err error)
		// CreateTag Create one tag from admin
		// check if the tag of name is already exist
		CreateTag(ctx context.Context, in *model.CreateTagInput) (out *model.CreateTagOutput, err error)
		UpdateTag(ctx context.Context, tag *do.Tag) (err error)
		DeleteTag(ctx context.Context) (err error)
	}
)

var (
	localTag ITag
)

func Tag() ITag {
	if localTag == nil {
		panic("implement not found for interface ITag, forgot register?")
	}
	return localTag
}

func RegisterTag(i ITag) {
	localTag = i
}
