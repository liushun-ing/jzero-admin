// Code generated by jzero. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest"
	"server/internal/svc"

	auth "server/internal/handler/auth"
	route "server/internal/handler/route"
	system_managemenu "server/internal/handler/system_manage/menu"
	system_managerole "server/internal/handler/system_manage/role"
	system_manageuser "server/internal/handler/system_manage/user"
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
					Path:    "/auth/error",
					Handler: auth.Error(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/auth/getUserInfo",
					Handler: auth.GetUserInfo(serverCtx),
				},
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
			},
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
					Handler: system_managemenu.GetAllPages(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getMenuList",
					Handler: system_managemenu.List(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getMenuTree",
					Handler: system_managemenu.Tree(serverCtx),
				},
			},
			rest.WithPrefix("/api/v1"),
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getAllRoles",
					Handler: system_managerole.GetAll(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getRoleList",
					Handler: system_managerole.List(serverCtx),
				},
			},
			rest.WithPrefix("/api/v1"),
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getUserList",
					Handler: system_manageuser.List(serverCtx),
				},
			},
			rest.WithPrefix("/api/v1"),
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/get",
					Handler: version.Get(serverCtx),
				},
			},
			rest.WithPrefix("/api/v1"),
		)
	}

}
