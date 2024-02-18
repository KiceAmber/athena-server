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
		GetArticleList(ctx context.Context, in *model.AdminGetArticleListInput) (out *model.AdminGetArticleListOutput, err error)
		AddArticle(ctx context.Context, in *model.AdminAddArticleInput) (out *model.AdminAddArticleOutput, err error)
		DeleteArticle(ctx context.Context, in *model.AdminDeleteArticleInput) (out *model.AdminDeleteArticleOutput, err error)
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
