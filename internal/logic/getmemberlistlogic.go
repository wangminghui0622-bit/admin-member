package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberListLogic {
	return &GetMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberListLogic) GetMemberList(in *member.MemberListReq) (*member.MemberListResp, error) {
	q := l.svcCtx.Query
	conds := make([]gen.Condition, 0, 5)
	if in.Username != nil && *in.Username != "" {
		conds = append(conds, q.MmsMember.Username.Like("%"+*in.Username+"%"))
	}
	if in.Nickname != nil && *in.Nickname != "" {
		conds = append(conds, q.MmsMember.Nickname.Like("%"+*in.Nickname+"%"))
	}
	if in.Mobile != nil && *in.Mobile != "" {
		conds = append(conds, q.MmsMember.Mobile.Like("%"+*in.Mobile+"%"))
	}
	if in.Email != nil && *in.Email != "" {
		conds = append(conds, q.MmsMember.Email.Like("%"+*in.Email+"%"))
	}
	if in.RankId != nil && *in.RankId != 0 {
		conds = append(conds, q.MmsMember.RankID.Eq(int64(*in.RankId)))
	}
	if in.WechatId != nil && *in.WechatId != "" {
		conds = append(conds, q.MmsMember.WechatOpenID.Eq(*in.WechatId))
	}

	do := q.MmsMember.WithContext(l.ctx).Where(conds...)
	total, err := do.Count()
	if err != nil {
		return nil, dbError(err)
	}

	page, pageSize := pageArgs(in.Page, in.PageSize)
	members, err := do.Order(q.MmsMember.ID.Desc()).Offset((page - 1) * pageSize).Limit(pageSize).Find()
	if err != nil {
		return nil, dbError(err)
	}

	rankCodes := make(map[int64]string)
	data := make([]*member.MemberInfo, 0, len(members))
	for _, m := range members {
		code := ""
		if m.RankID != nil {
			if cached, ok := rankCodes[*m.RankID]; ok {
				code = cached
			} else {
				code = rankCodeByID(l.ctx, q, m.RankID)
				rankCodes[*m.RankID] = code
			}
		}
		data = append(data, memberInfo(m, false, code))
	}

	return &member.MemberListResp{Total: uint64(total), Data: data}, nil
}
