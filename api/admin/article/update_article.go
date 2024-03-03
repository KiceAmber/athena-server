package article

import "github.com/gogf/gf/v2/frame/g"

type UpdateArticleReq struct {
	g.Meta      `path:"/updateArticle" method:"POST" tags:"article" sm:"更新文章内容"`
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	AuthorId    int    `json:"authorId"`
	TagList     []int  `json:"tagList"`
	CategoryId  int    `json:"categoryId"`
	IsVisible   int    `json:"isVisible"`
	//Cover       string `json:"cover"`
}

type UpdateArticleRes struct {
}
