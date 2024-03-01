package category

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

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

// admin

func (s sCategory) AdminGetCategoryList(ctx context.Context, in *model.AdminGetCategoryListInput) (out *model.AdminGetCategoryListOutput, err error) {
	out = &model.AdminGetCategoryListOutput{
		CategoryList: []*model.CategoryItem{},
	}

	// 这里要查询创建分类的用户名称，这里暂定为 Amber
	authorName := "Amber"

	categoryList := []*entity.Category{}
	if err = dao.Category.Ctx(ctx).Scan(&categoryList); err != nil {
		return nil, err
	}

	for _, item := range categoryList {
		category := &model.CategoryItem{
			Id:         item.Id,
			Name:       item.Name,
			IsVisible:  item.IsVisible,
			AuthorName: authorName,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			DeletedAt:  item.DeletedAt,
		}
		out.CategoryList = append(out.CategoryList, category)
	}

	return
}

func (s sCategory) AdminAddCategory(ctx context.Context, in *model.AdminAddCategoryInput) (out *model.AdminAddCategoryOutput, err error) {
	out = &model.AdminAddCategoryOutput{}

	_, err = dao.Category.Ctx(ctx).Insert(&do.Category{
		Name:      in.Name,
		IsVisible: in.IsVisible,
		AuthorId:  1,
	})
	if err != nil {
		return nil, err
	}

	return
}

func (s sCategory) AdminDeleteCategory(ctx context.Context, in *model.AdminDeleteCategoryInput) (out *model.AdminDeleteCategoryOutput, err error) {
	out = &model.AdminDeleteCategoryOutput{}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		// 删除该分类
		_, err = dao.Category.Ctx(ctx).Unscoped().Delete("id = ?", in.Id)
		if err != nil {
			return err
		}

		// 删除与该分类关联的所有文章，将该文章的 category_id 设置为 0
		_, err = dao.Article.Ctx(ctx).Where("category_id = ?", in.Id).Update(&do.Article{
			CategoryId: 0,
		})
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return
	}

	return
}

func (s sCategory) AdminUpdateCategory(ctx context.Context, in *model.AdminUpdateCategoryInput) (out *model.AdminUpdateCategoryOutput, err error) {
	out = &model.AdminUpdateCategoryOutput{}

	_, err = dao.Category.Ctx(ctx).Where("id = ?", in.Id).Update(&do.Category{
		Name:      in.Name,
		IsVisible: in.IsVisible,
	})
	if err != nil {
		return nil, err
	}

	return
}
