package logic

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type CreateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMemberLogic {
	return &CreateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: member
func (l *CreateMemberLogic) CreateMember(in *member.MemberInfo) (*member.BaseUUIDResp, error) {
	q := l.svcCtx.Query

	if in.Mobile != nil && *in.Mobile != "" {
		count, err := q.MmsMember.WithContext(l.ctx).Where(q.MmsMember.Mobile.Eq(*in.Mobile)).Count()
		if err != nil {
			return nil, dbError(err)
		}
		if count > 0 {
			return nil, errorx.New(errorx.InvalidArgument, "手机号已存在")
		}
	}

	if in.Email != nil && *in.Email != "" {
		count, err := q.MmsMember.WithContext(l.ctx).Where(q.MmsMember.Email.Eq(*in.Email)).Count()
		if err != nil {
			return nil, dbError(err)
		}
		if count > 0 {
			return nil, errorx.New(errorx.InvalidArgument, "邮箱已存在")
		}
	}

	now := time.Now()
	m := &model.MmsMember{
		ID:        uuid.NewString(),
		CreatedAt: now,
		UpdatedAt: now,
		Status:    ptr(int32(1)),
		Username:  val(in.Username, ""),
		Nickname:  val(in.Nickname, ""),
	}
	if in.Status != nil {
		m.Status = ptr(int32(*in.Status))
	}
	// 等级不提供时保持 NULL（参考实现 SetNotNilRankID 语义）；
	// 注册流由网关显式传 DefaultRankId。
	if in.RankId != nil {
		m.RankID = ptr(int64(*in.RankId))
	}
	if in.Mobile != nil {
		m.Mobile = ptr(*in.Mobile)
	}
	if in.Email != nil {
		m.Email = ptr(*in.Email)
	}
	if in.Avatar != nil {
		m.Avatar = ptr(*in.Avatar)
	}
	if in.WechatId != nil {
		m.WechatOpenID = ptr(*in.WechatId)
	}
	if in.ExpiredAt != nil {
		expiredAt := time.UnixMilli(*in.ExpiredAt)
		m.ExpiredAt = &expiredAt
	}
	if in.Password != nil && *in.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*in.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errorx.New(errorx.ServerError, err.Error())
		}
		m.Password = string(hashed)
	}

	if err := q.MmsMember.WithContext(l.ctx).Create(m); err != nil {
		return nil, dbError(err)
	}

	return &member.BaseUUIDResp{Id: m.ID, Msg: successMsg}, nil
}
