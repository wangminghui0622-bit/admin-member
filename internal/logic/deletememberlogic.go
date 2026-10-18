package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberLogic {
	return &DeleteMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMemberLogic) DeleteMember(in *member.UUIDsReq) (*member.BaseResp, error) {
	if len(in.Ids) > 0 {
		q := l.svcCtx.Query
		if _, err := q.MmsMember.WithContext(l.ctx).Where(q.MmsMember.ID.In(in.Ids...)).Delete(); err != nil {
			return nil, dbError(err)
		}
	}

	return &member.BaseResp{Msg: successMsg}, nil
}
