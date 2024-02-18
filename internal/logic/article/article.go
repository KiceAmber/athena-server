package article

import (
	"athena-server/internal/dao"
	"athena-server/internal/model"
	"athena-server/internal/model/do"
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

// Admin Service

func (s sArticle) AdminGetArticleList(ctx context.Context, in *model.AdminGetArticleListInput) (out *model.AdminGetArticleListOutput, err error) {
	out = &model.AdminGetArticleListOutput{
		ArticleList: make([]*model.ArticleItem, 0),
	}

	articleList := make([]*entity.Article, 0)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
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
				Id:          articleItem.Id,
				Title:       articleItem.Title,
				Content:     articleItem.Content,
				Image:       articleItem.Image,
				AuthorName:  user.Passport,
				TagList:     tagNameList,
				Description: articleItem.Description,
				IsVisible:   articleItem.IsVisible,
				CreatedAt:   articleItem.CreatedAt,
			}
			out.ArticleList = append(out.ArticleList, article)
		}
		return nil
	})
	return
}

func (s sArticle) AdminAddArticle(ctx context.Context, in *model.AdminAddArticleInput) (out *model.AdminAddArticleOutput, err error) {
	out = &model.AdminAddArticleOutput{}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 先插入文章表
		article := &do.Article{
			Title:       in.Title,
			Content:     in.Content,
			Description: in.Description,
			IsVisible:   in.IsVisible,
			AuthorId:    in.AuthorId,
		}
		r, err := dao.Article.Ctx(ctx).Save(article)
		if err != nil {
			return err
		}

		// 再插入文章标签表
		articleId, err := r.LastInsertId()
		if err != nil {
			return err
		}
		for _, tagId := range in.TagList {
			articleTag := &entity.ArticleTag{
				ArticleId: int(articleId),
				TagId:     tagId,
			}
			_, err := dao.ArticleTag.Ctx(ctx).Save(articleTag)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return
	}

	return
}

func (s sArticle) AdminDeleteArticle(ctx context.Context, in *model.AdminDeleteArticleInput) (out *model.AdminDeleteArticleOutput, err error) {
	out = &model.AdminDeleteArticleOutput{}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除文章表的记录
		_, err := dao.Article.Ctx(ctx).Where("id = ?", in.Id).Delete()
		if err != nil {
			return err
		}

		// 删除文章标签关联表的记录
		_, err = dao.ArticleTag.Ctx(ctx).Where("article_id = ?", in.Id).Delete()
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return out, err
	}

	return
}

// Blog Service

func (s sArticle) BlogGetArticleList(ctx context.Context, in *model.BlogGetArticleListInput) (out *model.BlogGetArticleListOutput, err error) {
	out = &model.BlogGetArticleListOutput{
		ArticleList: make([]*model.ArticleItem, 0),
	}

	articleList := make([]*entity.Article, 0)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 查询文章表
		if err = dao.Article.Ctx(ctx).Scan(&articleList); err != nil {
			return err
		}

		for _, articleItem := range articleList {
			// 查询文章标签表
			tagList := make([]*entity.Tag, 0)

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
				Id:          articleItem.Id,
				Title:       articleItem.Title,
				Content:     articleItem.Content,
				Image:       articleItem.Image,
				AuthorName:  user.Passport,
				TagList:     tagNameList,
				Description: articleItem.Description,
				IsVisible:   articleItem.IsVisible,
				CreatedAt:   articleItem.CreatedAt,
			}
			out.ArticleList = append(out.ArticleList, article)
		}
		return nil
	})
	return
}
