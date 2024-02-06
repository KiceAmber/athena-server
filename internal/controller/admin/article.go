package admin

import (
	"athena-server/api/admin/article"
	"athena-server/internal/model"
	"athena-server/internal/service"
	"context"
)

type Article struct{}

func NewArticle() *Article {
	return &Article{}
}

func (a *Article) GetArticleList(ctx context.Context, req *article.GetArticleListReq) (res *article.GetArticleListRes, err error) {

	// init res
	res = &article.GetArticleListRes{}

	out, err := service.Article().GetArticleList(ctx, &model.GetArticleListInput{})
	if err != nil {
		return res, err
	}

	res.ArticleList = out.ArticleList
	return
}

func (a *Article) AddArticle(ctx context.Context, req *article.AddArticleReq) (res *article.AddArticleRes, err error) {

	return
}
