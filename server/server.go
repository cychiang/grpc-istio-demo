package server

import (
	"context"
	"github.com/cychiang/grpc-istio-demo/proto"
)

func (s *Server) Get(ctx context.Context, in *service.GetParams) *service.GetResponse {

	return nil
}

func (s *Server) Put(ctx context.Context, in *service.PutParams) *service.PutResponse {
	return nil
}
