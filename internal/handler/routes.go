// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"ATM/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/atm/ping",
				Handler: PingHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/atm/user/register",
				Handler: UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/atm/user/login",
				Handler: UserLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/atm/user/info",
				Handler: UserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}