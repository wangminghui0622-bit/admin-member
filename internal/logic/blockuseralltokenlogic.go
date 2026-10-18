package logic

import (
	"context"
	"time"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUserAllTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUserAllTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUserAllTokenLogic {
	return &BlockUserAllTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUserAllTokenLogic) BlockUserAllToken(in *member.UUIDReq) (*member.BaseResp, error) {
	q := l.svcCtx.Query

	if _, err := q.MmsToken.WithContext(l.ctx).Where(q.MmsToken.UUID.Eq(in.Id)).
		Updates(map[string]any{"status": 0}); err != nil {
		return nil, dbError(err)
	}

	tokens, err := q.MmsToken.WithContext(l.ctx).
		Where(q.MmsToken.UUID.Eq(in.Id), q.MmsToken.Status.Eq(0), q.MmsToken.ExpiredAt.Gt(time.Now())).
		Find()
	if err != nil {
		return nil, dbError(err)
	}
	for _, t := range tokens {
		setTokenBlacklist(l.ctx, l.svcCtx.Redis, t.Token, t.ExpiredAt)
	}

	return &member.BaseResp{Msg: successMsg}, nil
}
