package logic

import (
	"ATM/internal/common/consts"
	"context"
	"encoding/json"
	"time"

	"ATM/internal/svc"
	"ATM/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTransactionMoneyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTransactionMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTransactionMoneyLogic {
	return &GetTransactionMoneyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTransactionMoneyLogic) GetTransactionMoney(req *types.GetTransactionMoneyRequest) (resp *types.GetTransactionMoneyResponse, err error) {
	meId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()
	// 获取流水记录
	count, transactions, err := l.svcCtx.TransactionModel.FindPage(l.ctx, meId, req.PageNum, req.PageSize)
	if err != nil {
		l.Errorf("分页查询流水记录失败，err：%v")
		return nil, err
	}

	resp = new(types.GetTransactionMoneyResponse)
	resp.Total = count
	resp.Transactions = make([]*types.Transaction, 0, len(transactions))
	for _, transaction := range transactions {
		resp.Transactions = append(resp.Transactions, &types.Transaction{
			Id:          transaction.Id,
			UserId:      transaction.UserId,
			Type:        transaction.Type,
			Amount:      transaction.Amount,
			Description: transaction.Description,
			CreateTime:  transaction.CreateTime.Format(time.DateTime),
		})
	}

	return
}
