package logic

import (
	"context"
	"time"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOauthProviderLogic {
	return &CreateOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: oauthprovider
func (l *CreateOauthProviderLogic) CreateOauthProvider(in *member.OauthProviderInfo) (*member.BaseIDResp, error) {
	now := time.Now()
	p := &model.MmsOauthProvider{
		CreatedAt:    now,
		UpdatedAt:    now,
		Name:         val(in.Name, ""),
		ClientID:     val(in.ClientId, ""),
		ClientSecret: val(in.ClientSecret, ""),
		RedirectURL:  val(in.RedirectUrl, ""),
		Scopes:       val(in.Scopes, ""),
		AuthURL:      val(in.AuthUrl, ""),
		TokenURL:     val(in.TokenUrl, ""),
		AuthStyle:    int64(val(in.AuthStyle, 0)),
		InfoURL:      val(in.InfoUrl, ""),
	}

	if err := l.svcCtx.Query.MmsOauthProvider.WithContext(l.ctx).Create(p); err != nil {
		return nil, dbError(err)
	}
	invalidateOauthProvider(p.Name)

	return &member.BaseIDResp{Id: uint64(p.ID), Msg: successMsg}, nil
}
