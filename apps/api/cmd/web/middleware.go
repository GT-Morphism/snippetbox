package main

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func commonHeaders(ctx huma.Context, next func(huma.Context)) {
	ctx.SetHeader("Server", "Go")
	ctx.SetHeader("Content-Security-Policy", "default-src 'self'; style-src 'self';frame-ancestors 'none'")
	ctx.SetHeader("Referrer-Policy", "strict-origin-when-cross-origin")
	ctx.SetHeader("X-Content-Type-Options", "nosniff")
	next(ctx)
}

func (app *application) logRequest(ctx huma.Context, next func(huma.Context)) {
	var (
		ip     = ctx.RemoteAddr()
		proto  = ctx.Version().Proto
		method = ctx.Operation().Method
		uri    = ctx.URL().Path
	)

	app.logger.Info("received request", "ip", ip, "proto", proto, "method", method, "uri", uri)

	next(ctx)
}

func recoverPanic(api huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		defer func() {
			if err := recover(); err != nil {
				ctx.SetHeader("Connection", "close")

				huma.WriteErr(api, ctx, http.StatusInternalServerError, "Brother, something happend on our site", fmt.Errorf("%s", err))
			}
		}()

		next(ctx)
	}
}
