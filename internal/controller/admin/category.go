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

func (c Category) DeleteCategory(ctx context.Context, req *category.DeleteCategoryReq) (res *category.DeleteCategoryRes, err error) {
	res = &category.DeleteCategoryRes{}

	_, err = service.Category().AdminDeleteCategory(ctx, &model.AdminDeleteCategoryInput{
		Id: req.Id,
	})

	return
}

func (c Category) UpdateCategory(ctx context.Context, req *category.UpdateCategoryReq) (res *category.UpdateCategoryRes, err error) {
	res = &category.UpdateCategoryRes{}

	_, err = service.Category().AdminUpdateCategory(ctx, &model.AdminUpdateCategoryInput{
		Id:        req.Id,
		Name:      req.Name,
		IsVisible: req.IsVisible,
	})
	if err != nil {
		return nil, err
	}

	return
}
