package logic

import (
	"ATM/internal/common/consts"
	"ATM/internal/svc"
	"ATM/internal/types"
	"ATM/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutMoneyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutMoneyLogic {
	return &PutMoneyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutMoneyLogic) PutMoney(req *types.PutMoneyRequest) (resp *types.PutMoneyResponse, err error) {
	// 检查金额是否精确到两位小数
	amountStr := strconv.FormatFloat(req.Amount, 'f', 2, 64)
	req.Amount, err = strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return nil, fmt.Errorf("存款金额不合法")
	}

	meId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()
	dbUser, err := l.svcCtx.UserModel.FindOne(l.ctx, meId)
	if err != nil {
		l.Errorf("查询用户失败，err：%v", err)
		return nil, err
	}

	// 更新余额
	dbUser.Balance += req.Amount
	err = l.svcCtx.Tx.Transact(func(session sqlx.Session) error {
		if err := l.svcCtx.UserModel.TxUpdate(l.ctx, session, dbUser); err != nil {
			l.Errorf("更新用户余额失败，err：%v", err)
			return err
		}
		// 插入交易记录
		if _, err := l.svcCtx.TransactionModel.TxInsert(l.ctx, session, &model.Transaction{
			Amount:      req.Amount,
			Type:        consts.TransactionTypePut,
			Description: fmt.Sprintf("您存入%v元", req.Amount),
			UserId:      meId,
		}); err != nil {
			l.Errorf("插入交易记录失败，err：%v", err)
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	resp = new(types.PutMoneyResponse)
	resp.Balance = dbUser.Balance

	return
}
