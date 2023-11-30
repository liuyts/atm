package logic

import (
	"ATM/internal/common/consts"
	"ATM/internal/common/utils"
	"ATM/internal/svc"
	"ATM/internal/types"
	"ATM/model"
	"context"
	"math/rand"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// 生成银行卡号
	cardNumber, err := l.generateBankCardNumber(req.BankName)
	if err != nil {
		return nil, err
	}

	// 创建用户
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Name:          req.Name,
		AccountNumber: cardNumber,
		Password:      utils.EncryptPassword(req.Password), // 对密码进行hash加密
		BankName:      req.BankName,
		IdCard:        req.IdCard,
		Phone:         req.Phone,
	})

	resp = new(types.UserRegisterResponse)
	resp.UserId, _ = res.LastInsertId()
	resp.AccountNumber = cardNumber

	return
}

func (l *UserRegisterLogic) generateBankCardNumber(bankName string) (string, error) {
	// 根据银行名称生成对应的发卡行标识
	bankIdentifier, ok := consts.BankNameNum[bankName]
	if !ok {
		return "", consts.ErrBankNameNotFound
	}

	// 生成校验位
	cardNumber := bankIdentifier + l.generateRandomDigits(10)
	checkDigit := calculateCheckDigit(cardNumber)

	return cardNumber + checkDigit, nil
}

func (l *UserRegisterLogic) generateRandomDigits(length int) string {
	digits := ""
	for i := 0; i < length; i++ {
		digits += strconv.Itoa(rand.Intn(10))
	}
	return digits
}

func calculateCheckDigit(cardNumber string) string {
	// 使用Luhn算法计算校验位
	sum := 0
	parity := len(cardNumber) % 2
	for i, digit := range cardNumber {
		d := int(digit - '0')
		if i%2 == parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	checkDigit := (10 - (sum % 10)) % 10
	return strconv.Itoa(checkDigit)
}
