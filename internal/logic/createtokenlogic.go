package logic

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: token
func (l *CreateTokenLogic) CreateToken(in *member.TokenInfo) (*member.BaseUUIDResp, error) {
	now := time.Now()
	t := &model.MmsToken{
		ID:        uuid.NewString(),
		CreatedAt: now,
		UpdatedAt: now,
		Status:    ptr(int32(val(in.Status, 1))),
		UUID:      val(in.Uuid, ""),
		Token:     val(in.Token, ""),
		Username:  val(in.Username, "unknown"),
		Source:    val(in.Source, ""),
		ExpiredAt: now,
	}
	if in.ExpiredAt != nil {
		t.ExpiredAt = time.UnixMilli(*in.ExpiredAt)
	}

	if err := l.svcCtx.Query.MmsToken.WithContext(l.ctx).Create(t); err != nil {
		return nil, dbError(err)
	}

	return &member.BaseUUIDResp{Id: t.ID, Msg: successMsg}, nil
}
