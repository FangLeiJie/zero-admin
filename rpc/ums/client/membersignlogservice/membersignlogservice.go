// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: ums.proto

package membersignlogservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddMemberAddressReq                        = umsclient.AddMemberAddressReq
	AddMemberAddressResp                       = umsclient.AddMemberAddressResp
	AddMemberBrandAttentionReq                 = umsclient.AddMemberBrandAttentionReq
	AddMemberBrandAttentionResp                = umsclient.AddMemberBrandAttentionResp
	AddMemberConsumeSettingReq                 = umsclient.AddMemberConsumeSettingReq
	AddMemberConsumeSettingResp                = umsclient.AddMemberConsumeSettingResp
	AddMemberGrowthLogReq                      = umsclient.AddMemberGrowthLogReq
	AddMemberGrowthLogResp                     = umsclient.AddMemberGrowthLogResp
	AddMemberLevelReq                          = umsclient.AddMemberLevelReq
	AddMemberLevelResp                         = umsclient.AddMemberLevelResp
	AddMemberPointsLogReq                      = umsclient.AddMemberPointsLogReq
	AddMemberPointsLogResp                     = umsclient.AddMemberPointsLogResp
	AddMemberProductCategoryRelationReq        = umsclient.AddMemberProductCategoryRelationReq
	AddMemberProductCategoryRelationResp       = umsclient.AddMemberProductCategoryRelationResp
	AddMemberProductCollectionReq              = umsclient.AddMemberProductCollectionReq
	AddMemberProductCollectionResp             = umsclient.AddMemberProductCollectionResp
	AddMemberReadHistoryReq                    = umsclient.AddMemberReadHistoryReq
	AddMemberReadHistoryResp                   = umsclient.AddMemberReadHistoryResp
	AddMemberRuleSettingReq                    = umsclient.AddMemberRuleSettingReq
	AddMemberRuleSettingResp                   = umsclient.AddMemberRuleSettingResp
	AddMemberSignLogReq                        = umsclient.AddMemberSignLogReq
	AddMemberSignLogResp                       = umsclient.AddMemberSignLogResp
	AddMemberStatisticsInfoReq                 = umsclient.AddMemberStatisticsInfoReq
	AddMemberStatisticsInfoResp                = umsclient.AddMemberStatisticsInfoResp
	AddMemberTagRelationReq                    = umsclient.AddMemberTagRelationReq
	AddMemberTagRelationResp                   = umsclient.AddMemberTagRelationResp
	AddMemberTagReq                            = umsclient.AddMemberTagReq
	AddMemberTagResp                           = umsclient.AddMemberTagResp
	AddMemberTaskRelationReq                   = umsclient.AddMemberTaskRelationReq
	AddMemberTaskRelationResp                  = umsclient.AddMemberTaskRelationResp
	AddMemberTaskReq                           = umsclient.AddMemberTaskReq
	AddMemberTaskResp                          = umsclient.AddMemberTaskResp
	DeleteMemberAddressReq                     = umsclient.DeleteMemberAddressReq
	DeleteMemberAddressResp                    = umsclient.DeleteMemberAddressResp
	DeleteMemberBrandAttentionReq              = umsclient.DeleteMemberBrandAttentionReq
	DeleteMemberBrandAttentionResp             = umsclient.DeleteMemberBrandAttentionResp
	DeleteMemberConsumeSettingReq              = umsclient.DeleteMemberConsumeSettingReq
	DeleteMemberConsumeSettingResp             = umsclient.DeleteMemberConsumeSettingResp
	DeleteMemberInfoReq                        = umsclient.DeleteMemberInfoReq
	DeleteMemberInfoResp                       = umsclient.DeleteMemberInfoResp
	DeleteMemberLevelReq                       = umsclient.DeleteMemberLevelReq
	DeleteMemberLevelResp                      = umsclient.DeleteMemberLevelResp
	DeleteMemberLoginLogReq                    = umsclient.DeleteMemberLoginLogReq
	DeleteMemberLoginLogResp                   = umsclient.DeleteMemberLoginLogResp
	DeleteMemberProductCollectionReq           = umsclient.DeleteMemberProductCollectionReq
	DeleteMemberProductCollectionResp          = umsclient.DeleteMemberProductCollectionResp
	DeleteMemberReadHistoryReq                 = umsclient.DeleteMemberReadHistoryReq
	DeleteMemberReadHistoryResp                = umsclient.DeleteMemberReadHistoryResp
	DeleteMemberRuleSettingReq                 = umsclient.DeleteMemberRuleSettingReq
	DeleteMemberRuleSettingResp                = umsclient.DeleteMemberRuleSettingResp
	DeleteMemberTagReq                         = umsclient.DeleteMemberTagReq
	DeleteMemberTagResp                        = umsclient.DeleteMemberTagResp
	DeleteMemberTaskReq                        = umsclient.DeleteMemberTaskReq
	DeleteMemberTaskResp                       = umsclient.DeleteMemberTaskResp
	LoginReq                                   = umsclient.LoginReq
	LoginResp                                  = umsclient.LoginResp
	MemberAddressListData                      = umsclient.MemberAddressListData
	MemberBrandAttentionListData               = umsclient.MemberBrandAttentionListData
	MemberConsumeSettingListData               = umsclient.MemberConsumeSettingListData
	MemberGrowthLogListData                    = umsclient.MemberGrowthLogListData
	MemberInfoListData                         = umsclient.MemberInfoListData
	MemberLevelListData                        = umsclient.MemberLevelListData
	MemberLoginLogListData                     = umsclient.MemberLoginLogListData
	MemberPointsLogListData                    = umsclient.MemberPointsLogListData
	MemberProductCategoryRelationListData      = umsclient.MemberProductCategoryRelationListData
	MemberProductCollectionListData            = umsclient.MemberProductCollectionListData
	MemberReadHistoryListData                  = umsclient.MemberReadHistoryListData
	MemberRuleSettingListData                  = umsclient.MemberRuleSettingListData
	MemberSignLogListData                      = umsclient.MemberSignLogListData
	MemberTagListData                          = umsclient.MemberTagListData
	MemberTagRelationListData                  = umsclient.MemberTagRelationListData
	MemberTaskListData                         = umsclient.MemberTaskListData
	MemberTaskRelationListData                 = umsclient.MemberTaskRelationListData
	QueryMemberAddressDetailReq                = umsclient.QueryMemberAddressDetailReq
	QueryMemberAddressDetailResp               = umsclient.QueryMemberAddressDetailResp
	QueryMemberAddressListReq                  = umsclient.QueryMemberAddressListReq
	QueryMemberAddressListResp                 = umsclient.QueryMemberAddressListResp
	QueryMemberBrandAttentionDetailReq         = umsclient.QueryMemberBrandAttentionDetailReq
	QueryMemberBrandAttentionDetailResp        = umsclient.QueryMemberBrandAttentionDetailResp
	QueryMemberBrandAttentionListReq           = umsclient.QueryMemberBrandAttentionListReq
	QueryMemberBrandAttentionListResp          = umsclient.QueryMemberBrandAttentionListResp
	QueryMemberConsumeSettingDetailReq         = umsclient.QueryMemberConsumeSettingDetailReq
	QueryMemberConsumeSettingDetailResp        = umsclient.QueryMemberConsumeSettingDetailResp
	QueryMemberConsumeSettingListReq           = umsclient.QueryMemberConsumeSettingListReq
	QueryMemberConsumeSettingListResp          = umsclient.QueryMemberConsumeSettingListResp
	QueryMemberGrowthLogDetailReq              = umsclient.QueryMemberGrowthLogDetailReq
	QueryMemberGrowthLogDetailResp             = umsclient.QueryMemberGrowthLogDetailResp
	QueryMemberGrowthLogListReq                = umsclient.QueryMemberGrowthLogListReq
	QueryMemberGrowthLogListResp               = umsclient.QueryMemberGrowthLogListResp
	QueryMemberInfoDetailReq                   = umsclient.QueryMemberInfoDetailReq
	QueryMemberInfoDetailResp                  = umsclient.QueryMemberInfoDetailResp
	QueryMemberInfoListReq                     = umsclient.QueryMemberInfoListReq
	QueryMemberInfoListResp                    = umsclient.QueryMemberInfoListResp
	QueryMemberLevelDetailReq                  = umsclient.QueryMemberLevelDetailReq
	QueryMemberLevelDetailResp                 = umsclient.QueryMemberLevelDetailResp
	QueryMemberLevelListReq                    = umsclient.QueryMemberLevelListReq
	QueryMemberLevelListResp                   = umsclient.QueryMemberLevelListResp
	QueryMemberLoginLogListReq                 = umsclient.QueryMemberLoginLogListReq
	QueryMemberLoginLogListResp                = umsclient.QueryMemberLoginLogListResp
	QueryMemberPointsLogDetailReq              = umsclient.QueryMemberPointsLogDetailReq
	QueryMemberPointsLogDetailResp             = umsclient.QueryMemberPointsLogDetailResp
	QueryMemberPointsLogListReq                = umsclient.QueryMemberPointsLogListReq
	QueryMemberPointsLogListResp               = umsclient.QueryMemberPointsLogListResp
	QueryMemberProductCategoryRelationListReq  = umsclient.QueryMemberProductCategoryRelationListReq
	QueryMemberProductCategoryRelationListResp = umsclient.QueryMemberProductCategoryRelationListResp
	QueryMemberProductCollectionDetailReq      = umsclient.QueryMemberProductCollectionDetailReq
	QueryMemberProductCollectionDetailResp     = umsclient.QueryMemberProductCollectionDetailResp
	QueryMemberProductCollectionListReq        = umsclient.QueryMemberProductCollectionListReq
	QueryMemberProductCollectionListResp       = umsclient.QueryMemberProductCollectionListResp
	QueryMemberReadHistoryDetailReq            = umsclient.QueryMemberReadHistoryDetailReq
	QueryMemberReadHistoryDetailResp           = umsclient.QueryMemberReadHistoryDetailResp
	QueryMemberReadHistoryListReq              = umsclient.QueryMemberReadHistoryListReq
	QueryMemberReadHistoryListResp             = umsclient.QueryMemberReadHistoryListResp
	QueryMemberRuleSettingDetailReq            = umsclient.QueryMemberRuleSettingDetailReq
	QueryMemberRuleSettingDetailResp           = umsclient.QueryMemberRuleSettingDetailResp
	QueryMemberRuleSettingListReq              = umsclient.QueryMemberRuleSettingListReq
	QueryMemberRuleSettingListResp             = umsclient.QueryMemberRuleSettingListResp
	QueryMemberSignLogDetailReq                = umsclient.QueryMemberSignLogDetailReq
	QueryMemberSignLogDetailResp               = umsclient.QueryMemberSignLogDetailResp
	QueryMemberSignLogListReq                  = umsclient.QueryMemberSignLogListReq
	QueryMemberSignLogListResp                 = umsclient.QueryMemberSignLogListResp
	QueryMemberStatisticsInfoDetailReq         = umsclient.QueryMemberStatisticsInfoDetailReq
	QueryMemberStatisticsInfoDetailResp        = umsclient.QueryMemberStatisticsInfoDetailResp
	QueryMemberStatisticsInfoListData          = umsclient.QueryMemberStatisticsInfoListData
	QueryMemberStatisticsInfoListReq           = umsclient.QueryMemberStatisticsInfoListReq
	QueryMemberStatisticsInfoListResp          = umsclient.QueryMemberStatisticsInfoListResp
	QueryMemberTagDetailReq                    = umsclient.QueryMemberTagDetailReq
	QueryMemberTagDetailResp                   = umsclient.QueryMemberTagDetailResp
	QueryMemberTagListReq                      = umsclient.QueryMemberTagListReq
	QueryMemberTagListResp                     = umsclient.QueryMemberTagListResp
	QueryMemberTagRelationDetailReq            = umsclient.QueryMemberTagRelationDetailReq
	QueryMemberTagRelationDetailResp           = umsclient.QueryMemberTagRelationDetailResp
	QueryMemberTagRelationListReq              = umsclient.QueryMemberTagRelationListReq
	QueryMemberTagRelationListResp             = umsclient.QueryMemberTagRelationListResp
	QueryMemberTaskDetailReq                   = umsclient.QueryMemberTaskDetailReq
	QueryMemberTaskDetailResp                  = umsclient.QueryMemberTaskDetailResp
	QueryMemberTaskListReq                     = umsclient.QueryMemberTaskListReq
	QueryMemberTaskListResp                    = umsclient.QueryMemberTaskListResp
	QueryMemberTaskRelationDetailReq           = umsclient.QueryMemberTaskRelationDetailReq
	QueryMemberTaskRelationDetailResp          = umsclient.QueryMemberTaskRelationDetailResp
	QueryMemberTaskRelationListReq             = umsclient.QueryMemberTaskRelationListReq
	QueryMemberTaskRelationListResp            = umsclient.QueryMemberTaskRelationListResp
	RegisterReq                                = umsclient.RegisterReq
	RegisterResp                               = umsclient.RegisterResp
	UpdateCouponStatusReq                      = umsclient.UpdateCouponStatusReq
	UpdateMemberAddressReq                     = umsclient.UpdateMemberAddressReq
	UpdateMemberAddressResp                    = umsclient.UpdateMemberAddressResp
	UpdateMemberAddressStatusReq               = umsclient.UpdateMemberAddressStatusReq
	UpdateMemberAddressStatusResp              = umsclient.UpdateMemberAddressStatusResp
	UpdateMemberConsumeSettingReq              = umsclient.UpdateMemberConsumeSettingReq
	UpdateMemberConsumeSettingResp             = umsclient.UpdateMemberConsumeSettingResp
	UpdateMemberConsumeSettingStatusResp       = umsclient.UpdateMemberConsumeSettingStatusResp
	UpdateMemberInfoReq                        = umsclient.UpdateMemberInfoReq
	UpdateMemberInfoResp                       = umsclient.UpdateMemberInfoResp
	UpdateMemberInfoStatusReq                  = umsclient.UpdateMemberInfoStatusReq
	UpdateMemberInfoStatusResp                 = umsclient.UpdateMemberInfoStatusResp
	UpdateMemberLevelReq                       = umsclient.UpdateMemberLevelReq
	UpdateMemberLevelResp                      = umsclient.UpdateMemberLevelResp
	UpdateMemberLevelStatusReq                 = umsclient.UpdateMemberLevelStatusReq
	UpdateMemberLevelStatusResp                = umsclient.UpdateMemberLevelStatusResp
	UpdateMemberPointsReq                      = umsclient.UpdateMemberPointsReq
	UpdateMemberPointsResp                     = umsclient.UpdateMemberPointsResp
	UpdateMemberRuleSettingReq                 = umsclient.UpdateMemberRuleSettingReq
	UpdateMemberRuleSettingResp                = umsclient.UpdateMemberRuleSettingResp
	UpdateMemberRuleSettingStatusReq           = umsclient.UpdateMemberRuleSettingStatusReq
	UpdateMemberRuleSettingStatusResp          = umsclient.UpdateMemberRuleSettingStatusResp
	UpdateMemberTagReq                         = umsclient.UpdateMemberTagReq
	UpdateMemberTagResp                        = umsclient.UpdateMemberTagResp
	UpdateMemberTagStatusReq                   = umsclient.UpdateMemberTagStatusReq
	UpdateMemberTagStatusResp                  = umsclient.UpdateMemberTagStatusResp
	UpdateMemberTaskRelationStatusResp         = umsclient.UpdateMemberTaskRelationStatusResp
	UpdateMemberTaskReq                        = umsclient.UpdateMemberTaskReq
	UpdateMemberTaskResp                       = umsclient.UpdateMemberTaskResp
	UpdateMemberTaskStatusReq                  = umsclient.UpdateMemberTaskStatusReq
	UpdateMemberTaskStatusResp                 = umsclient.UpdateMemberTaskStatusResp
	UpdateStatusReq                            = umsclient.UpdateStatusReq

	MemberSignLogService interface {
		// 添加会员签到记录
		AddMemberSignLog(ctx context.Context, in *AddMemberSignLogReq, opts ...grpc.CallOption) (*AddMemberSignLogResp, error)
		// 查询会员签到记录详情
		QueryMemberSignLogDetail(ctx context.Context, in *QueryMemberSignLogDetailReq, opts ...grpc.CallOption) (*QueryMemberSignLogDetailResp, error)
		// 查询会员签到记录列表
		QueryMemberSignLogList(ctx context.Context, in *QueryMemberSignLogListReq, opts ...grpc.CallOption) (*QueryMemberSignLogListResp, error)
	}

	defaultMemberSignLogService struct {
		cli zrpc.Client
	}
)

