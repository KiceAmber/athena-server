package admin

import (
	"athena-server/api/admin/category"
	"athena-server/internal/model"
	"athena-server/internal/service"
	"context"
)

type Category struct{}

func NewCategory() *Category {
	return &Category{}
}

func (c Category) GetCategoryList(ctx context.Context, req *category.GetCategoryListReq) (res *category.GetCategoryListRes, err error) {
	res = &category.GetCategoryListRes{}

	out, err := service.Category().AdminGetCategoryList(ctx, &model.AdminGetCategoryListInput{})
	if err != nil {
		return res, err
	}

	res.CategoryList = out.CategoryList
	return
}

func (c Category) AddCategory(ctx context.Context, req *category.AddCategoryReq) (res *category.AddCategoryRes, err error) {
	res = &category.AddCategoryRes{}

	_, err = service.Category().AdminAddCategory(ctx, &model.AdminAddCategoryInput{
		Name:      req.Name,
		IsVisible: req.IsVisible,
	})
	if err != nil {
		return nil, err
	}

	return
}
