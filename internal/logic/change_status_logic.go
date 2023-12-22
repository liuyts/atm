package logic

import (
	"context"

	"ATM/internal/svc"
	"ATM/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeStatusLogic {
	return &ChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeStatusLogic) ChangeStatus(req *types.ChangeStatusRequest) (resp *types.ChangeStatusResponse, err error) {
	dbUser, err := l.svcCtx.UserModel.FindOneByAccountNumber(l.ctx, req.AccountNumber)
	dbUser.Status = req.Status
	err = l.svcCtx.UserModel.Update(l.ctx, dbUser)
	if err != nil {
		return
	}

	resp = new(types.ChangeStatusResponse)

	return
}
