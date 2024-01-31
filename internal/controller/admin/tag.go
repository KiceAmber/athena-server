package admin

import (
	"athena-server/api/admin"
	"athena-server/internal/model"
	"athena-server/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) GetTagList(ctx context.Context, req *admin.GetTagListReq) (res *admin.GetTagListRes, err error) {
	out, err := service.Tag().GetTagList(ctx, &model.GetTagListInput{})
	if err != nil {
		return nil, err
	}
	res = &admin.GetTagListRes{
		TagList: out.TagList,
	}
	return
}

func (t *Tag) CreateTag(ctx context.Context, req *admin.CreateTagReq) (res *admin.CreateTagRes, err error) {
	g.Log("tag").Info(ctx, "req.TagName =>", req.Name)
	_, err = service.Tag().CreateTag(ctx, &model.CreateTagInput{
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	res = &admin.CreateTagRes{}
	return
}

// func (t *Tag) UpdateTag(ctx context.Context, req *tag.UpdateTagReq) (res *tag.UpdateTagRes, err error) {

// 	return
// }