func NewMemberSignLogService(cli zrpc.Client) MemberSignLogService {
	return &defaultMemberSignLogService{
		cli: cli,
	}
}

// 添加会员签到记录
func (m *defaultMemberSignLogService) AddMemberSignLog(ctx context.Context, in *AddMemberSignLogReq, opts ...grpc.CallOption) (*AddMemberSignLogResp, error) {
	client := umsclient.NewMemberSignLogServiceClient(m.cli.Conn())
	return client.AddMemberSignLog(ctx, in, opts...)
}

// 查询会员签到记录详情
func (m *defaultMemberSignLogService) QueryMemberSignLogDetail(ctx context.Context, in *QueryMemberSignLogDetailReq, opts ...grpc.CallOption) (*QueryMemberSignLogDetailResp, error) {
	client := umsclient.NewMemberSignLogServiceClient(m.cli.Conn())
	return client.QueryMemberSignLogDetail(ctx, in, opts...)
}

// 查询会员签到记录列表
func (m *defaultMemberSignLogService) QueryMemberSignLogList(ctx context.Context, in *QueryMemberSignLogListReq, opts ...grpc.CallOption) (*QueryMemberSignLogListResp, error) {
	client := umsclient.NewMemberSignLogServiceClient(m.cli.Conn())
	return client.QueryMemberSignLogList(ctx, in, opts...)
}
