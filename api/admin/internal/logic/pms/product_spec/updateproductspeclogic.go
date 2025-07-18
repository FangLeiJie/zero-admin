package product_spec

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductSpecLogic 更新商品规格
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type UpdateProductSpecLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductSpecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSpecLogic {
	return &UpdateProductSpecLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductSpec 更新商品规格
func (l *UpdateProductSpecLogic) UpdateProductSpec(req *types.UpdateProductSpecReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductSpecService.UpdateProductSpec(l.ctx, &pmsclient.UpdateProductSpecReq{
		Id:         req.Id,         //
		CategoryId: req.CategoryId, // 分类ID
		Name:       req.Name,       // 规格名称
		Sort:       req.Sort,       // 排序
		Status:     req.Status,     // 状态：0->禁用；1->启用
		UpdateBy:   userId,         // 更新人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品规格失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新商品规格成功",
	}, nil
}
