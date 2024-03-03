// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"athena-server/internal/model"
	"context"
)

type (
	IFriendLink interface {
		AdminAddFriendLink(ctx context.Context, in *model.AdminAddFriendLinkInput) (out *model.AdminAddFriendLinkOutput, err error)
		AdminGetFriendLinkList(ctx context.Context, in *model.AdminGetFriendLinkListInput) (out *model.AdminGetFriendLinkListOutput, err error)
		AdminDeleteFriendLink(ctx context.Context, in *model.AdminDeleteFriendLinkInput) (out *model.AdminDeleteFriendLinkOutput, err error)
		AdminUpdateFriendLink(ctx context.Context, in *model.AdminUpdateFriendLinkInput) (out *model.AdminUpdateFriendLinkOutput, err error)
	}
)

var (
	localFriendLink IFriendLink
)

func FriendLink() IFriendLink {
	if localFriendLink == nil {
		panic("implement not found for interface IFriendLink, forgot register?")
	}
	return localFriendLink
}

func RegisterFriendLink(i IFriendLink) {
	localFriendLink = i
}
