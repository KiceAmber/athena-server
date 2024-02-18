// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"athena-server/internal/model"
	"context"
)

type (
	ICategory interface {
		AdminAddCategory(ctx context.Context, in *model.AdminAddCategoryInput) (out *model.AdminAddCategoryOutput, err error)
		AdminGetCategoryList(ctx context.Context, in *model.AdminGetCategoryListInput) (out *model.AdminGetCategoryListOutput, err error)
	}
)

var (
	localCategory ICategory
)

func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

func RegisterCategory(i ICategory) {
	localCategory = i
}
