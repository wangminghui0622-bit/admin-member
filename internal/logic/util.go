package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/dal/query"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-member/member"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

const (
	successMsg = "success"

	blacklistTokenPrefix = "BLACKLIST:TOKEN:"
)

func ptr[T any](v T) *T {
	return &v
}

func val[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}

// dbError maps gorm errors to the common error codes.
func dbError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.New(errorx.InvalidArgument, "目标不存在")
	}
	msg := strings.ToLower(err.Error())
	if strings.Contains(msg, "duplicate") || strings.Contains(msg, "error 1062") || strings.Contains(msg, "constraint") {
		return errorx.New(errorx.InvalidArgument, err.Error())
	}
	return errorx.New(errorx.ServerError, err.Error())
}

func pageArgs(page, pageSize uint64) (int, int) {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	return int(page), int(pageSize)
}

func memberInfo(m *model.MmsMember, withPassword bool, rankCode string) *member.MemberInfo {
	info := &member.MemberInfo{
		Id:        ptr(m.ID),
		CreatedAt: ptr(m.CreatedAt.UnixMilli()),
		UpdatedAt: ptr(m.UpdatedAt.UnixMilli()),
		Status:    ptr(uint32(val(m.Status, 1))),
		Username:  ptr(m.Username),
		Nickname:  ptr(m.Nickname),
		Avatar:    m.Avatar,
		Mobile:    m.Mobile,
		Email:     m.Email,
		WechatId:  m.WechatOpenID,
	}
	if m.RankID != nil {
		info.RankId = ptr(uint64(*m.RankID))
	}
	if m.ExpiredAt != nil {
		info.ExpiredAt = ptr(m.ExpiredAt.UnixMilli())
	}
	if withPassword {
		info.Password = ptr(m.Password)
	}
	if rankCode != "" {
		info.RankCode = ptr(rankCode)
	}
	return info
}

func rankInfo(r *model.MmsRank) *member.MemberRankInfo {
	return &member.MemberRankInfo{
		Id:          ptr(uint64(r.ID)),
		CreatedAt:   ptr(r.CreatedAt.UnixMilli()),
		UpdatedAt:   ptr(r.UpdatedAt.UnixMilli()),
		Name:        ptr(r.Name),
		Code:        ptr(r.Code),
		Description: ptr(r.Description),
		Remark:      ptr(r.Remark),
	}
}

func oauthProviderInfo(p *model.MmsOauthProvider) *member.OauthProviderInfo {
	return &member.OauthProviderInfo{
		Id:           ptr(uint64(p.ID)),
		CreatedAt:    ptr(p.CreatedAt.UnixMilli()),
		UpdatedAt:    ptr(p.UpdatedAt.UnixMilli()),
		Name:         ptr(p.Name),
		ClientId:     ptr(p.ClientID),
		ClientSecret: ptr(p.ClientSecret),
		RedirectUrl:  ptr(p.RedirectURL),
		Scopes:       ptr(p.Scopes),
		AuthUrl:      ptr(p.AuthURL),
		TokenUrl:     ptr(p.TokenURL),
		AuthStyle:    ptr(uint64(p.AuthStyle)),
		InfoUrl:      ptr(p.InfoURL),
	}
}

func tokenInfo(t *model.MmsToken) *member.TokenInfo {
	return &member.TokenInfo{
		Id:        ptr(t.ID),
		CreatedAt: ptr(t.CreatedAt.UnixMilli()),
		UpdatedAt: ptr(t.UpdatedAt.UnixMilli()),
		Status:    ptr(uint32(val(t.Status, 1))),
		Uuid:      ptr(t.UUID),
		Token:     ptr(t.Token),
		Source:    ptr(t.Source),
		ExpiredAt: ptr(t.ExpiredAt.UnixMilli()),
		Username:  ptr(t.Username),
	}
}

// rankCodeByID resolves the rank code of a member, empty when it does not exist.
func rankCodeByID(ctx context.Context, q *query.Query, rankID *int64) string {
	if rankID == nil {
		return ""
	}
	rank, err := q.MmsRank.WithContext(ctx).Where(q.MmsRank.ID.Eq(*rankID)).First()
	if err != nil {
		return ""
	}
	return rank.Code
}

func setTokenBlacklist(ctx context.Context, rds *redis.Redis, token string, expiredAt time.Time) {
	seconds := int(time.Until(expiredAt).Seconds())
	if seconds < 1 {
		seconds = 1
	}
	if err := rds.SetexCtx(ctx, blacklistTokenPrefix+token, "1", seconds); err != nil {
		logx.Errorf("set token blacklist failed: %v", err)
	}
}

func clearTokenBlacklist(ctx context.Context, rds *redis.Redis, token string) {
	if _, err := rds.DelCtx(ctx, blacklistTokenPrefix+token); err != nil {
		logx.Errorf("clear token blacklist failed: %v", err)
	}
}
