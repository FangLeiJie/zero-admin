package memberinfoservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogic 会员登录
/*
Author: LiuFeiHua
Date: 2025/5/21 15:27
*/
type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 会员登录
func (l *LoginLogic) Login(in *umsclient.LoginReq) (*umsclient.LoginResp, error) {
	q := query.UmsMemberInfo
	// 1.校验参数
	member, err := q.WithContext(l.ctx).Where(q.Mobile.Eq(in.Mobile)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "账号不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("账号不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员异常")
	}

	// 判断密码
	if member.Password != in.Password {
		logc.Errorf(l.ctx, "账号密码不对,请求参数：%+v", in)
		return nil, errors.New("账号密码不对")
	}

	// 2.添加登录日志
	log := &model.UmsMemberLoginLog{
		MemberID:   member.ID,
		CreateTime: time.Now(),
		MemberIP:   in.Ip,
		City:       "todo",
		LoginType:  in.Source,
		Province:   "todo",
	}
	err = query.UmsMemberLoginLog.WithContext(l.ctx).Create(log)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员登录日志失败,参数：%+v,异常:%s", log, err.Error())
		// 为了兼容，这里不返回错误
	}

	// 3.返回数据
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	secret := l.svcCtx.Config.JWT.AccessSecret
	token, err := createJwtToken(secret, member.Nickname, member.Mobile, accessExpire, member.MemberID)

	if err != nil {
		logc.Errorf(l.ctx, "生成token失败,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("生成token失败")
	}

	l.svcCtx.RabbitMQ.PublishSimple("test", token)
	return &umsclient.LoginResp{
		Token: token,
	}, nil
}
