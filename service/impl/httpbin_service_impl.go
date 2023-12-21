package impl

import (
	"context"

	"github.com/donnyirianto/go-be-fiber/client"
	"github.com/donnyirianto/go-be-fiber/common"
	"github.com/donnyirianto/go-be-fiber/model"
	"github.com/donnyirianto/go-be-fiber/service"
)

func NewHttpBinServiceImpl(httpBinClient *client.HttpBinClient) service.HttpBinService {
	return &httpBinServiceImpl{HttpBinClient: *httpBinClient}
}

type httpBinServiceImpl struct {
	client.HttpBinClient
}

func (h *httpBinServiceImpl) PostMethod(ctx context.Context) {
	httpBin := model.HttpBin{
		Name: "rizki",
	}
	var response map[string]interface{}
	h.HttpBinClient.PostMethod(ctx, &httpBin, &response)
	common.NewLogger().Info("log response service ", response)
}
