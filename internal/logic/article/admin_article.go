package article

import (
	"athena-server/internal/dao"
	"athena-server/internal/model"
	"athena-server/internal/model/entity"
	"athena-server/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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
	err = dao.Article.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 查询文章表
		if err = dao.Article.Ctx(ctx).Scan(&articleList); err != nil {
			return err
		}

		for _, articleItem := range articleList {
			// 查询文章标签表
			tagList := make([]*entity.Tag, 0)
			// select * from article_tag
			if err := g.Model("tag t").
				LeftJoin("article_tag at", "t.id = at.tag_id").
				Fields("t.id, t.name").
				Where("at.article_id = ?", articleItem.Id).
				Scan(&tagList); err != nil {
				return err
			}

			tagNameList := make([]string, 0)
			for _, tag := range tagList {
				tagNameList = append(tagNameList, tag.Name)
			}

			// 查询文章作者
			var user = &entity.User{}
			if err = dao.User.Ctx(ctx).Where("id = ?", articleItem.AuthorId).Scan(&user); err != nil {
				return err
			}

			var article = &model.ArticleItem{
				Id:         articleItem.Id,
				Title:      articleItem.Title,
				Content:    articleItem.Content,
				Image:      articleItem.Image,
				AuthorName: user.Passport,
				TagList:    tagNameList,
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
