package logic

import (
	"ATM/internal/common/consts"
	"ATM/internal/common/utils"
	"context"
	"encoding/json"

	"ATM/internal/svc"
	"ATM/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordRequest) (resp *types.ChangePasswordResponse, err error) {
	meId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()
	// 查找用户信息
	dbUser, err := l.svcCtx.UserModel.FindOne(l.ctx, meId)
	if err != nil {
		return
	}
	// 对修改的密码进行hash加密
	dbUser.Password = utils.EncryptPassword(req.NewPassword)
	// 更新用户信息
	if err = l.svcCtx.UserModel.Update(l.ctx, dbUser); err != nil {
		return
	}

	resp = new(types.ChangePasswordResponse)

	return
}
