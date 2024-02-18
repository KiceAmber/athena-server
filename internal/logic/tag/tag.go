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

func (s sTag) AdminGetTagList(ctx context.Context, in *model.AdminGetTagListInput) (out *model.AdminGetTagListOutput, err error) {
	out = &model.AdminGetTagListOutput{
		TagList: []*model.TagItem{},
	}

	tagList := []*entity.Tag{}
	if err = dao.Tag.Ctx(ctx).Scan(&tagList); err != nil {
		return nil, err
	}

	for _, tagEntity := range tagList {
		out.TagList = append(out.TagList, &model.TagItem{
			Id:        tagEntity.Id,
			Name:      tagEntity.Name,
			IsVisible: tagEntity.IsVisible,
			CreatedAt: tagEntity.CreatedAt,
		})
	}
	return
}

func (s sTag) AdminAddTag(ctx context.Context, in *model.AdminAddTagInput) (out *model.AdminAddTagOutput, err error) {
	out = &model.AdminAddTagOutput{}

	var tag *entity.Tag
	if err = dao.Tag.Ctx(ctx).Where(do.Tag{
		Name: in.Name,
	}).Scan(&tag); err != nil {
		return nil, err
	}

	if tag != nil {
		return nil, consts.ErrTagExist
	}

	_, err = dao.Tag.Ctx(ctx).Data(do.Tag{
		Name:      in.Name,
		IsVisible: in.IsVisible,
	}).Insert()
	if err != nil {
		return nil, err
	}

	return
}

func (s sTag) AdminUpdateTag(ctx context.Context, in *model.AdminUpdateTagInput) (out *model.AdminUpdateTagOutput, err error) {
	out = &model.AdminUpdateTagOutput{}

	_, err = dao.Tag.Ctx(ctx).Where(do.Tag{Id: in.Id}).Data(do.Tag{
		Name:      in.Name,
		IsVisible: in.IsVisible,
	}).Update()
	if err != nil {
		return nil, err
	}

	return
}

func (s sTag) AdminDeleteTag(ctx context.Context, in *model.AdminDeleteTagInput) (out *model.AdminDeleteTagOutput, err error) {
	// Unscoped 表示不使用软删除来删除数据库记录
	_, err = dao.Tag.Ctx(ctx).Where("id = ?", in.Id).Unscoped().Delete()
	return
}
