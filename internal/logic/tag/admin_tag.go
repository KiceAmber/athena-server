package tag

import (
	"athena-server/internal/consts"
	"athena-server/internal/dao"
	"athena-server/internal/model"
	"athena-server/internal/model/do"
	"athena-server/internal/model/entity"
	"athena-server/internal/service"
	"context"
)

type sTag struct{}

func init() {
	service.RegisterTag(New())
}

func New() *sTag {
	return &sTag{}
}

func (s sTag) GetTagList(ctx context.Context, in *model.GetTagListInput) (out *model.GetTagListOutput, err error) {
	tagList := []*entity.Tag{}
	if err = dao.Tag.Ctx(ctx).Scan(&tagList); err != nil {
		return nil, err
	}

	out = &model.GetTagListOutput{
		TagList: []*model.TagItem{},
	}
	for _, tagEntity := range tagList {
		out.TagList = append(out.TagList, &model.TagItem{
			Id:        tagEntity.Id,
			Name:      tagEntity.Name,
			CreatedAt: tagEntity.CreatedAt,
		})
	}
	return
}

func (s sTag) AddTag(ctx context.Context, in *model.AddTagInput) (out *model.AddTagOutput, err error) {
	var tag *entity.Tag
	if err = dao.Tag.Ctx(ctx).Where(do.Tag{Name: in.Name}).Scan(&tag); err != nil {
		return nil, err
	}

	if tag != nil {
		return nil, consts.ErrTagExist
	}

	_, err = dao.Tag.Ctx(ctx).Data(do.Tag{
		Name: in.Name,
	}).Insert()
	if err != nil {
		return nil, err
	}

	out = &model.AddTagOutput{}
	return
}

func (s sTag) UpdateTag(ctx context.Context, in *model.UpdateTagInput) (out *model.UpdateTagOutput, err error) {
	_, err = dao.Tag.Ctx(ctx).Where(do.Tag{Id: in.Id}).Data(do.Tag{
		Name: in.Name,
	}).Update()
	if err != nil {
		return nil, err
	}

	out = &model.UpdateTagOutput{}
	return
}

func (s sTag) DeleteTag(ctx context.Context, in *model.DeleteTagInput) (out *model.DeleteTagOutput, err error) {
	// Unscoped 表示不使用软删除来删除数据库记录
	_, err = dao.Tag.Ctx(ctx).Where("id = ?", in.Id).Unscoped().Delete()

	return
}
