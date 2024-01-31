package tag

import (
	"athena-server/internal/dao"
	"athena-server/internal/model"
	"athena-server/internal/model/do"
	"athena-server/internal/model/entity"
	"athena-server/internal/service"
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
)

type sTag struct{}

func init() {
	service.RegisterTag(New())
}

func New() *sTag {
	return &sTag{}
}

func (s sTag) GetTagList(ctx context.Context, in *model.GetTagListInput) (out *model.GetTagListOutput, err error) {
	tagEntityList := []*entity.Tag{}
	if err = dao.Tag.Ctx(ctx).Scan(&tagEntityList); err != nil {
		return &model.GetTagListOutput{}, err
	}

	out = &model.GetTagListOutput{
		TagList: []*model.TagItem{},
	}
	for _, tagEntity := range tagEntityList {
		out.TagList = append(out.TagList, &model.TagItem{
			Id:   tagEntity.Id,
			Name: tagEntity.Name,
		})
	}
	return
}

func (s sTag) CreateTag(ctx context.Context, in *model.CreateTagInput) (out *model.CreateTagOutput, err error) {
	g.Log("tag").Debug(ctx, "in.Name => ", in.Name)
	var tag *entity.Tag
	if err = dao.Tag.Ctx(ctx).Where(do.Tag{Name: in.Name}).Scan(&tag); err != nil {
		return nil, err
	}

	if tag != nil {
		return nil, errors.New("该标签已存在")
	}

	// execute insert sql
	_, err = dao.Tag.Ctx(ctx).Data(do.Tag{
		Name: in.Name,
	}).Insert()
	if err != nil {
		return nil, errors.New("插入标签失败")
	}

	out = &model.CreateTagOutput{}
	return
}

func (s sTag) UpdateTag(ctx context.Context, tag *do.Tag) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s sTag) DeleteTag(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}
