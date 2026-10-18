package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMemberRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberRankLogic {
	return &DeleteMemberRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMemberRankLogic) DeleteMemberRank(in *member.IDsReq) (*member.BaseResp, error) {
	if len(in.Ids) > 0 {
		ids := make([]int64, 0, len(in.Ids))
		for _, id := range in.Ids {
			ids = append(ids, int64(id))
		}

		q := l.svcCtx.Query
		if _, err := q.MmsRank.WithContext(l.ctx).Where(q.MmsRank.ID.In(ids...)).Delete(); err != nil {
			return nil, dbError(err)
		}
	}

	return &member.BaseResp{Msg: successMsg}, nil
}
