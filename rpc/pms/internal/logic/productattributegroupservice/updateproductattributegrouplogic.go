package productattributegroupservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateProductAttributeGroupLogic 更新商品属性分组
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type UpdateProductAttributeGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductAttributeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeGroupLogic {
	return &UpdateProductAttributeGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductAttributeGroup 更新商品属性分组
func (l *UpdateProductAttributeGroupLogic) UpdateProductAttributeGroup(in *pmsclient.UpdateProductAttributeGroupReq) (*pmsclient.UpdateProductAttributeGroupResp, error) {
	group := query.PmsProductAttributeGroup
	q := group.WithContext(l.ctx)

	// 1.根据商品属性分组id查询商品属性分组是否已存在
	detail, err := q.Where(group.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品属性分组不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品属性分组不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品属性分组异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品属性分组异常")
	}

	now := time.Now()
	item := &model.PmsProductAttributeGroup{
		ID:         in.Id,             // 主键id
		CategoryID: in.CategoryId,     // 分类ID
		Name:       in.Name,           // 分组名称
		Sort:       in.Sort,           // 排序
		Status:     in.Status,         // 状态：0->禁用；1->启用
		CreateBy:   detail.CreateBy,   // 创建人ID
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   &in.UpdateBy,      // 更新人ID
		UpdateTime: &now,              // 更新时间
	}

	// 2.商品属性分组存在时,则直接更新商品属性分组
	err = l.svcCtx.DB.Model(&model.PmsProductAttributeGroup{}).WithContext(l.ctx).Where(group.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新商品属性分组失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新商品属性分组失败")
	}

	return &pmsclient.UpdateProductAttributeGroupResp{}, nil
}
