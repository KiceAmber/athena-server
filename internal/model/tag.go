package model

type TagItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetTagListInput struct {
}

type GetTagListOutput struct {
	TagList []*TagItem `json:"tag_list"`
}

type CreateTagInput struct {
	Name string `json:"name"`
}

type CreateTagOutput struct {
}
