package model

import "github.com/gogf/gf/v2/os/gtime"

type CategoryItem struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	IsVisible  int         `json:"isVisible"`
	AuthorName string      `json:"authorName"`
	CreatedAt  *gtime.Time `json:"createdAt"`
	UpdatedAt  *gtime.Time `json:"updatedAt"`
	DeletedAt  *gtime.Time `json:"deletedAt"`
}

type AdminAddCategoryInput struct {
	Name      string `json:"name"`
	IsVisible int    `json:"isVisible"`
}

type AdminAddCategoryOutput struct{}

type AdminGetCategoryListInput struct {
}

type AdminGetCategoryListOutput struct {
	CategoryList []*CategoryItem `json:"categoryList"`
}
