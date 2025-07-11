package seckill_activity

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSeckillActivityStatusLogic 更新秒杀活动状态状态
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type UpdateSeckillActivityStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSeckillActivityStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillActivityStatusLogic {
	return &UpdateSeckillActivityStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateSeckillActivityStatus 更新秒杀活动状态
func (l *UpdateSeckillActivityStatusLogic) UpdateSeckillActivityStatus(req *types.UpdateSeckillActivityStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.SeckillActivityService.UpdateSeckillActivityStatus(l.ctx, &smsclient.UpdateSeckillActivityStatusReq{
		Ids:      req.Ids,    // 编号
		Status:   req.Status, // 状态:0-上线,1-下线
		UpdateBy: userId,     // 更新人ID

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀活动状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新秒杀活动状态成功",
	}, nil
}
