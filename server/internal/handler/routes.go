// Code generated by jzero. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest"
	"server/internal/svc"

	auth "server/internal/handler/auth"
	route "server/internal/handler/route"
	systemmenu "server/internal/handler/system/menu"
	systemrole "server/internal/handler/system/role"
	systemuser "server/internal/handler/system/user"
	version "server/internal/handler/version"
)

var (
	_ = http.StatusOK
	_ = time.Now()
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/auth/login",
					Handler: auth.Login(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/auth/refreshToken",
					Handler: auth.RefreshToken(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/auth/register",
					Handler: auth.Register(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/auth/sendVerificationCode",
					Handler: auth.SendVerificationCode(serverCtx),
				},
			},
		)

		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/auth/error",
					Handler: auth.Error(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/auth/getUserInfo",
					Handler: auth.GetUserInfo(serverCtx),
				},
			},
			rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/route/getConstantRoutes",
					Handler: route.GetConstantRoutes(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/route/getUserRoutes",
					Handler: route.GetUserRoutes(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/route/isRouteExist",
					Handler: route.IsRouteExist(serverCtx),
				},
			},
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getAllPages",
					Handler: systemmenu.GetAllPages(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getMenuList/v2",
					Handler: systemmenu.List(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getMenuTree",
					Handler: systemmenu.Tree(serverCtx),
				},
			},
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getAllRoles",
					Handler: systemrole.GetAll(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getRoleList",
					Handler: systemrole.List(serverCtx),
				},
			},
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getUserList",
					Handler: systemuser.List(serverCtx),
				},
			},
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/version",
					Handler: version.Get(serverCtx),
				},
			},
			rest.WithPrefix("/api/v1"),
		)
	}

}
