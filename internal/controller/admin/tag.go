package admin

import (
	"athena-server/api/admin/tag"
	"athena-server/internal/model"
	"athena-server/internal/service"
	"context"
)

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) GetTagList(ctx context.Context, req *tag.GetTagListReq) (res *tag.GetTagListRes, err error) {
	out, err := service.Tag().AdminGetTagList(ctx, &model.AdminGetTagListInput{})
	if err != nil {
		return nil, err
	}
	res = &tag.GetTagListRes{
		TagList: out.TagList,
	}
	return
}

func (t *Tag) AddTag(ctx context.Context, req *tag.AddTagReq) (res *tag.AddTagRes, err error) {
	_, err = service.Tag().AdminAddTag(ctx, &model.AdminAddTagInput{
		Name:      req.Name,
		IsVisible: req.IsVisible,
	})

	if err != nil {
		return nil, err
	}

	res = &tag.AddTagRes{}
	return
}

func (t *Tag) DeleteTag(ctx context.Context, req *tag.DeleteTagReq) (res *tag.DeleteTagRes, err error) {

	_, err = service.Tag().AdminDeleteTag(ctx, &model.AdminDeleteTagInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	res = &tag.DeleteTagRes{}
	return
}

func (t *Tag) UpdateTag(ctx context.Context, req *tag.UpdateTagReq) (res *tag.UpdateTagRes, err error) {
	_, err = service.Tag().AdminUpdateTag(ctx, &model.AdminUpdateTagInput{
		Id:        req.Id,
		Name:      req.Name,
		IsVisible: req.IsVisible,
	})
	if err != nil {
		return nil, err
	}

	res = &tag.UpdateTagRes{}
	return
}
