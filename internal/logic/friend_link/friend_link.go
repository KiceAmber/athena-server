package friendlink

import (
	"athena-server/internal/dao"
	"athena-server/internal/model"
	"athena-server/internal/model/do"
	"athena-server/internal/model/entity"
	"athena-server/internal/service"
	"context"
)

type sFriendLink struct{}

func init() {
	service.RegisterFriendLink(New())
}

func New() *sFriendLink {
	return &sFriendLink{}
}

func (s sFriendLink) AdminGetFriendLinkList(ctx context.Context, in *model.AdminGetFriendLinkListInput) (out *model.AdminGetFriendLinkListOutput, err error) {
	out = &model.AdminGetFriendLinkListOutput{
		FriendLinkList: []*model.FriendLinkItem{},
	}
	var friendLinkList = []*entity.FriendLink{}

	if err = dao.FriendLink.Ctx(ctx).Scan(&friendLinkList); err != nil {
		return
	}
	for _, item := range friendLinkList {
		friendLink := &model.FriendLinkItem{
			Id:        item.Id,
			Name:      item.Name,
			Icon:      item.Icon,
			Url:       item.Url,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			DeletedAt: item.DeletedAt,
		}
		out.FriendLinkList = append(out.FriendLinkList, friendLink)
	}

	return
}

func (s sFriendLink) AdminAddFriendLink(ctx context.Context, in *model.AdminAddFriendLinkInput) (out *model.AdminAddFriendLinkOutput, err error) {
	out = &model.AdminAddFriendLinkOutput{}

	_, err = dao.FriendLink.Ctx(ctx).Data(&do.FriendLink{
		Name: in.Name,
		Url:  in.Url,
		Icon: in.Icon,
	}).Insert()

	return
}

func (s sFriendLink) AdminDeleteFriendLink(ctx context.Context, in *model.AdminDeleteFriendLinkInput) (out *model.AdminDeleteFriendLinkOutput, err error) {
	out = &model.AdminDeleteFriendLinkOutput{}

	_, err = dao.FriendLink.Ctx(ctx).Unscoped().Delete("id = ?", in.Id)

	return
}

func (s sFriendLink) AdminUpdateFriendLink(ctx context.Context, in *model.AdminUpdateFriendLinkInput) (out *model.AdminUpdateFriendLinkOutput, err error) {
	out = &model.AdminUpdateFriendLinkOutput{}

	_, err = dao.FriendLink.Ctx(ctx).Where("id = ?", in.Id).Update(&do.FriendLink{
		Name: in.Name,
		Icon: in.Icon,
		Url:  in.Url,
	})

	return
}
