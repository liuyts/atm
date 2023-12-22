package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DailyLimitModel = (*customDailyLimitModel)(nil)

type (
	// DailyLimitModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDailyLimitModel.
	DailyLimitModel interface {
		dailyLimitModel
		FindPage(ctx context.Context, pageNum int, pageSize int) ([]*DailyLimit, error)
		FindByUserIdDate(ctx context.Context, userId int64, date string) (*DailyLimit, error)
	}

	customDailyLimitModel struct {
		*defaultDailyLimitModel
	}
)

func (m *customDailyLimitModel) FindByUserIdDate(ctx context.Context, userId int64, date string) (*DailyLimit, error) {
	query := fmt.Sprintf("select %s from %s where user_id=? and date=?", dailyLimitRows, m.table)
	var resp DailyLimit
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId, date)

	return &resp, err
}

func (m *customDailyLimitModel) FindPage(ctx context.Context, pageNum int, pageSize int) ([]*DailyLimit, error) {
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", dailyLimitRows, m.table)
	var resp []*DailyLimit
	err := m.conn.QueryRowsCtx(ctx, &resp, query, (pageNum-1)*pageSize, pageSize)

	return resp, err
}

// NewDailyLimitModel returns a model for the database table.
func NewDailyLimitModel(conn sqlx.SqlConn) DailyLimitModel {
	return &customDailyLimitModel{
		defaultDailyLimitModel: newDailyLimitModel(conn),
	}
}
