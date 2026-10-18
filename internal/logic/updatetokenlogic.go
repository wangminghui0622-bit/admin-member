package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTokenLogic {
	return &UpdateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTokenLogic) UpdateToken(in *member.TokenInfo) (*member.BaseResp, error) {
	if in.Id == nil {
		return nil, errorx.New(errorx.InvalidArgument, "id 不能为空")
	}

	q := l.svcCtx.Query
	old, err := q.MmsToken.WithContext(l.ctx).Where(q.MmsToken.ID.Eq(*in.Id)).First()
	if err != nil {
		return nil, dbError(err)
	}

	updates := map[string]any{}
	if in.Status != nil {
		updates["status"] = *in.Status
	}
	if in.Source != nil {
		updates["source"] = *in.Source
	}

	if len(updates) > 0 {
		if _, err := q.MmsToken.WithContext(l.ctx).Where(q.MmsToken.ID.Eq(*in.Id)).Updates(updates); err != nil {
			return nil, dbError(err)
		}
	}

	// keep the token blacklist in sync with the status change.
	if in.Status != nil {
		switch *in.Status {
		case 2:
			setTokenBlacklist(l.ctx, l.svcCtx.Redis, old.Token, old.ExpiredAt)
		case 1:
			clearTokenBlacklist(l.ctx, l.svcCtx.Redis, old.Token)
		}
	}

	return &member.BaseResp{Msg: successMsg}, nil
}
