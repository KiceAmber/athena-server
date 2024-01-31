package common

type PageData struct {
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
}

type ListRes struct {
	Total interface{} `json:"total"`
}
