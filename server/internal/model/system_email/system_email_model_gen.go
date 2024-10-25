// Code generated by goctl. DO NOT EDIT.

package system_email

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	systemEmailFieldNames          = builder.RawFieldNames(&SystemEmail{})
	systemEmailRows                = strings.Join(systemEmailFieldNames, ",")
	systemEmailRowsExpectAutoSet   = strings.Join(stringx.Remove(systemEmailFieldNames, "`id`"), ",")
	systemEmailRowsWithPlaceHolder = strings.Join(stringx.Remove(systemEmailFieldNames, "`id`"), "=?,") + "=?"
)

type (
	systemEmailModel interface {
		Insert(ctx context.Context, data *SystemEmail) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*SystemEmail, error)
		Update(ctx context.Context, data *SystemEmail) error
		Delete(ctx context.Context, id uint64) error

		// custom interface generated by jzero
		BulkInsert(ctx context.Context, datas []*SystemEmail) error
		FindByCondition(ctx context.Context, conds ...condition.Condition) ([]*SystemEmail, error)
		FindOneByCondition(ctx context.Context, conds ...condition.Condition) (*SystemEmail, error)
		PageByCondition(ctx context.Context, conds ...condition.Condition) ([]*SystemEmail, int64, error)
		UpdateFieldsByCondition(ctx context.Context, field map[string]any, conds ...condition.Condition) error
		DeleteByCondition(ctx context.Context, conds ...condition.Condition) error
	}

	defaultSystemEmailModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SystemEmail struct {
		Id         uint64        `db:"id"`
		CreateTime time.Time     `db:"create_time"`
		UpdateTime time.Time     `db:"update_time"`
		CreateBy   sql.NullInt64 `db:"create_by"`
		UpdateBy   sql.NullInt64 `db:"update_by"`
		From       string        `db:"from"`
		Host       string        `db:"host"`
		Port       int64         `db:"port"`
		Username   string        `db:"username"`
		Password   string        `db:"password"`
		EnableSsl  int64         `db:"enable_ssl"`
		IsVerify   int64         `db:"is_verify"`
	}
)

func newSystemEmailModel(conn sqlx.SqlConn) *defaultSystemEmailModel {
	return &defaultSystemEmailModel{
		conn:  conn,
		table: "`system_email`",
	}
}

func (m *defaultSystemEmailModel) Delete(ctx context.Context, id uint64) error {
	sb := sqlbuilder.DeleteFrom(m.table)
	sb.Where(sb.EQ("`id`", id))
	sql, args := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	return err
}

func (m *defaultSystemEmailModel) FindOne(ctx context.Context, id uint64) (*SystemEmail, error) {
	sb := sqlbuilder.Select(systemEmailRows).From(m.table)
	sb.Where(sb.EQ("`id`", id))
	sb.Limit(1)
	sql, args := sb.Build()
	var resp SystemEmail
	err := m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSystemEmailModel) Insert(ctx context.Context, data *SystemEmail) (sql.Result, error) {
	sql, args := sqlbuilder.NewInsertBuilder().
		InsertInto(m.table).
		Cols(systemEmailRowsExpectAutoSet).
		Values(data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.From, data.Host, data.Port, data.Username, data.Password, data.EnableSsl, data.IsVerify).Build()
	ret, err := m.conn.ExecCtx(ctx, sql, args...)
	return ret, err
}

func (m *defaultSystemEmailModel) Update(ctx context.Context, data *SystemEmail) error {
	sb := sqlbuilder.Update(m.table)
	split := strings.Split(systemEmailRowsExpectAutoSet, ",")
	var assigns []string
	for _, s := range split {
		assigns = append(assigns, sb.Assign(s, nil))
	}
	sb.Set(assigns...)
	sb.Where(sb.EQ("`id`", nil))
	sql, _ := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.From, data.Host, data.Port, data.Username, data.Password, data.EnableSsl, data.IsVerify, data.Id)
	return err
}

func (m *defaultSystemEmailModel) tableName() string {
	return m.table
}

func (m *customSystemEmailModel) BulkInsert(ctx context.Context, datas []*SystemEmail) error {
	sb := sqlbuilder.InsertInto(m.table)
	sb.Cols(systemEmailRowsExpectAutoSet)
	for _, data := range datas {
		sb.Values(data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.From, data.Host, data.Port, data.Username, data.Password, data.EnableSsl, data.IsVerify)
	}

	sql, args := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	return err
}

func (m *customSystemEmailModel) FindByCondition(ctx context.Context, conds ...condition.Condition) ([]*SystemEmail, error) {
	sb := sqlbuilder.Select(systemEmailFieldNames...).From(m.table)
	condition.ApplySelect(sb, conds...)
	var resp []*SystemEmail

	sql, args := sb.Build()
	err := m.conn.QueryRowCtx(ctx, &resp, sql, args...)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customSystemEmailModel) FindOneByCondition(ctx context.Context, conds ...condition.Condition) (*SystemEmail, error) {
	sb := sqlbuilder.Select(systemEmailFieldNames...).From(m.table)
	condition.ApplySelect(sb, conds...)
	sb.Limit(1)
	var resp SystemEmail
	sql, args := sb.Build()
	err := m.conn.QueryRowCtx(ctx, &resp, sql, args...)

	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *customSystemEmailModel) PageByCondition(ctx context.Context, conds ...condition.Condition) ([]*SystemEmail, int64, error) {
	sb := sqlbuilder.Select(systemEmailFieldNames...).From(m.table)
	countsb := sqlbuilder.Select("count(*)").From(m.table)

	condition.ApplySelect(sb, conds...)

	var countConds []condition.Condition
	for _, cond := range conds {
		if cond.Operator != condition.Limit && cond.Operator != condition.Offset {
			countConds = append(countConds, cond)
		}
	}
	condition.ApplySelect(countsb, countConds...)

	var resp []*SystemEmail

	sql, args := sb.Build()
	err := m.conn.QueryRowsCtx(ctx, &resp, sql, args...)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	sql, args = countsb.Build()
	err = m.conn.QueryRowCtx(ctx, &total, sql, args...)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func (m *customSystemEmailModel) UpdateFieldsByCondition(ctx context.Context, field map[string]any, conds ...condition.Condition) error {
	if field == nil {
		return nil
	}

	sb := sqlbuilder.Update(m.table)
	condition.ApplyUpdate(sb, conds...)

	var assigns []string
	for key, value := range field {
		assigns = append(assigns, sb.Assign(key, value))
	}
	sb.Set(assigns...)

	sql, args := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (m *customSystemEmailModel) DeleteByCondition(ctx context.Context, conds ...condition.Condition) error {
	if len(conds) == 0 {
		return nil
	}
	sb := sqlbuilder.DeleteFrom(m.table)
	condition.ApplyDelete(sb, conds...)
	sql, args := sb.Build()
	_, err := m.conn.ExecCtx(ctx, sql, args...)
	return err
}
