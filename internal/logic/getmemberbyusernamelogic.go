package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByUsernameLogic {
	return &GetMemberByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberByUsernameLogic) GetMemberByUsername(in *member.UsernameReq) (*member.MemberInfo, error) {
	q := l.svcCtx.Query
	m, err := q.MmsMember.WithContext(l.ctx).Where(q.MmsMember.Username.Eq(in.Username)).First()
	if err != nil {
		return nil, dbError(err)
	}

	// password is returned intentionally: login needs the hash for verification.
	return memberInfo(m, true, rankCodeByID(l.ctx, q, m.RankID)), nil
}
