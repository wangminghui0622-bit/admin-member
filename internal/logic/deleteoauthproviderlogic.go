package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOauthProviderLogic {
	return &DeleteOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOauthProviderLogic) DeleteOauthProvider(in *member.IDsReq) (*member.BaseResp, error) {
	if len(in.Ids) > 0 {
		ids := make([]int64, 0, len(in.Ids))
		for _, id := range in.Ids {
			ids = append(ids, int64(id))
		}

		q := l.svcCtx.Query
		providers, err := q.MmsOauthProvider.WithContext(l.ctx).Where(q.MmsOauthProvider.ID.In(ids...)).Find()
		if err != nil {
			return nil, dbError(err)
		}
		for _, p := range providers {
			invalidateOauthProvider(p.Name)
		}

		if _, err := q.MmsOauthProvider.WithContext(l.ctx).Where(q.MmsOauthProvider.ID.In(ids...)).Delete(); err != nil {
			return nil, dbError(err)
		}
	}

	return &member.BaseResp{Msg: successMsg}, nil
}
