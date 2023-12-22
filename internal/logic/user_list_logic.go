package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"ATM/internal/svc"
	"ATM/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListRequest) (resp *types.UserListResponse, err error) {
	count, users, err := l.svcCtx.UserModel.FindPage(l.ctx, req.PageNum, req.PageSize)
	if err != nil {
		return
	}

	resp = new(types.UserListResponse)
	resp.Users = make([]*types.User, 0, len(users))
	resp.Total = count
	_ = copier.Copy(&resp.Users, &users)

	return
}
