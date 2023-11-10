package logic

import (
	"ATM/internal/common/consts"
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"

	"ATM/internal/svc"
	"ATM/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	meId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()

	dbUser, err := l.svcCtx.UserModel.FindOne(l.ctx, meId)
	if err != nil {
		return
	}

	resp = new(types.UserInfoResponse)
	resp.User = new(types.User)
	_ = copier.Copy(resp.User, dbUser)

	return
}
