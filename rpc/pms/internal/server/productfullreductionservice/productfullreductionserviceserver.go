// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: pms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/logic/productfullreductionservice"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
)

type ProductFullReductionServiceServer struct {
	svcCtx *svc.ServiceContext
	pmsclient.UnimplementedProductFullReductionServiceServer
}

func NewProductFullReductionServiceServer(svcCtx *svc.ServiceContext) *ProductFullReductionServiceServer {
	return &ProductFullReductionServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加产品满减表(只针对同商品)
func (s *ProductFullReductionServiceServer) AddProductFullReduction(ctx context.Context, in *pmsclient.AddProductFullReductionReq) (*pmsclient.AddProductFullReductionResp, error) {
	l := productfullreductionservicelogic.NewAddProductFullReductionLogic(ctx, s.svcCtx)
	return l.AddProductFullReduction(in)
}

// 删除产品满减表(只针对同商品)
func (s *ProductFullReductionServiceServer) DeleteProductFullReduction(ctx context.Context, in *pmsclient.DeleteProductFullReductionReq) (*pmsclient.DeleteProductFullReductionResp, error) {
	l := productfullreductionservicelogic.NewDeleteProductFullReductionLogic(ctx, s.svcCtx)
	return l.DeleteProductFullReduction(in)
}

// 查询产品满减表(只针对同商品)列表
func (s *ProductFullReductionServiceServer) QueryProductFullReductionList(ctx context.Context, in *pmsclient.QueryProductFullReductionListReq) (*pmsclient.QueryProductFullReductionListResp, error) {
	l := productfullreductionservicelogic.NewQueryProductFullReductionListLogic(ctx, s.svcCtx)
	return l.QueryProductFullReductionList(in)
}
