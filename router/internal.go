package router

import (
	"context"

	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/health"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/pagination"
)

func Initialize(ctx context.Context, s *http.Server) (err error) {

	//Middleware for adding config to ctx
	s.Engine.Use(config.Middleware())
	s.Engine.Use(pagination.Middleware())

	s.Engine.Use(log.RequestLogMiddleware(log.MiddlewareOptions{
		Format:      config.GetString(ctx, "log.format"),
		Level:       config.GetString(ctx, "log.level"),
		LogRequest:  config.GetBool(ctx, "log.request"),
		LogResponse: config.GetBool(ctx, "log.response"),
	}))
	s.GET("/health", health.HealthcheckHandler())
	return nil
}
