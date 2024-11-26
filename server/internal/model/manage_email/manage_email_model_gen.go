// Code generated by goctl. Templates Edited by jzero. DO NOT EDIT.

package manage_email

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/eddieowens/opts"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	manageEmailFieldNames          = builder.RawFieldNames(&ManageEmail{})
	manageEmailRows                = strings.Join(manageEmailFieldNames, ",")
	manageEmailRowsExpectAutoSet   = strings.Join(stringx.Remove(manageEmailFieldNames, "`id`"), ",")
	manageEmailRowsWithPlaceHolder = strings.Join(stringx.Remove(manageEmailFieldNames, "`id`"), "=?,") + "=?"
)

type (
	manageEmailModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *ManageEmail) (sql.Result, error)
		InsertWithCache(ctx context.Context, session sqlx.Session, data *ManageEmail) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id uint64) (*ManageEmail, error)
		FindOneWithCache(ctx context.Context, session sqlx.Session, id uint64) (*ManageEmail, error)
		Update(ctx context.Context, session sqlx.Session, data *ManageEmail) error
		UpdateWithCache(ctx context.Context, session sqlx.Session, data *ManageEmail) error
		Delete(ctx context.Context, session sqlx.Session, id uint64) error
		DeleteWithCache(ctx context.Context, session sqlx.Session, id uint64) error

		// custom interface generated by jzero
		BulkInsert(ctx context.Context, session sqlx.Session, datas []*ManageEmail) error
		FindByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) ([]*ManageEmail, error)
		FindOneByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) (*ManageEmail, error)
		PageByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) ([]*ManageEmail, int64, error)
		UpdateFieldsByCondition(ctx context.Context, session sqlx.Session, field map[string]any, conds ...condition.Condition) error
		DeleteByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) error
	}

	defaultManageEmailModel struct {
		cachedConn sqlc.CachedConn
		conn       sqlx.SqlConn
		table      string
	}

	ManageEmail struct {
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

func newManageEmailModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) *defaultManageEmailModel {
	o := opts.DefaultApply(op...)
	var cachedConn sqlc.CachedConn
	if len(o.CacheConf) > 0 {
		cachedConn = sqlc.NewConn(conn, o.CacheConf, o.CacheOpts...)
	}
	if o.CachedConn != nil {
		cachedConn = *o.CachedConn
	}
	return &defaultManageEmailModel{
		cachedConn: cachedConn,
		conn:       conn,
		table:      "`manage_email`",
	}
}
func (m *defaultManageEmailModel) Delete(ctx context.Context, session sqlx.Session, id uint64) error {
	sb := sqlbuilder.DeleteFrom(m.table)
	sb.Where(sb.EQ("`id`", id))
	statement, args := sb.Build()
	var err error
	if session != nil {
		_, err = session.ExecCtx(ctx, statement, args...)
	} else {
		_, err = m.conn.ExecCtx(ctx, statement, args...)
	}
	return err
}

func (m *defaultManageEmailModel) DeleteWithCache(ctx context.Context, session sqlx.Session, id uint64) error {
	return m.Delete(ctx, session, id)
}

func (m *defaultManageEmailModel) FindOne(ctx context.Context, session sqlx.Session, id uint64) (*ManageEmail, error) {
	sb := sqlbuilder.Select(manageEmailRows).From(m.table)
	sb.Where(sb.EQ("`id`", id))
	sb.Limit(1)
	sql, args := sb.Build()
	var resp ManageEmail
	var err error
	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, sql, args...)
	} else {
		err = m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	}
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultManageEmailModel) FindOneWithCache(ctx context.Context, session sqlx.Session, id uint64) (*ManageEmail, error) {
	return m.FindOne(ctx, session, id)
}

func (m *defaultManageEmailModel) Insert(ctx context.Context, session sqlx.Session, data *ManageEmail) (sql.Result, error) {
	statement, args := sqlbuilder.NewInsertBuilder().
		InsertInto(m.table).
		Cols(manageEmailRowsExpectAutoSet).
		Values(data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.From, data.Host, data.Port, data.Username, data.Password, data.EnableSsl, data.IsVerify).Build()
	if session != nil {
		return session.ExecCtx(ctx, statement, args...)
	}
	return m.conn.ExecCtx(ctx, statement, args...)
}

