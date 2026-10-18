package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberRankByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberRankByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberRankByIdLogic {
	return &GetMemberRankByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberRankByIdLogic) GetMemberRankById(in *member.IDReq) (*member.MemberRankInfo, error) {
	q := l.svcCtx.Query
	r, err := q.MmsRank.WithContext(l.ctx).Where(q.MmsRank.ID.Eq(int64(in.Id))).First()
	if err != nil {
		return nil, dbError(err)
	}

	return rankInfo(r), nil
}
