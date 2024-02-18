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

func (a Article) GetArticleList(ctx context.Context, req *article.GetArticleListReq) (res *article.GetArticleListRes, err error) {
	res = &article.GetArticleListRes{}

	out, err := service.Article().GetArticleList(ctx, &model.AdminGetArticleListInput{})
	if err != nil {
		return res, err
	}

	res.ArticleList = out.ArticleList
	return
}

func (a Article) AddArticle(ctx context.Context, req *article.AddArticleReq) (res *article.AddArticleRes, err error) {
	res = &article.AddArticleRes{}

	_, err = service.Article().AddArticle(ctx, &model.AdminAddArticleInput{
		Title:       req.Title,
		Content:     req.Content,
		Description: req.Description,
		IsVisible:   req.IsVisible,
		AuthorId:    req.AuthorId,
		TagList:     req.TagList,
	})

	return
}

func (a Article) DeleteArticle(ctx context.Context, req *article.DeleteArticleReq) (res *article.DeleteArticleRes, err error) {
	res = &article.DeleteArticleRes{}

	_, err = service.Article().DeleteArticle(ctx, &model.AdminDeleteArticleInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}

	return
}
