package model

type ArticleItem struct {
	Id         int      `json:"id"`
	Title      string   `json:"title"`
	AuthorName string   `json:"author_name"`
	Content    string   `json:"content"`
	Image      string   `json:"image"`
	TagList    []string `json:"tagList"`
}

type GetArticleListInput struct{}

type GetArticleListOutput struct {
	ArticleList []*ArticleItem `json:"articleList"`
}

type AddArticleInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AddArticleOutput struct {
}
