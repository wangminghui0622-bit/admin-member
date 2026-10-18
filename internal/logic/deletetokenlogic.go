package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/dal/query"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTokenLogic) DeleteToken(in *member.UUIDsReq) (*member.BaseResp, error) {
	if len(in.Ids) == 0 {
		return &member.BaseResp{Msg: successMsg}, nil
	}

	err := l.svcCtx.DB.WithContext(l.ctx).Transaction(func(tx *gorm.DB) error {
		tq := query.Use(tx)

		tokens, err := tq.MmsToken.WithContext(l.ctx).Where(tq.MmsToken.ID.In(in.Ids...)).Find()
		if err != nil {
			return err
		}
		// blacklist the tokens before removing the rows.
		for _, t := range tokens {
			setTokenBlacklist(l.ctx, l.svcCtx.Redis, t.Token, t.ExpiredAt)
		}

		_, err = tq.MmsToken.WithContext(l.ctx).Where(tq.MmsToken.ID.In(in.Ids...)).Delete()
		return err
	})
	if err != nil {
		return nil, dbError(err)
	}

	return &member.BaseResp{Msg: successMsg}, nil
}
