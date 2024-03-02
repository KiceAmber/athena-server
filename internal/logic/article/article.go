package article

import (
	"athena-server/internal/dao"
	"athena-server/internal/model"
	"athena-server/internal/model/do"
	"athena-server/internal/model/entity"
	"athena-server/internal/service"
	"context"
	"fmt"

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
			user := &entity.User{}
			if err = dao.User.Ctx(ctx).Where("id = ?", articleItem.AuthorId).Scan(&user); err != nil {
				return err
			}

			article := &model.ArticleItem{
				Id:          articleItem.Id,
				Title:       articleItem.Title,
				Content:     articleItem.Content,
				Cover:       articleItem.Cover,
				AuthorName:  user.Passport,
				TagList:     tagNameList,
				Description: articleItem.Description,
				IsVisible:   articleItem.IsVisible,
				CreatedAt:   articleItem.CreatedAt,
			}

			// 查询文章分类，但如果分类 ID 为 0，则说明该文章没有分类
			categoryItem := &entity.Category{}
			if articleItem.CategoryId != 0 {
				if err = dao.Category.Ctx(ctx).Where("id = ?", articleItem.CategoryId).Scan(&categoryItem); err != nil {
					return err
				}
			} else {
				article.CategoryName = "无分类"
			}

			out.ArticleList = append(out.ArticleList, article)
		}
		return nil
	})
	fmt.Println(err)

	if len(articleList) == 0 {
		return out, nil
	}
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
			CategoryId:  in.CategoryId,
			// Cover: in.Cover,
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
			user := &entity.User{}
			if err = dao.User.Ctx(ctx).Where("id = ?", articleItem.AuthorId).Scan(&user); err != nil {
				return err
			}

			article := &model.ArticleItem{
				Id:          articleItem.Id,
				Title:       articleItem.Title,
				Content:     articleItem.Content,
				Cover:       articleItem.Cover,
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

func (s sArticle) BlogGetArticleDetail(ctx context.Context, in *model.BlogGetArticleDetailInput) (out *model.BlogGetArticleDetailOutput, err error) {
	out = &model.BlogGetArticleDetailOutput{}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 获取文章的详情，也就是 content 的数据
		articleItem := &entity.Article{}
		if err = dao.Article.Ctx(ctx).Where("id = ?", in.Id).Scan(&articleItem); err != nil {
			return err
		}

		out.Content = articleItem.Content
		return nil
	}); err != nil {
		return
	}

	return
}