func (m *defaultManageEmailModel) InsertWithCache(ctx context.Context, session sqlx.Session, data *ManageEmail) (sql.Result, error) {
	return m.Insert(ctx, session, data)
}
func (m *defaultManageEmailModel) Update(ctx context.Context, session sqlx.Session, data *ManageEmail) error {
	sb := sqlbuilder.Update(m.table)
	split := strings.Split(manageEmailRowsExpectAutoSet, ",")
	var assigns []string
	for _, s := range split {
		assigns = append(assigns, sb.Assign(s, nil))
	}
	sb.Set(assigns...)
	sb.Where(sb.EQ("`id`", nil))
	statement, _ := sb.Build()

	var err error
	if session != nil {
		_, err = session.ExecCtx(ctx, statement, data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.From, data.Host, data.Port, data.Username, data.Password, data.EnableSsl, data.IsVerify, data.Id)
	} else {
		_, err = m.conn.ExecCtx(ctx, statement, data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.From, data.Host, data.Port, data.Username, data.Password, data.EnableSsl, data.IsVerify, data.Id)
	}
	return err
}

func (m *defaultManageEmailModel) UpdateWithCache(ctx context.Context, session sqlx.Session, data *ManageEmail) error {
	return m.Update(ctx, session, data)
}

func (m *defaultManageEmailModel) tableName() string {
	return m.table
}

func (m *customManageEmailModel) BulkInsert(ctx context.Context, session sqlx.Session, datas []*ManageEmail) error {
	sb := sqlbuilder.InsertInto(m.table)
	sb.Cols(manageEmailRowsExpectAutoSet)
	for _, data := range datas {
		sb.Values(data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.From, data.Host, data.Port, data.Username, data.Password, data.EnableSsl, data.IsVerify)
	}
	statement, args := sb.Build()

	var err error
	if session != nil {
		_, err = session.ExecCtx(ctx, statement, args...)
	} else {
		_, err = m.conn.ExecCtx(ctx, statement, args...)
	}
	return err
}

func (m *customManageEmailModel) FindByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) ([]*ManageEmail, error) {
	sb := sqlbuilder.Select(manageEmailFieldNames...).From(m.table)
	builder := condition.Select(*sb, conds...)
	statement, args := builder.Build()

	var resp []*ManageEmail
	var err error

	if session != nil {
		err = session.QueryRowsCtx(ctx, &resp, statement, args...)
	} else {
		err = m.conn.QueryRowsCtx(ctx, &resp, statement, args...)
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customManageEmailModel) FindOneByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) (*ManageEmail, error) {
	sb := sqlbuilder.Select(manageEmailFieldNames...).From(m.table)

	builder := condition.Select(*sb, conds...)
	builder.Limit(1)
	statement, args := builder.Build()

	var resp ManageEmail
	var err error

	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, statement, args...)
	} else {
		err = m.conn.QueryRowCtx(ctx, &resp, statement, args...)
	}
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *customManageEmailModel) PageByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) ([]*ManageEmail, int64, error) {
	sb := sqlbuilder.Select(manageEmailFieldNames...).From(m.table)
	countsb := sqlbuilder.Select("count(*)").From(m.table)

	builder := condition.Select(*sb, conds...)

	var countConds []condition.Condition
	for _, cond := range conds {
		if cond.Operator != condition.Limit && cond.Operator != condition.Offset {
			countConds = append(countConds, cond)
		}
	}
	countBuilder := condition.Select(*countsb, countConds...)

	var resp []*ManageEmail
	var err error

	statement, args := builder.Build()

	if session != nil {
		err = session.QueryRowsCtx(ctx, &resp, statement, args...)
	} else {
		err = m.conn.QueryRowsCtx(ctx, &resp, statement, args...)
	}
	if err != nil {
		return nil, 0, err
	}

	var total int64
	statement, args = countBuilder.Build()
	if session != nil {
		err = session.QueryRowCtx(ctx, &total, statement, args...)
	} else {
		err = m.conn.QueryRowCtx(ctx, &total, statement, args...)
	}
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func (m *customManageEmailModel) UpdateFieldsByCondition(ctx context.Context, session sqlx.Session, field map[string]any, conds ...condition.Condition) error {
	if field == nil {
		return nil
	}

	sb := sqlbuilder.Update(m.table)
	builder := condition.Update(*sb, conds...)

	var assigns []string
	for key, value := range field {
		assigns = append(assigns, sb.Assign(key, value))
	}
	builder.Set(assigns...)

	statement, args := builder.Build()

	var err error
	if session != nil {
		_, err = session.ExecCtx(ctx, statement, args...)
	} else {
		_, err = m.conn.ExecCtx(ctx, statement, args...)
	}
	if err != nil {
		return err
	}
	return nil
}

func (m *customManageEmailModel) DeleteByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) error {
	if len(conds) == 0 {
		return nil
	}
	sb := sqlbuilder.DeleteFrom(m.table)
	builder := condition.Delete(*sb, conds...)
	statement, args := builder.Build()

	var err error
	if session != nil {
		_, err = session.ExecCtx(ctx, statement, args...)
	} else {
		_, err = m.conn.ExecCtx(ctx, statement, args...)
	}
	return err
}