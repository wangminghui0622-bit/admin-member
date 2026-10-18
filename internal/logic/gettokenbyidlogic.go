package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenByIdLogic {
	return &GetTokenByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenByIdLogic) GetTokenById(in *member.UUIDReq) (*member.TokenInfo, error) {
	q := l.svcCtx.Query
	t, err := q.MmsToken.WithContext(l.ctx).Where(q.MmsToken.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, dbError(err)
	}

	return tokenInfo(t), nil
}
