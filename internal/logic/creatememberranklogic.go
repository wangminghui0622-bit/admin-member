package logic

import (
	"context"
	"time"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMemberRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMemberRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMemberRankLogic {
	return &CreateMemberRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: memberrank
func (l *CreateMemberRankLogic) CreateMemberRank(in *member.MemberRankInfo) (*member.BaseIDResp, error) {
	now := time.Now()
	r := &model.MmsRank{
		CreatedAt:   now,
		UpdatedAt:   now,
		Name:        val(in.Name, ""),
		Code:        val(in.Code, ""),
		Description: val(in.Description, ""),
		Remark:      val(in.Remark, ""),
	}

	if err := l.svcCtx.Query.MmsRank.WithContext(l.ctx).Create(r); err != nil {
		return nil, dbError(err)
	}

	return &member.BaseIDResp{Id: uint64(r.ID), Msg: successMsg}, nil
}
