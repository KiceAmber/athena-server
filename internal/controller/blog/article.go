package blog

import (
	"athena-server/api/blog/article"
	"athena-server/internal/model"
	"athena-server/internal/service"
	"context"
)

type Article struct{}

func NewArticle() *Article {
	return &Article{}
}

func (a Article) GetArticleList(ctx context.Context, req *article.GetArticleListReq) (res *article.GetArticleListRes, err error) {
	res = &article.GetArticleListRes{}

	out, err := service.Article().GetArticleList(ctx, &model.GetArticleListInput{})
	if err != nil {
		return res, err
	}

	res.ArticleList = out.ArticleList
	return
}
