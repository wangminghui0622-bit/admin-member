package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOauthProviderLogic {
	return &UpdateOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOauthProviderLogic) UpdateOauthProvider(in *member.OauthProviderInfo) (*member.BaseResp, error) {
	if in.Id == nil {
		return nil, errorx.New(errorx.InvalidArgument, "id 不能为空")
	}

	q := l.svcCtx.Query
	old, err := q.MmsOauthProvider.WithContext(l.ctx).Where(q.MmsOauthProvider.ID.Eq(int64(*in.Id))).First()
	if err != nil {
		return nil, dbError(err)
	}

	updates := map[string]any{}
	if in.Name != nil {
		updates["name"] = *in.Name
	}
	if in.ClientId != nil {
		updates["client_id"] = *in.ClientId
	}
	if in.ClientSecret != nil {
		updates["client_secret"] = *in.ClientSecret
	}
	if in.RedirectUrl != nil {
		updates["redirect_url"] = *in.RedirectUrl
	}
	if in.Scopes != nil {
		updates["scopes"] = *in.Scopes
	}
	if in.AuthUrl != nil {
		updates["auth_url"] = *in.AuthUrl
	}
	if in.TokenUrl != nil {
		updates["token_url"] = *in.TokenUrl
	}
	if in.AuthStyle != nil {
		updates["auth_style"] = *in.AuthStyle
	}
	if in.InfoUrl != nil {
		updates["info_url"] = *in.InfoUrl
	}

	if len(updates) > 0 {
		if _, err := q.MmsOauthProvider.WithContext(l.ctx).
			Where(q.MmsOauthProvider.ID.Eq(int64(*in.Id))).Updates(updates); err != nil {
			return nil, dbError(err)
		}
	}

	// the cached oauth config must be refreshed after the provider changed.
	invalidateOauthProvider(old.Name)
	if in.Name != nil {
		invalidateOauthProvider(*in.Name)
	}

	return &member.BaseResp{Msg: successMsg}, nil
}
