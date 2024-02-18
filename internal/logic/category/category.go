package category

import (
	"athena-server/internal/dao"
	"athena-server/internal/model"
	"athena-server/internal/model/do"
	"athena-server/internal/model/entity"
	"athena-server/internal/service"
	"context"
)

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

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
