package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthProviderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderByIdLogic {
	return &GetOauthProviderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOauthProviderByIdLogic) GetOauthProviderById(in *member.IDReq) (*member.OauthProviderInfo, error) {
	q := l.svcCtx.Query
	p, err := q.MmsOauthProvider.WithContext(l.ctx).Where(q.MmsOauthProvider.ID.Eq(int64(in.Id))).First()
	if err != nil {
		return nil, dbError(err)
	}

	return oauthProviderInfo(p), nil
}
