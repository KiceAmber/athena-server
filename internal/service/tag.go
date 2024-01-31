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
		GetTagList(ctx context.Context, in *model.GetTagListInput) (out *model.GetTagListOutput, err error)
		AddTag(ctx context.Context, in *model.AddTagInput) (out *model.AddTagOutput, err error)
		UpdateTag(ctx context.Context, in *model.UpdateTagInput) (out *model.UpdateTagOutput, err error)
		DeleteTag(ctx context.Context, in *model.DeleteTagInput) (out *model.DeleteTagOutput, err error)
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
