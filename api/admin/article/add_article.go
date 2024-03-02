package article

import "github.com/gogf/gf/v2/frame/g"

type AddArticleReq struct {
	g.Meta      `path:"/addArticle" method:"POST" tags:"article" sm:"添加文章"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	AuthorId    int    `json:"authorId"`
	TagList     []int  `json:"tagList"`
	CategoryId  int    `json:"categoryId"`
	IsVisible   int    `json:"isVisible"`
	//Cover       string `json:"cover"`
}

type AddArticleRes struct {
}
