package logic

import (
	"ATM/internal/common/consts"
	"ATM/model"
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"ATM/internal/svc"
	"ATM/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBalanceLogic {
	return &GetBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBalanceLogic) GetBalance(req *types.GetBalanceRequest) (resp *types.GetBalanceResponse, err error) {
	meId, _ := l.ctx.Value(consts.UserId).(json.Number).Int64()

	dbUser, err := l.svcCtx.UserModel.FindOne(l.ctx, meId)
	date := time.Now().Format("2006-01-02")
	dailyLimit, err := l.svcCtx.DailyLimitModel.FindByUserIdDate(l.ctx, meId, date)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			// 如果没有找到，创建新的每日限额
			dailyLimit = &model.DailyLimit{
				UserId: meId,
				Date:   date,
				Limit:  dbUser.DailyLimit,
			}
			_, err = l.svcCtx.DailyLimitModel.Insert(l.ctx, dailyLimit)
			if err != nil {
				return
			}
		} else {
			return
		}
	}

	resp = new(types.GetBalanceResponse)
	resp.TotalBalance = dbUser.Balance
	resp.DailyLimit = dailyLimit.Limit
	resp.Date = dailyLimit.Date

	return
}
