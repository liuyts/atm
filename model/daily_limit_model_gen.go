// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	dailyLimitFieldNames          = builder.RawFieldNames(&DailyLimit{})
	dailyLimitRows                = strings.Join(dailyLimitFieldNames, ",")
	dailyLimitRowsExpectAutoSet   = strings.Join(stringx.Remove(dailyLimitFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dailyLimitRowsWithPlaceHolder = strings.Join(stringx.Remove(dailyLimitFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dailyLimitModel interface {
		Insert(ctx context.Context, data *DailyLimit) (sql.Result, error)
		TxInsert(ctx context.Context, tx sqlx.Session, data *DailyLimit) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DailyLimit, error)
		Update(ctx context.Context, data *DailyLimit) error
		TxUpdate(ctx context.Context, tx sqlx.Session, data *DailyLimit) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDailyLimitModel struct {
		conn  sqlx.SqlConn
		table string
	}

	DailyLimit struct {
		Id     int64   `db:"id"`
		UserId int64   `db:"user_id"`
		Limit  float64 `db:"limit"`
		Date   string  `db:"date"`
	}
)

func newDailyLimitModel(conn sqlx.SqlConn) *defaultDailyLimitModel {
	return &defaultDailyLimitModel{
		conn:  conn,
		table: "`daily_limit`",
	}
}

func (m *defaultDailyLimitModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultDailyLimitModel) FindOne(ctx context.Context, id int64) (*DailyLimit, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dailyLimitRows, m.table)
	var resp DailyLimit
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDailyLimitModel) Insert(ctx context.Context, data *DailyLimit) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, dailyLimitRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Limit, data.Date)
	return ret, err
}

func (m *defaultDailyLimitModel) TxInsert(ctx context.Context, tx sqlx.Session, data *DailyLimit) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, dailyLimitRowsExpectAutoSet)
	ret, err := tx.ExecCtx(ctx, query, data.UserId, data.Limit, data.Date)
	return ret, err
}

func (m *defaultDailyLimitModel) Update(ctx context.Context, data *DailyLimit) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, dailyLimitRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Limit, data.Date, data.Id)
	return err
}

func (m *defaultDailyLimitModel) TxUpdate(ctx context.Context, tx sqlx.Session, data *DailyLimit) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, dailyLimitRowsWithPlaceHolder)
	_, err := tx.ExecCtx(ctx, query, data.UserId, data.Limit, data.Date, data.Id)
	return err
}

func (m *defaultDailyLimitModel) tableName() string {
	return m.table
}
