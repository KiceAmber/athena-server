package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type TagItem struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	IsVisible int         `json:"isVisible"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
	DeletedAt *gtime.Time `json:"deletedAt"`
}

type AdminGetTagListInput struct {
}

type AdminGetTagListOutput struct {
	TagList []*TagItem `json:"tagList"`
}

type AdminAddTagInput struct {
	Name      string `json:"name"`
	IsVisible int    `json:"isVisible"`
}

type AdminAddTagOutput struct {
}

type AdminDeleteTagInput struct {
	Id int `json:"id"`
}

type AdminDeleteTagOutput struct{}

type AdminUpdateTagInput struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsVisible int    `json:"isVisible"`
}

type AdminUpdateTagOutput struct{}
