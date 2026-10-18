package logic

import (
	"context"
	"errors"
	"strings"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type OauthCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOauthCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthCallbackLogic {
	return &OauthCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OauthCallbackLogic) OauthCallback(in *member.CallbackReq) (*member.MemberInfo, error) {
	// state is formatted as "<anything>-<providerName>".
	stateParts := strings.Split(in.State, "-")
	if len(stateParts) < 2 {
		return nil, errorx.New(errorx.InvalidArgument, "无效的 state")
	}
	provider := stateParts[1]

	cfg, err := loadOauthProvider(l.ctx, l.svcCtx.Query, provider)
	if err != nil {
		return nil, err
	}

	userInfo, err := fetchOauthUserInfo(l.ctx, cfg, in.Code)
	if err != nil {
		return nil, errorx.New(errorx.InvalidArgument, err.Error())
	}
	if userInfo.Email == "" {
		return nil, errorx.New(errorx.InvalidArgument, "邮箱为空")
	}

	// oauth login does not register new members automatically.
	q := l.svcCtx.Query
	m, err := q.MmsMember.WithContext(l.ctx).Where(q.MmsMember.Email.Eq(userInfo.Email)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorx.New(errorx.InvalidArgument, "用户不存在")
		}
		return nil, dbError(err)
	}

	return memberInfo(m, false, rankCodeByID(l.ctx, q, m.RankID)), nil
}
