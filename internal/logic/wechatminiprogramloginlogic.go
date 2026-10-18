package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wangminghui0622-bit/admin-common/errorx"
	"github.com/wangminghui0622-bit/admin-member/internal/svc"
	"github.com/wangminghui0622-bit/admin-member/member"

	"github.com/zeromicro/go-zero/core/logx"
)

type WechatMiniProgramLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWechatMiniProgramLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatMiniProgramLoginLogic {
	return &WechatMiniProgramLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WechatMiniProgramLoginLogic) WechatMiniProgramLogin(in *member.OauthLoginReq) (*member.BaseResp, error) {
	cfg, err := loadOauthProvider(l.ctx, l.svcCtx.Query, in.Provider)
	if err != nil {
		return nil, err
	}

	// state carries the js_code of the mini program.
	url := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		cfg.conf.ClientID, cfg.conf.ClientSecret, in.State,
	)
	req, err := http.NewRequestWithContext(l.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errorx.New(errorx.ServerError, err.Error())
	}

	resp, err := oauthHTTPClient.Do(req)
	if err != nil {
		return nil, errorx.New(errorx.ServerError, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errorx.New(errorx.ServerError, err.Error())
	}

	var result struct {
		OpenID     string `json:"openid"`
		SessionKey string `json:"session_key"`
		UnionID    string `json:"unionid"`
		ErrCode    int32  `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errorx.New(errorx.ServerError, err.Error())
	}
	if result.ErrCode != 0 {
		return nil, errorx.New(errorx.InvalidArgument, result.ErrMsg)
	}

	return &member.BaseResp{Msg: result.OpenID}, nil
}
