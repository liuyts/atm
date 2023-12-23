package logic

import (
	"ATM/internal/common/consts"
	"ATM/internal/svc"
	"ATM/internal/types"
	"ATM/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferMoneyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferMoneyLogic {
	return &TransferMoneyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferMoneyLogic) TransferMoney(req *types.TransferMoneyRequest) (resp *types.TransferMoneyResponse, err error) {
	meId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()
	// 查询用户余额
	meUser, err := l.svcCtx.UserModel.FindOne(l.ctx, meId)
	if err != nil {
		return
	}
	if meUser.Balance < req.Amount {
		return nil, consts.ErrInsufficientBalance
	}
	// 查询对方是否存在
	toUser, err := l.svcCtx.UserModel.FindOneByAccountNumber(l.ctx, req.ToAccountNumber)
	if err != nil {
		return
	}

	if toUser.Id == meId {
		return nil, consts.ErrTransferToSelf
	}

	// 查看每日限额
	date := time.Now().Format(time.DateOnly)
	dailyLimit, err := l.svcCtx.DailyLimitModel.FindByUserIdDate(l.ctx, meId, date)
	if err != nil {
		return
	}
	if req.Amount > dailyLimit.Limit {
		return nil, errors.New("转账额度超过每日限额")
	}

	// 开启事务转账
	err = l.svcCtx.Tx.Transact(func(session sqlx.Session) (err error) {
		meUser.Balance -= req.Amount
		err = l.svcCtx.UserModel.TxUpdate(l.ctx, session, meUser)
		if err != nil {
			return
		}

		toUser.Balance += req.Amount
		err = l.svcCtx.UserModel.TxUpdate(l.ctx, session, toUser)
		if err != nil {
			return
		}

		_, err = l.svcCtx.TransactionModel.TxInsert(l.ctx, session, &model.Transaction{
			Amount:      req.Amount,
			Type:        consts.TransactionTypeTransfer,
			Description: fmt.Sprintf("你转账给%v(%v)%v元", toUser.Name, toUser.AccountNumber, req.Amount),
			UserId:      meId,
		})
		if err != nil {
			return
		}

		_, err = l.svcCtx.TransactionModel.TxInsert(l.ctx, session, &model.Transaction{
			Amount:      req.Amount,
			Type:        consts.TransactionTypeTransfer,
			Description: fmt.Sprintf("%v(%v)转账给你%v元", meUser.Name, meUser.AccountNumber, req.Amount),
			UserId:      toUser.Id,
		})
		if err != nil {
			return
		}

		// 更新每日限额
		dailyLimit.Limit -= req.Amount
		err = l.svcCtx.DailyLimitModel.TxUpdate(l.ctx, session, dailyLimit)
		if err != nil {
			return
		}

		return
	})

	resp = new(types.TransferMoneyResponse)

	return
}
