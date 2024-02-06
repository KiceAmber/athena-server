package article

import (
	"athena-server/internal/dao"
	"athena-server/internal/model"
	"athena-server/internal/model/entity"
	"athena-server/internal/service"
	"context"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/v2/database/gdb"
)

type sArticle struct{}

func init() {
	service.RegisterArticle(New())
}

func New() *sArticle {
	return &sArticle{}
}

func (s sArticle) GetArticleList(ctx context.Context, in *model.GetArticleListInput) (out *model.GetArticleListOutput, err error) {
	// init out
	out = &model.GetArticleListOutput{
		ArticleList: make([]*model.ArticleItem, 0),
	}

	articleList := make([]*entity.Article, 0)
	// use transaction to query the articleList
	err = dao.Article.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if err = dao.Article.Ctx(ctx).Scan(&articleList); err != nil {
			return err
		}
		glog.Info("articleList => ", articleList)

		for _, item := range articleList {
			var user = &entity.User{}
			if err = dao.User.Ctx(ctx).Where("id = ?", item.AuthorId).Scan(&user); err != nil {
				return err
			}

			var article = &model.ArticleItem{
				Id:         item.Id,
				Title:      item.Title,
				Content:    item.Content,
				AuthorName: user.Passport,
			}

			out.ArticleList = append(out.ArticleList, article)
		}
		return err
	})
	return
}

func (s sArticle) AddArticle(ctx context.Context, in *model.AddArticleInput) (out *model.AddArticleOutput, err error) {
	return
}
