package request

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/chaitin/libveinmind/go/plugin/log"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote/transport"
)

type Option func(*options) error

type options struct {
	context context.Context
	t       http.RoundTripper
	auth    authn.Authenticator
}

func WithAuth(auth authn.Authenticator) Option {
	return func(o *options) error {
		o.auth = auth
		return nil
	}
}

type catalog struct {
	Repos []string `json:"repositories"`
}

// CatalogPageV1 adapt v1 auth
func CatalogPageV1(target name.Registry, last string, n int, ops ...Option) ([]string, error) {
	o := &options{
		context: context.Background(),
	}

	for _, option := range ops {
		err := option(o)
		if err != nil {
			log.Error(err)
		}
	}

	if o.auth == nil {
		o.auth = authn.Anonymous
	}

	o.t = &basicTransport{
		inner: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
		target: target.RegistryStr(),
		auth:   o.auth,
	}

	query := fmt.Sprintf("last=%s&n=%d", url.QueryEscape(last), n)

	uri := url.URL{
		Scheme:   target.Scheme(),
		Host:     target.RegistryStr(),
		Path:     "/v2/_catalog",
		RawQuery: query,
	}

	client := http.Client{Transport: o.t}
	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req.WithContext(o.context))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := transport.CheckError(resp, http.StatusOK); err != nil {
		return nil, err
	}

	var parsed catalog
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	return parsed.Repos, nil
}
