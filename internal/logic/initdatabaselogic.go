package logic

import (
	"context"

	_ "embed"

	"time"

	"github.com/google/uuid"
	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

//go:embed init.sql
var initSQL string

const initDatabaseKey = "initDatabase:member"

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: base
func (l *InitDatabaseLogic) InitDatabase(in *member.Empty) (*member.BaseResp, error) {
	if ok, err := l.svcCtx.Redis.ExistsCtx(l.ctx, initDatabaseKey); err == nil && ok {
		return &member.BaseResp{Msg: "already initialized"}, nil
	}

	if err := l.svcCtx.DB.WithContext(l.ctx).Exec(initSQL).Error; err != nil {
		l.Errorf("init database failed: %v", err)
		return nil, errorx.New(errorx.ServerError, "初始化数据库失败: "+err.Error())
	}

	if err := l.seedMembers(); err != nil {
		l.Errorf("seed members failed: %v", err)
		return nil, errorx.New(errorx.ServerError, "初始化会员数据失败: "+err.Error())
	}

	if err := l.svcCtx.Redis.SetCtx(l.ctx, initDatabaseKey, "1"); err != nil {
		l.Errorf("set init database key failed: %v", err)
	}

	return &member.BaseResp{Msg: successMsg}, nil
}

// seedMembers creates the built-in member accounts, skipping existing usernames.
func (l *InitDatabaseLogic) seedMembers() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte("simple-admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	expiredAt := time.Date(2099, 1, 1, 0, 0, 0, 0, time.Local)
	seeds := []*model.MmsMember{
		{
			ID:        uuid.NewString(),
			CreatedAt: now,
			UpdatedAt: now,
			Status:    ptr(int32(1)),
			Username:  "normalMember",
			Password:  string(hashed),
			Nickname:  "Normal Member",
			Mobile:    ptr("18888888888"),
			Email:     ptr("simpleadmin@gmail.com"),
			Avatar:    ptr(""),
			ExpiredAt: &expiredAt,
			RankID:    ptr(int64(1)),
		},
		{
			ID:        uuid.NewString(),
			CreatedAt: now,
			UpdatedAt: now,
			Status:    ptr(int32(1)),
			Username:  "VIPMember",
			Password:  string(hashed),
			Nickname:  "VIP Member",
			Mobile:    ptr("18888888889"),
			Email:     ptr("vip@gmail.com"),
			Avatar:    ptr(""),
			ExpiredAt: &expiredAt,
			RankID:    ptr(int64(2)),
		},
	}

	for _, seed := range seeds {
		count, err := l.svcCtx.Query.MmsMember.WithContext(l.ctx).
			Where(l.svcCtx.Query.MmsMember.Username.Eq(seed.Username)).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		if err := l.svcCtx.Query.MmsMember.WithContext(l.ctx).Create(seed); err != nil {
			return err
		}
	}
	return nil
}
