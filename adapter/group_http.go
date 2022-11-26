package adapter

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/go-wonk/si/sicore"
	"github.com/go-wonk/si/sihttp"
	"github.com/w-woong/common"
	"github.com/w-woong/product/dto"
)

type groupHttp struct {
	client      *sihttp.Client
	baseUrl     string
	host        string
	bearerToken string
}

func NewGroupHttp(client *http.Client, baseUrl string, bearerToken string) *groupHttp {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json; charset=utf-8"

	c := sihttp.NewClient(client, sihttp.WithBaseUrl(baseUrl),
		sihttp.WithRequestOpt(sihttp.WithBearerToken(bearerToken)),
		sihttp.WithWriterOpt(sicore.SetJsonEncoder()),
		sihttp.WithReaderOpt(sicore.SetJsonDecoder()),
		sihttp.WithDefaultHeaders(headers))

	a := &groupHttp{
		client:      c,
		baseUrl:     baseUrl,
		bearerToken: bearerToken,
	}
	if u, err := url.Parse(baseUrl); err == nil {
		a.host = u.Host
	}
	return a
}

func (a *groupHttp) ReadGroup(ctx context.Context, id string) (dto.Group, error) {

	group := dto.Group{}
	res := common.HttpBody{
		Document: &group,
	}

	err := a.client.RequestGetDecodeContext(ctx, "/v1/product/group/"+id, nil, nil, &res)
	if err != nil {
		return dto.NilGroup, err
	}

	if res.Status != http.StatusOK {
		return group, errors.New(res.Message)
	}

	return group, nil
}
