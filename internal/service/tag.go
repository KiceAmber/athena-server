// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"athena-server/internal/model"
	"context"
)

type (
	ITag interface {
		AdminGetTagList(ctx context.Context, in *model.AdminGetTagListInput) (out *model.AdminGetTagListOutput, err error)
		AdminAddTag(ctx context.Context, in *model.AdminAddTagInput) (out *model.AdminAddTagOutput, err error)
		AdminUpdateTag(ctx context.Context, in *model.AdminUpdateTagInput) (out *model.AdminUpdateTagOutput, err error)
		AdminDeleteTag(ctx context.Context, in *model.AdminDeleteTagInput) (out *model.AdminDeleteTagOutput, err error)
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
