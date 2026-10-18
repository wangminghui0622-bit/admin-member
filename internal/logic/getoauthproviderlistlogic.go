package logic

import (
	"context"

	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen"
)

type GetOauthProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderListLogic {
	return &GetOauthProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOauthProviderListLogic) GetOauthProviderList(in *member.OauthProviderListReq) (*member.OauthProviderListResp, error) {
	q := l.svcCtx.Query
	conds := make([]gen.Condition, 0, 1)
	if in.Name != nil && *in.Name != "" {
		conds = append(conds, q.MmsOauthProvider.Name.Like("%"+*in.Name+"%"))
	}

	do := q.MmsOauthProvider.WithContext(l.ctx).Where(conds...)
	total, err := do.Count()
	if err != nil {
		return nil, dbError(err)
	}

	page, pageSize := pageArgs(in.Page, in.PageSize)
	providers, err := do.Order(q.MmsOauthProvider.ID.Desc()).Offset((page - 1) * pageSize).Limit(pageSize).Find()
	if err != nil {
		return nil, dbError(err)
	}

	data := make([]*member.OauthProviderInfo, 0, len(providers))
	for _, p := range providers {
		data = append(data, oauthProviderInfo(p))
	}

	return &member.OauthProviderListResp{Total: uint64(total), Data: data}, nil
}
