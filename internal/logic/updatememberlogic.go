package logic

import (
	"context"
	"time"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type UpdateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLogic {
	return &UpdateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMemberLogic) UpdateMember(in *member.MemberInfo) (*member.BaseResp, error) {
	if in.Id == nil || *in.Id == "" {
		return nil, errorx.New(errorx.InvalidArgument, "id 不能为空")
	}

	q := l.svcCtx.Query
	if _, err := q.MmsMember.WithContext(l.ctx).Where(q.MmsMember.ID.Eq(*in.Id)).First(); err != nil {
		return nil, dbError(err)
	}

	updates := map[string]any{}
	if in.Status != nil {
		updates["status"] = *in.Status
	}
	if in.Username != nil {
		updates["username"] = *in.Username
	}
	if in.Nickname != nil {
		updates["nickname"] = *in.Nickname
	}
	if in.RankId != nil {
		updates["rank_id"] = *in.RankId
	}
	if in.Mobile != nil {
		updates["mobile"] = *in.Mobile
	}
	if in.Email != nil {
		updates["email"] = *in.Email
	}
	if in.Avatar != nil {
		updates["avatar"] = *in.Avatar
	}
	if in.WechatId != nil {
		updates["wechat_open_id"] = *in.WechatId
	}
	if in.ExpiredAt != nil {
		updates["expired_at"] = time.UnixMilli(*in.ExpiredAt)
	}
	if in.Password != nil && *in.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*in.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errorx.New(errorx.ServerError, err.Error())
		}
		updates["password"] = string(hashed)
	}

	if len(updates) > 0 {
		if _, err := q.MmsMember.WithContext(l.ctx).Where(q.MmsMember.ID.Eq(*in.Id)).Updates(updates); err != nil {
			return nil, dbError(err)
		}
	}

	return &member.BaseResp{Msg: successMsg}, nil
}
