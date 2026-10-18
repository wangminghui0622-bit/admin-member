package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/wangminghui0622-bit/admin-common/dal/model"
	"github.com/wangminghui0622-bit/admin-common/dal/query"
	"github.com/wangminghui0622-bit/admin-common/errorx"
	"golang.org/x/oauth2"
)

type oauthProviderConfig struct {
	conf      oauth2.Config
	infoURL   string
	authStyle int64
}

var (
	oauthProviderMu    sync.RWMutex
	oauthProviderCache = map[string]*oauthProviderConfig{}

	oauthHTTPClient = &http.Client{Timeout: 10 * time.Second}
)

// loadOauthProvider returns the oauth2 config of a provider, cached in memory by name.
func loadOauthProvider(ctx context.Context, q *query.Query, name string) (*oauthProviderConfig, error) {
	oauthProviderMu.RLock()
	cached, ok := oauthProviderCache[name]
	oauthProviderMu.RUnlock()
	if ok {
		return cached, nil
	}

	p, err := q.MmsOauthProvider.WithContext(ctx).Where(q.MmsOauthProvider.Name.Eq(name)).First()
	if err != nil {
		return nil, dbError(err)
	}

	cfg := buildOauthProvider(p)
	oauthProviderMu.Lock()
	oauthProviderCache[name] = cfg
	oauthProviderMu.Unlock()
	return cfg, nil
}

func buildOauthProvider(p *model.MmsOauthProvider) *oauthProviderConfig {
	return &oauthProviderConfig{
		conf: oauth2.Config{
			ClientID:     p.ClientID,
			ClientSecret: p.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:   replaceKeywords(p.AuthURL, p),
				TokenURL:  p.TokenURL,
				AuthStyle: oauth2.AuthStyle(p.AuthStyle),
			},
			RedirectURL: p.RedirectURL,
			Scopes:      strings.Split(p.Scopes, " "),
		},
		infoURL:   replaceKeywords(p.InfoURL, p),
		authStyle: p.AuthStyle,
	}
}

func invalidateOauthProvider(name string) {
	oauthProviderMu.Lock()
	delete(oauthProviderCache, name)
	oauthProviderMu.Unlock()
}

func replaceKeywords(raw string, p *model.MmsOauthProvider) string {
	raw = strings.ReplaceAll(raw, "CLIENT_ID", p.ClientID)
	raw = strings.ReplaceAll(raw, "SECRET", p.ClientSecret)
	raw = strings.ReplaceAll(raw, "REDIRECT_URL", p.RedirectURL)
	raw = strings.ReplaceAll(raw, "SCOPE", p.Scopes)
	return raw
}

type oauthUserInfo struct {
	Email    string `json:"email"`
	NickName string `json:"nickName"`
	Picture  string `json:"picture"`
	Mobile   string `json:"mobile"`
}

// fetchOauthUserInfo exchanges the oauth code and requests the user info endpoint.
func fetchOauthUserInfo(ctx context.Context, cfg *oauthProviderConfig, code string) (*oauthUserInfo, error) {
	token, err := cfg.conf.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	var (
		resp *http.Response
		req  *http.Request
	)
	switch cfg.authStyle {
	case 1:
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, strings.ReplaceAll(cfg.infoURL, "TOKEN", token.AccessToken), nil)
		if err != nil {
			return nil, err
		}
		resp, err = oauthHTTPClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to get user's information: %s", err.Error())
		}
	case 2:
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, cfg.infoURL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Bearer "+token.AccessToken)
		resp, err = oauthHTTPClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed getting user info: %s", err.Error())
		}
	default:
		return nil, errorx.New(errorx.InvalidArgument, "不支持的认证方式")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err.Error())
	}

	var info oauthUserInfo
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user info: %s", err.Error())
	}
	return &info, nil
}
