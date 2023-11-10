package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TransactionModel = (*customTransactionModel)(nil)

type (
	// TransactionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTransactionModel.
	TransactionModel interface {
		transactionModel
		FindPage(ctx context.Context, meId int64, pageNum int64, pageSize int64) (int64, []*Transaction, error)
	}

	customTransactionModel struct {
		*defaultTransactionModel
	}
)

func (m *customTransactionModel) FindPage(ctx context.Context, meId int64, pageNum int64, pageSize int64) (int64, []*Transaction, error) {
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	query := fmt.Sprintf("select count(*) from %s where user_id = ?", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query, meId)
	if err != nil {
		return 0, nil, err
	}

	query = fmt.Sprintf("select %s from %s where user_id = ? order by create_time desc limit ?,?", transactionRows, m.table)
	var resp []*Transaction
	err = m.conn.QueryRowsCtx(ctx, &resp, query, meId, (pageNum-1)*pageSize, pageSize)
	return count, resp, err
}

// NewTransactionModel returns a model for the database table.
func NewTransactionModel(conn sqlx.SqlConn) TransactionModel {
	return &customTransactionModel{
		defaultTransactionModel: newTransactionModel(conn),
	}
}
