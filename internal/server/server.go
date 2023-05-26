package server

import (
	"context"

	"github.com/escalopa/itis-tables/internal/application"
)

type Server struct {
	ctx context.Context
	uc  *application.UseCase
}

func New(ctx context.Context, uc *application.UseCase) *Server {
	return &Server{
		ctx: ctx,
		uc:  uc,
	}
}

func (s *Server) Run(port string) error {
	return nil
}
