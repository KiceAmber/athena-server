package model

type UserItem struct {
	Passport string `json:"passport"`
	Password string `json:"password"`
}

type LoginInput struct {
	Passport string `json:"passport"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}
