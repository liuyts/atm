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

	"github.com/zeromicro/go-zero/core/logx"
)

type TakeMoneyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTakeMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TakeMoneyLogic {
	return &TakeMoneyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TakeMoneyLogic) TakeMoney(req *types.TakeMoneyRequest) (resp *types.TakeMoneyResponse, err error) {
	meId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()
	dbUser, err := l.svcCtx.UserModel.FindOne(l.ctx, meId)
	if err != nil {
		l.Errorf("查询用户失败，err：%v", err)
		return nil, err
	}

	// 更新余额
	if dbUser.Balance < req.Amount {
		l.Errorf("余额不足，err：%v", err)
		return nil, consts.ErrBalanceNotEnough
	}
	dbUser.Balance -= req.Amount
	err = l.svcCtx.Tx.Transact(func(session sqlx.Session) error {
		if err := l.svcCtx.UserModel.TxUpdate(l.ctx, session, dbUser); err != nil {
			l.Errorf("更新用户余额失败，err：%v", err)
			return err
		}
		// 插入交易记录
		if _, err := l.svcCtx.TransactionModel.TxInsert(l.ctx, session, &model.Transaction{
			Amount:      req.Amount,
			Type:        consts.TransactionTypeTake,
			Description: fmt.Sprintf("您取出%v元", req.Amount),
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

	resp = new(types.TakeMoneyResponse)
	resp.Balance = dbUser.Balance

	return
}
