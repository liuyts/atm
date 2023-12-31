package logic

import (
	"ATM/internal/common/consts"
	"ATM/internal/common/utils"
	"ATM/model"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"ATM/internal/svc"
	"ATM/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	// 根据银行卡号查询用户信息
	dbUser, err := l.svcCtx.UserModel.FindOneByAccountNumber(l.ctx, req.AccountNumber)
	if errors.Is(err, model.ErrNotFound) {
		return nil, errors.New("银行卡号不存在")
	}
	if err != nil {
		return nil, err
	}
	// 将用户输入的密码进行hash加密后与数据库中的密码进行比对
	if !utils.VerifyPassword(req.Password, dbUser.Password) {
		return nil, errors.New("密码错误，请重新输入")
	}
	// 该卡是否被管理员封禁
	if dbUser.Status == "封禁" {
		return nil, errors.New("银行卡已被封禁，请联系管理员")
	}
	// 生成token，登录成功
	auth := l.svcCtx.Config.Auth
	now := time.Now().Unix()
	accessToken, _ := l.getJwtToken(auth.AccessSecret, now, auth.AccessExpire, dbUser.Id)

	resp = new(types.UserLoginResponse)
	resp.Token = accessToken
	resp.UserId = dbUser.Id
	resp.AccountNumber = dbUser.AccountNumber

	return
}

func (l *UserLoginLogic) getJwtToken(secretKey string, iat, seconds, userID int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[consts.UserId] = userID
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
