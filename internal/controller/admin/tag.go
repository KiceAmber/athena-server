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
	out, err := service.Tag().GetTagList(ctx, &model.GetTagListInput{})
	if err != nil {
		return nil, err
	}
	res = &tag.GetTagListRes{
		TagList: out.TagList,
	}
	return
}

func (t *Tag) AddTag(ctx context.Context, req *tag.AddTagReq) (res *tag.AddTagRes, err error) {
	_, err = service.Tag().AddTag(ctx, &model.AddTagInput{
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	res = &tag.AddTagRes{}
	return
}

func (t *Tag) DeleteTag(ctx context.Context, req *tag.DeleteTagReq) (res *tag.DeleteTagRes, err error) {

	_, err = service.Tag().DeleteTag(ctx, &model.DeleteTagInput{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	res = &tag.DeleteTagRes{}
	return
}

func (t *Tag) UpdateTag(ctx context.Context, req *tag.UpdateTagReq) (res *tag.UpdateTagRes, err error) {
	_, err = service.Tag().UpdateTag(ctx, &model.UpdateTagInput{
		Id:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	res = &tag.UpdateTagRes{}
	return
}
