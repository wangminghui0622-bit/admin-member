package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetMemberRankListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberRankListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberRankListLogic {
	return &GetMemberRankListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberRankListLogic) GetMemberRankList(in *member.MemberRankListReq) (*member.MemberRankListResp, error) {
	q := l.svcCtx.Query
	conds := make([]gen.Condition, 0, 3)
	if in.Name != nil && *in.Name != "" {
		conds = append(conds, q.MmsRank.Name.Like("%"+*in.Name+"%"))
	}
	if in.Description != nil && *in.Description != "" {
		conds = append(conds, q.MmsRank.Description.Like("%"+*in.Description+"%"))
	}
	if in.Remark != nil && *in.Remark != "" {
		conds = append(conds, q.MmsRank.Remark.Like("%"+*in.Remark+"%"))
	}

	do := q.MmsRank.WithContext(l.ctx).Where(conds...)
	total, err := do.Count()
	if err != nil {
		return nil, dbError(err)
	}

	page, pageSize := pageArgs(in.Page, in.PageSize)
	ranks, err := do.Order(q.MmsRank.ID.Desc()).Offset((page - 1) * pageSize).Limit(pageSize).Find()
	if err != nil {
		return nil, dbError(err)
	}

	data := make([]*member.MemberRankInfo, 0, len(ranks))
	for _, r := range ranks {
		data = append(data, rankInfo(r))
	}

	return &member.MemberRankListResp{Total: uint64(total), Data: data}, nil
}
