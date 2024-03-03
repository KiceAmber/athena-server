package admin

import (
	"athena-server/api/admin/article"
	"athena-server/internal/model"
	"athena-server/internal/service"
	"context"
	"fmt"
)

type Article struct{}

func NewArticle() *Article {
	return &Article{}
}

func (a Article) GetArticleList(ctx context.Context, req *article.GetArticleListReq) (res *article.GetArticleListRes, err error) {
	res = &article.GetArticleListRes{}

	out, err := service.Article().AdminGetArticleList(ctx, &model.AdminGetArticleListInput{})
	if err != nil {
		fmt.Println("====>", out.ArticleList)
		return res, err
	}

	res.ArticleList = out.ArticleList
	return
}

func (a Article) AddArticle(ctx context.Context, req *article.AddArticleReq) (res *article.AddArticleRes, err error) {
	res = &article.AddArticleRes{}

	_, err = service.Article().AdminAddArticle(ctx, &model.AdminAddArticleInput{
		Title:       req.Title,
		Content:     req.Content,
		Description: req.Description,
		IsVisible:   req.IsVisible,
		AuthorId:    req.AuthorId,
		TagList:     req.TagList,
		CategoryId:  req.CategoryId,
		// Cover: req.Cover,
	})

	return
}

func (a Article) DeleteArticle(ctx context.Context, req *article.DeleteArticleReq) (res *article.DeleteArticleRes, err error) {
	res = &article.DeleteArticleRes{}

	_, err = service.Article().AdminDeleteArticle(ctx, &model.AdminDeleteArticleInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}

	return
}

func (a Article) GetArticleDetail(ctx context.Context, req *article.GetArticleDetailReq) (res *article.GetArticleDetailRes, err error) {
	res = &article.GetArticleDetailRes{}

	out, err := service.Article().AdminGetArticleDetail(ctx, &model.AdminGetArticleDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return
	}
	res.Article = out.Articel

	return
}

func (a Article) UpdateArticle(ctx context.Context, req *article.UpdateArticleReq) (res *article.UpdateArticleRes, err error) {
	res = &article.UpdateArticleRes{}

	_, err = service.Article().AdminUpdateArticle(ctx, &model.AdminUpdateArticleInput{
		Id:          req.Id,
		Title:       req.Title,
		Content:     req.Content,
		Description: req.Description,
		AuthorId:    req.AuthorId,
		TagList:     req.TagList,
		CategoryId:  req.CategoryId,
		IsVisible:   req.IsVisible,
	})

	return
}
