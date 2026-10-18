package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByIdLogic {
	return &GetMemberByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberByIdLogic) GetMemberById(in *member.UUIDReq) (*member.MemberInfo, error) {
	q := l.svcCtx.Query
	m, err := q.MmsMember.WithContext(l.ctx).Where(q.MmsMember.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, dbError(err)
	}

	return memberInfo(m, false, rankCodeByID(l.ctx, q, m.RankID)), nil
}
