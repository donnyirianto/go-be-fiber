package client

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/model"
)

type HttpBinClient interface {
	PostMethod(ctx context.Context, requestBody *model.HttpBin, response *map[string]interface{})
}
