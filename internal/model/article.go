package model

import "github.com/gogf/gf/v2/os/gtime"

type ArticleItem struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	AuthorName  string      `json:"authorName"`
	Content     string      `json:"content"`
	Image       string      `json:"image"`
	TagList     []string    `json:"tagList"`
	Description string      `json:"description"`
	IsVisible   int         `json:"isVisible"`
	CreatedAt   *gtime.Time `json:"createdAt"`
	UpdatedAt   *gtime.Time `json:"updatedAt"`
	DeletedAt   *gtime.Time `json:"deletedAt"`
}

// Admin Module

type AdminGetArticleListInput struct{}

type AdminGetArticleListOutput struct {
	ArticleList []*ArticleItem `json:"articleList"`
}

type AdminAddArticleInput struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	IsVisible   int    `json:"isVisible"`
	AuthorId    int    `json:"authorId"`
	TagList     []int  `json:"tagList"`
}

type AdminAddArticleOutput struct {
}

type AdminDeleteArticleInput struct {
	Id int `json:"id"`
}

type AdminDeleteArticleOutput struct {
}

// Blog Module

type BlogGetArticleListInput struct {
}

type BlogGetArticleListOutput struct {
	ArticleList []*ArticleItem `json:"articleList"`
}
