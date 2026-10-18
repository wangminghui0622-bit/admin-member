package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetTokenListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenListLogic) GetTokenList(in *member.TokenListReq) (*member.TokenListResp, error) {
	q := l.svcCtx.Query

	// member filters are resolved first, tokens are then queried by the member UUID.
	memberConds := make([]gen.Condition, 0, 4)
	if in.Uuid != nil && *in.Uuid != "" {
		memberConds = append(memberConds, q.MmsMember.ID.Eq(*in.Uuid))
	}
	if in.Username != nil && *in.Username != "" {
		memberConds = append(memberConds, q.MmsMember.Username.Eq(*in.Username))
	}
	if in.Email != nil && *in.Email != "" {
		memberConds = append(memberConds, q.MmsMember.Email.Eq(*in.Email))
	}
	if in.Nickname != nil && *in.Nickname != "" {
		memberConds = append(memberConds, q.MmsMember.Nickname.Eq(*in.Nickname))
	}

	do := q.MmsToken.WithContext(l.ctx)
	if len(memberConds) > 0 {
		m, err := q.MmsMember.WithContext(l.ctx).Where(memberConds...).First()
		if err != nil {
			return nil, dbError(err)
		}
		do = do.Where(q.MmsToken.UUID.Eq(m.ID))
	}

	total, err := do.Count()
	if err != nil {
		return nil, dbError(err)
	}

	page, pageSize := pageArgs(in.Page, in.PageSize)
	tokens, err := do.Order(q.MmsToken.ID.Desc()).Offset((page - 1) * pageSize).Limit(pageSize).Find()
	if err != nil {
		return nil, dbError(err)
	}

	data := make([]*member.TokenInfo, 0, len(tokens))
	for _, t := range tokens {
		data = append(data, tokenInfo(t))
	}

	return &member.TokenListResp{Total: uint64(total), Data: data}, nil
}
