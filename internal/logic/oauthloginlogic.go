package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OauthLoginLogic) OauthLogin(in *member.OauthLoginReq) (*member.OauthRedirectResp, error) {
	cfg, err := loadOauthProvider(l.ctx, l.svcCtx.Query, in.Provider)
	if err != nil {
		return nil, err
	}

	return &member.OauthRedirectResp{Url: cfg.conf.AuthCodeURL(in.State)}, nil
}
