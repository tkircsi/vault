package server

import (
	"context"
	"time"

	"github.com/tkircsi/vault/api/grpc/vaultpb"
	"github.com/tkircsi/vault/models"
)

type GRPCHandler struct {
	vault models.VaultHandler
}

func NewGRPCHandler(v models.VaultHandler) *GRPCHandler {
	return &GRPCHandler{
		vault: v,
	}
}

func (h *GRPCHandler) Get(ctx context.Context, r *vaultpb.GetRequest) (*vaultpb.GetResponse, error) {
	t := r.GetToken()
	token, err := h.vault.Get(t)
	if err != nil {
		return nil, err
	}
	return &vaultpb.GetResponse{
		Token: token.Token,
		Value: token.Value,
	}, nil
}

func (h *GRPCHandler) Put(ctx context.Context, r *vaultpb.PutRequest) (*vaultpb.PutResponse, error) {
	v := r.GetValue()
	e := r.GetExpire()
	token, err := h.vault.Put(v, time.Duration(e)*time.Second)
	if err != nil {
		return nil, err
	}
	return &vaultpb.PutResponse{
		Token: token.Token,
	}, nil
}
