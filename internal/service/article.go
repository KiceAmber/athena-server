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
		AdminGetArticleList(ctx context.Context, in *model.AdminGetArticleListInput) (out *model.AdminGetArticleListOutput, err error)
		AdminAddArticle(ctx context.Context, in *model.AdminAddArticleInput) (out *model.AdminAddArticleOutput, err error)
		AdminDeleteArticle(ctx context.Context, in *model.AdminDeleteArticleInput) (out *model.AdminDeleteArticleOutput, err error)
		AdminGetArticleDetail(ctx context.Context, in *model.AdminGetArticleDetailInput) (out *model.AdminGetArticleDetailOutput, err error)
		AdminUpdateArticle(ctx context.Context, in *model.AdminUpdateArticleInput) (out *model.AdminUpdateArticleOutput, err error)
		BlogGetArticleList(ctx context.Context, in *model.BlogGetArticleListInput) (out *model.BlogGetArticleListOutput, err error)
		BlogGetArticleDetail(ctx context.Context, in *model.BlogGetArticleDetailInput) (out *model.BlogGetArticleDetailOutput, err error)
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
