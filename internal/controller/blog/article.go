package blog

import (
	"athena-server/api/blog/article"
	"athena-server/internal/model"
	"athena-server/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type Article struct{}

func NewArticle() *Article {
	return &Article{}
}

func (a Article) GetArticleList(ctx context.Context, req *article.GetArticleListReq) (res *article.GetArticleListRes, err error) {
	res = &article.GetArticleListRes{}

	out, err := service.Article().BlogGetArticleList(ctx, &model.BlogGetArticleListInput{})
	if err != nil {
		return res, err
	}

	res.ArticleList = out.ArticleList
	return
}

// GetArticleDetail 获取文章详情页面
func (a Article) GetArticleDetail(ctx context.Context, req *article.GetArticleDetailReq) (res *article.GetArticleDetailRes, err error) {
	res = &article.GetArticleDetailRes{}

	out, err := service.Article().BlogGetArticleDetail(ctx, &model.BlogGetArticleDetailInput{
		Id: g.RequestFromCtx(ctx).Get("id").Int(),
	})
	if err != nil {
		return res, err
	}

	res.Content = out.Content
	return
}

func (a Article) DeleteArticle(ctx context.Context, req *article.GetArticleListReq) {
}
