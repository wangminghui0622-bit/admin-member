package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberRankLogic {
	return &UpdateMemberRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMemberRankLogic) UpdateMemberRank(in *member.MemberRankInfo) (*member.BaseResp, error) {
	if in.Id == nil {
		return nil, errorx.New(errorx.InvalidArgument, "id 不能为空")
	}

	q := l.svcCtx.Query
	if _, err := q.MmsRank.WithContext(l.ctx).Where(q.MmsRank.ID.Eq(int64(*in.Id))).First(); err != nil {
		return nil, dbError(err)
	}

	updates := map[string]any{}
	if in.Name != nil {
		updates["name"] = *in.Name
	}
	if in.Code != nil {
		updates["code"] = *in.Code
	}
	if in.Description != nil {
		updates["description"] = *in.Description
	}
	if in.Remark != nil {
		updates["remark"] = *in.Remark
	}

	if len(updates) > 0 {
		if _, err := q.MmsRank.WithContext(l.ctx).Where(q.MmsRank.ID.Eq(int64(*in.Id))).Updates(updates); err != nil {
			return nil, dbError(err)
		}
	}

	return &member.BaseResp{Msg: successMsg}, nil
}
