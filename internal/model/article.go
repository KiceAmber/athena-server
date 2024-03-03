package model

import "github.com/gogf/gf/v2/os/gtime"

type ArticleItem struct {
	Id           int         `json:"id"`
	Title        string      `json:"title"`
	AuthorName   string      `json:"authorName"`
	Content      string      `json:"content"`
	Cover        string      `json:"cover"`
	CategoryName string      `json:"categoryName"`
	TagList      []string    `json:"tagList"`
	Description  string      `json:"description"`
	IsVisible    int         `json:"isVisible"`
	CreatedAt    *gtime.Time `json:"createdAt"`
	UpdatedAt    *gtime.Time `json:"updatedAt"`
	DeletedAt    *gtime.Time `json:"deletedAt"`
}

// Admin

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
	CategoryId  int    `json:"categoryId"`
	// Cover       string `json:"cover"`
}

type AdminAddArticleOutput struct{}

type AdminDeleteArticleInput struct {
	Id int `json:"id"`
}

type AdminDeleteArticleOutput struct{}

// Blog

type BlogGetArticleListInput struct{}

type BlogGetArticleListOutput struct {
	ArticleList []*ArticleItem `json:"articleList"`
}

type BlogGetArticleDetailInput struct {
	Id int `json:"id"`
}

type BlogGetArticleDetailOutput struct {
	Content string `json:"content"`
}
