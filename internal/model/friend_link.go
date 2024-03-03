package model

import "github.com/gogf/gf/v2/os/gtime"

type FriendLinkItem struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Icon      string      `json:"icon"`
	Url       string      `json:"url"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
	DeletedAt *gtime.Time `json:"deletedAt"`
}

// Admin

type AdminGetFriendLinkListInput struct{}

type AdminGetFriendLinkListOutput struct {
	FriendLinkList []*FriendLinkItem `json:"friendLinkList"`
}

type AdminAddFriendLinkInput struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Url  string `json:"url"`
}

type AdminAddFriendLinkOutput struct{}

type AdminDeleteFriendLinkInput struct {
	Id int `json:"id"`
}

type AdminDeleteFriendLinkOutput struct{}

type AdminUpdateFriendLinkInput struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Url  string `json:"url"`
}

type AdminUpdateFriendLinkOutput struct{}

// Blog
