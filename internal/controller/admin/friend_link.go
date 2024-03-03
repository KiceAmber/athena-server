package admin

import (
	"athena-server/api/admin/friend_link"
	"athena-server/internal/model"
	"athena-server/internal/service"
	"context"
)

type FriendLink struct{}

func NewFriendLink() *FriendLink {
	return &FriendLink{}
}

func (c FriendLink) GetFriendLinkList(ctx context.Context, req *friendlink.GetFriendLinkListReq) (res *friendlink.GetFriendLinkListRes, err error) {
	res = &friendlink.GetFriendLinkListRes{}

	out, err := service.FriendLink().AdminGetFriendLinkList(ctx, &model.AdminGetFriendLinkListInput{})
	if err != nil {
		return res, err
	}

	res.FriendLinkList = out.FriendLinkList
	return
}

func (fl FriendLink) AddFriendLink(ctx context.Context, req *friendlink.AddFriendLinkReq) (res *friendlink.AddFriendLinkRes, err error) {
	res = &friendlink.AddFriendLinkRes{}

	_, err = service.FriendLink().AdminAddFriendLink(ctx, &model.AdminAddFriendLinkInput{
		Name: req.Name,
		Icon: req.Icon,
		Url:  req.Url,
	})

	return
}

func (c FriendLink) DeleteFriendLink(ctx context.Context, req *friendlink.DeleteFriendLinkReq) (res *friendlink.DeleteFriendLinkRes, err error) {
	res = &friendlink.DeleteFriendLinkRes{}

	_, err = service.FriendLink().AdminDeleteFriendLink(ctx, &model.AdminDeleteFriendLinkInput{
		Id: req.Id,
	})

	return
}

func (c FriendLink) UpdateFriendLink(ctx context.Context, req *friendlink.UpdateFriendLinkReq) (res *friendlink.UpdateFriendLinkRes, err error) {
	res = &friendlink.UpdateFriendLinkRes{}

	_, err = service.FriendLink().AdminUpdateFriendLink(ctx, &model.AdminUpdateFriendLinkInput{
		Id:   req.Id,
		Name: req.Name,
		Url:  req.Url,
		Icon: req.Icon,
	})
	if err != nil {
		return nil, err
	}

	return
}
