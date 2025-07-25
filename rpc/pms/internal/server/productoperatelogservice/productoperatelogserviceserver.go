// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: pms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/logic/productoperatelogservice"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

type ProductOperateLogServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedProductOperateLogServiceServer
}

func NewProductOperateLogServiceServer(svcCtx *svc.ServiceContext) *ProductOperateLogServiceServer {
	return &ProductOperateLogServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加商品操作日志
func (s *ProductOperateLogServiceServer) AddProductOperateLog(ctx context.Context, in *pmsclient.AddProductOperateLogReq) (*pmsclient.AddProductOperateLogResp, error) {
	l := productoperatelogservicelogic.NewAddProductOperateLogLogic(ctx, s.svcCtx)
	return l.AddProductOperateLog(in)
}

// 删除商品操作日志
func (s *ProductOperateLogServiceServer) DeleteProductOperateLog(ctx context.Context, in *pmsclient.DeleteProductOperateLogReq) (*pmsclient.DeleteProductOperateLogResp, error) {
	l := productoperatelogservicelogic.NewDeleteProductOperateLogLogic(ctx, s.svcCtx)
	return l.DeleteProductOperateLog(in)
}

// 查询商品操作日志详情
func (s *ProductOperateLogServiceServer) QueryProductOperateLogDetail(ctx context.Context, in *pmsclient.QueryProductOperateLogDetailReq) (*pmsclient.QueryProductOperateLogDetailResp, error) {
	l := productoperatelogservicelogic.NewQueryProductOperateLogDetailLogic(ctx, s.svcCtx)
	return l.QueryProductOperateLogDetail(in)
}

// 查询商品操作日志列表
func (s *ProductOperateLogServiceServer) QueryProductOperateLogList(ctx context.Context, in *pmsclient.QueryProductOperateLogListReq) (*pmsclient.QueryProductOperateLogListResp, error) {
	l := productoperatelogservicelogic.NewQueryProductOperateLogListLogic(ctx, s.svcCtx)
	return l.QueryProductOperateLogList(in)
}
