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
	IArticle interface {
		GetArticleList(ctx context.Context, in *model.GetArticleListInput) (out *model.GetArticleListOutput, err error)
		AddArticle(ctx context.Context, in *model.AddArticleInput) (out *model.AddArticleOutput, err error)
		DeleteArticle(ctx context.Context, in *model.DeleteArticleInput) (out *model.DeleteArticleOutput, err error)
	}
)

var (
	localArticle IArticle
)

func Article() IArticle {
	if localArticle == nil {
		panic("implement not found for interface IArticle, forgot register?")
	}
	return localArticle
}

func RegisterArticle(i IArticle) {
	localArticle = i
}
