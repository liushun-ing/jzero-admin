package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/system/role"
)

type GetAll struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAll(ctx context.Context, svcCtx *svc.ServiceContext) *GetAll {
	return &GetAll{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAll) GetAll(req *types.GetAllRequest) (resp []types.GetAllResponse, err error) {
	roles, err := l.svcCtx.Model.SystemRole.FindByCondition(l.ctx)
	if err != nil {
		return nil, err
	}
	var list []types.GetAllResponse

	for _, role := range roles {
		if role.Status == "1" {
			list = append(list, types.GetAllResponse{
				Id:       role.Id,
				RoleCode: role.Code,
				RoleName: role.Name,
			})
		}
	}

	return list, nil
}
