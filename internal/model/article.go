package model

type ArticleItem struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	AuthorName string `json:"author_name"`
	Content    string `json:"content"`
}

type GetArticleListInput struct{}

type GetArticleListOutput struct {
	ArticleList []*ArticleItem `json:"articleList"`
}

type AddArticleInput struct {
}

type AddArticleOutput struct {
}
