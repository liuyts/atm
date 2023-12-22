package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"ATM/internal/svc"
	"ATM/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) (resp *types.UpdateUserResponse, err error) {
	dbUser, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return
	}

	_ = copier.Copy(dbUser, req)

	err = l.svcCtx.UserModel.Update(l.ctx, dbUser)
	if err != nil {
		return
	}

	resp = new(types.UpdateUserResponse)

	return
}
