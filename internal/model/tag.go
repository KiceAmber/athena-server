package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type TagItem struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	CreatedAt *gtime.Time `json:"createdAt"`
}

type GetTagListInput struct {
}

type GetTagListOutput struct {
	TagList []*TagItem `json:"tagList"`
}

type AddTagInput struct {
	Name string `json:"name"`
}

type AddTagOutput struct {
}

type DeleteTagInput struct {
	Id int `json:"id"`
}

type DeleteTagOutput struct{}

type UpdateTagInput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateTagOutput struct{}
