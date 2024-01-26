// Code generated by wpb. DO NOT EDIT.
package db_user

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/jmoiron/sqlx"
	"github.com/walleframe/svc_db"
	"github.com/walleframe/svc_db/example/dbop"
	"github.com/walleframe/walle/util"
)

////////////////////////////////////////////////////////////////////////////////
// public interface

type UserTestOperation interface {
	Insert(ctx context.Context, data *dbop.UserTest) (res sql.Result, err error)
	InsertMany(ctx context.Context, datas []*dbop.UserTest) (res sql.Result, err error)

	FindByIndexUid(ctx context.Context, uid int64, limit, offset int) (datas []*dbop.UserTest, err error)
	FindExByIndexUid(ctx context.Context, uid int64, limit, offset int) (datas []*dbop.UserTestEx, err error)
	CountByIndexUid(ctx context.Context, uid int64) (count int, err error)
	DeleteByIndexUid(ctx context.Context, uid int64) (res sql.Result, err error)

	Where(bufSize int) *UserTestWhereStmt
	Select(ctx context.Context, where *UserTestWhereStmt) (datas []*dbop.UserTest, err error)
	SelectEx(ctx context.Context, where *UserTestWhereStmt) (datas []*dbop.UserTestEx, err error)

	DeleteMany(ctx context.Context, where *UserTestWhereStmt) (res sql.Result, err error)

	RangeAll(ctx context.Context, where *UserTestWhereStmt, f func(ctx context.Context, row *dbop.UserTest) bool) error
	RangeAllEx(ctx context.Context, where *UserTestWhereStmt, f func(ctx context.Context, row *dbop.UserTestEx) bool) error
	AllData(ctx context.Context, where *UserTestWhereStmt) (datas []*dbop.UserTest, err error)
	AllDataEx(ctx context.Context, where *UserTestWhereStmt) (datas []*dbop.UserTestEx, err error)

	// use for custom named sql
	DB() *sqlx.DB
}

var (
	UserTestUidUnamrshal   = svc_db.RawToInt64
	UserTestXxxUnamrshal   = svc_db.RawToInt64
	UserTestStateUnamrshal = svc_db.RawToInt8
)

var (
	globalUserTestOP atomic.Pointer[xUserTestOperation]
)

func init() {
	svc_db.RegisterDB("mysql", "db_user", "user_test", func(db *sqlx.DB) error {
		// sync table columns
		err := svc_db.SyncTableColumns(context.Background(), db, "user_test", UserTestSQL_Create, UserTestSQL_TableColumns)
		if err != nil {
			return fmt.Errorf("swap db_user.user_test pointer, sync columns failed, %w", err)
		}
		// sync table index
		err = svc_db.SyncTableIndex(context.Background(), db, "user_test", UserTestSQL_TableIndex)
		if err != nil {
			return fmt.Errorf("swap db_user.user_test pointer, sync index failed, %w", err)
		}
		//
		tableOP, err := NewUserTestOperation(db)
		if err != nil {
			return fmt.Errorf("swap db_user.user_test pointer, new table operation failed, %w", err)
		}

		globalUserTestOP.Store(tableOP)
		return nil
	})
}

var UserTestOP = func() UserTestOperation {
	return globalUserTestOP.Load()
}

func UserTestNamedSQL(bufSize int) *UserTestSQLWriter {
	sql := &UserTestSQLWriter{}
	sql.buf.Grow(bufSize)
	return sql
}

////////////////////////////////////////////////////////////////////////////////
// sql statement

const (
	UserTestSQL_Insert        = "insert user_test(`uid`,`xxx`,`state`) values(?,?,?)"
	UserTestSQL_InsertValues  = ",values(?,?,?)"
	UserTestSQL_InsertValues2 = ",values(?,?,?)"
	UserTestSQL_Delete        = "delete from user_test"
	UserTestSQL_Find          = "select user_test(`uid`,`xxx`,`state`) from user_test"
	UserTestSQL_FindRow       = "select user_test(`uid`,`xxx`,`state`,`modify_stamp`,`create_stamp`) from user_test"
	UserTestSQL_Count         = "select count(id) from user_test"
	UserTestSQL_Create        = "create table user_test (" +
		"`uid` bigint not null default 0," +
		"`xxx` bigint not null default 0," +
		"`state` tinyint not null default 0," +
		"`modify_stamp` timestamp default current_timestamp on update current_timestamp," +
		"`create_stamp` timestamp default current_timestamp" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;"
)

var (
	UserTestSQL_TableColumns = map[string]string{
		"uid":   "alter table user_test add `uid` bigint not null default 0;",
		"xxx":   "alter table user_test add `xxx` bigint not null default 0;",
		"state": "alter table user_test add `state` tinyint not null default 0;",
	}
	UserTestSQL_TableIndex = map[string]string{
		"user_test_uid": "create unique index user_test_uid on user_test(`uid`)",
	}
)

////////////////////////////////////////////////////////////////////////////////
// TblTestOperation impl

type xUserTestOperation struct {
	db           *sqlx.DB
	insert       *sql.Stmt
	idxUidFind   *sql.Stmt
	idxUidFindEx *sql.Stmt
	idxUidCount  *sql.Stmt
	idxUidDelete *sql.Stmt
}

func NewUserTestOperation(db *sqlx.DB) (_ *xUserTestOperation, err error) {
	t := &xUserTestOperation{
		db: db,
	}

	t.insert, err = db.Prepare(UserTestSQL_Insert)
	if err != nil {
		return nil, fmt.Errorf("prepare db_user.user_test insert failed,%w", err)
	}

	return t, nil
}

func (t *xUserTestOperation) Insert(ctx context.Context, data *dbop.UserTest) (res sql.Result, err error) {
	res, err = t.insert.ExecContext(ctx, data.Uid, data.Xxx, data.State)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_test insert failed,%w", err)
	}

	return
}

func (t *xUserTestOperation) InsertMany(ctx context.Context, datas []*dbop.UserTest) (res sql.Result, err error) {
	switch len(datas) {
	case 0:
		return svc_db.EmptyResult{}, nil
	case 1:
		return t.Insert(ctx, datas[0])
	}

	buf := util.Builder{}
	buf.Grow(len(UserTestSQL_Insert) + (len(datas)-1)*len(UserTestSQL_InsertValues))
	buf.Write([]byte(UserTestSQL_Insert))
	for i := 0; i < len(datas)-1; i++ {
		buf.Write([]byte(UserTestSQL_InsertValues))
	}
	args := make([]any, 0, len(datas)*3)
	for i := 0; i < len(datas); i++ {
		data := datas[i]
		args = append(args, data.Uid, data.Xxx, data.State)
	}
	res, err = t.db.DB.ExecContext(ctx, buf.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_test insert_many failed,%w", err)
	}
	return
}

func (t *xUserTestOperation) FindByIndexUid(ctx context.Context, uid int64, limit, offset int) (datas []*dbop.UserTest, err error) {
	if t.idxUidFind == nil {
		t.idxUidFind, err = t.db.PrepareContext(ctx, UserFriendSQL_Find+"where `uid`=? limit ?,?")
		if err != nil {
			return nil, fmt.Errorf("prepare db_user.user_test find_by_index_uid failed,%w", err)
		}
	}
	rows, err := t.idxUidFind.QueryContext(ctx, uid, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_test find_by_index_uid failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		data, err := scanUserTest(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return
}
func (t *xUserTestOperation) FindExByIndexUid(ctx context.Context, uid int64, limit, offset int) (datas []*dbop.UserTestEx, err error) {
	if t.idxUidFindEx == nil {
		t.idxUidFindEx, err = t.db.PrepareContext(ctx, UserFriendSQL_FindRow+"where `uid`=? limit ?,?")
		if err != nil {
			return nil, fmt.Errorf("prepare db_user.user_test findex_by_index_uid failed,%w", err)
		}
	}
	rows, err := t.idxUidFindEx.QueryContext(ctx, uid, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_test findex_by_index_uid failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		data, err := scanUserTestEx(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return
}
func (t *xUserTestOperation) CountByIndexUid(ctx context.Context, uid int64) (count int, err error) {
	if t.idxUidCount == nil {
		t.idxUidCount, err = t.db.PrepareContext(ctx, UserFriendSQL_Count+"where `uid`=?")
		if err != nil {
			return 0, fmt.Errorf("prepare db_user.user_test count_by_index_uid failed,%w", err)
		}
	}
	err = t.idxUidCount.QueryRowContext(ctx, uid).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("exec db_user.user_test count_by_index_uid failed,%w", err)
	}
	return
}

func (t *xUserTestOperation) DeleteByIndexUid(ctx context.Context, uid int64) (res sql.Result, err error) {
	if t.idxUidDelete == nil {
		t.idxUidDelete, err = t.db.PrepareContext(ctx, UserFriendSQL_Count+"where `uid`=?")
		if err != nil {
			return nil, fmt.Errorf("prepare db_user.user_test delete_by_index_uid failed,%w", err)
		}
	}
	res, err = t.idxUidDelete.ExecContext(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_test delete_by_index_uid failed,%w", err)
	}
	return
}

func (t *xUserTestOperation) Where(bufSize int) *UserTestWhereStmt {
	w := &UserTestWhereStmt{}
	w.buf.Grow(bufSize)
	return w
}

func (t *xUserTestOperation) Select(ctx context.Context, where *UserTestWhereStmt) (datas []*dbop.UserTest, err error) {
	where.applyLimitAndOffset()
	var findSql = UserTestSQL_Find + where.String()
	rows, err := t.db.QueryContext(ctx, findSql)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_test select failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {

		data, err := scanUserTest(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return
}
func (t *xUserTestOperation) SelectEx(ctx context.Context, where *UserTestWhereStmt) (datas []*dbop.UserTestEx, err error) {
	where.applyLimitAndOffset()
	var findSql = UserTestSQL_FindRow + where.String()
	rows, err := t.db.QueryContext(ctx, findSql)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_test selectex failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		data, err := scanUserTestEx(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return
}

func (t *xUserTestOperation) DeleteMany(ctx context.Context, where *UserTestWhereStmt) (res sql.Result, err error) {
	w := where.String()
	buf := util.Builder{}
	buf.Grow(len(UserTestSQL_Delete) + len(w))
	buf.Write([]byte(UserTestSQL_Delete))
	buf.WriteString(w)
	res, err = t.db.ExecContext(ctx, buf.String())
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_test delete_many failed,%w", err)
	}

	return
}

func (t *xUserTestOperation) RangeAll(ctx context.Context, where *UserTestWhereStmt, f func(ctx context.Context, row *dbop.UserTest) bool) error {
	var findSql = UserTestSQL_Find + where.String()
	limit := where.limit
	if limit == 0 {
		limit = 512
	}
	offset := 0
	count := 0
	for {
		buf := util.Builder{}
		buf.Grow(32)
		buf.Write([]byte(" limit "))
		buf.WriteInt(limit)
		buf.WriteByte(',')
		buf.WriteInt(offset)
		rows, err := t.db.QueryContext(ctx, findSql+buf.String())
		if err != nil {
			return fmt.Errorf("exec db_user.user_test range_all failed, offset:%d limit:%d %w", offset, limit, err)
		}
		defer rows.Close()
		count = 0
		for rows.Next() {
			data, err := scanUserTest(rows)
			if err != nil {
				return err
			}
			if !f(ctx, data) {
				return nil
			}
			count++
		}
		if count < limit {
			break
		}
		offset += limit
	}
	return nil
}

func (t *xUserTestOperation) RangeAllEx(ctx context.Context, where *UserTestWhereStmt, f func(ctx context.Context, row *dbop.UserTestEx) bool) error {
	var findSql = UserTestSQL_FindRow + where.String()
	limit := where.limit
	if limit == 0 {
		limit = 512
	}
	offset := 0
	count := 0
	for {
		buf := util.Builder{}
		buf.Grow(32)
		buf.Write([]byte(" limit "))
		buf.WriteInt(limit)
		buf.WriteByte(',')
		buf.WriteInt(offset)
		rows, err := t.db.QueryContext(ctx, findSql+buf.String())
		if err != nil {
			return fmt.Errorf("exec db_user.user_test range_all failed, offset:%d limit:%d %w", offset, limit, err)
		}
		defer rows.Close()
		count = 0
		for rows.Next() {
			data, err := scanUserTestEx(rows)
			if err != nil {
				return err
			}
			if !f(ctx, data) {
				return nil
			}
			count++
		}
		if count < limit {
			break
		}
		offset += limit
	}
	return nil
}

func (t *xUserTestOperation) AllData(ctx context.Context, where *UserTestWhereStmt) (datas []*dbop.UserTest, err error) {
	var findSql = UserTestSQL_Find + where.String()
	limit := where.limit
	if limit == 0 {
		limit = 512
	}
	offset := 0
	datas = make([]*dbop.UserTest, 0, limit)
	for {
		buf := util.Builder{}
		buf.Grow(32)
		buf.Write([]byte(" limit "))
		buf.WriteInt(limit)
		buf.WriteByte(',')
		buf.WriteInt(offset)
		rows, err := t.db.QueryContext(ctx, findSql+buf.String())
		if err != nil {
			return nil, fmt.Errorf("exec db_user.user_test all_data failed, offset:%d limit:%d %w", offset, limit, err)
		}
		defer rows.Close()

		for rows.Next() {
			data, err := scanUserTest(rows)
			if err != nil {
				return nil, err
			}
			datas = append(datas, data)
		}
		if len(datas) < offset+limit {
			break
		}
		offset += limit
	}
	return
}

func (t *xUserTestOperation) AllDataEx(ctx context.Context, where *UserTestWhereStmt) (datas []*dbop.UserTestEx, err error) {
	var findSql = UserTestSQL_FindRow + where.String()
	limit := where.limit
	if limit == 0 {
		limit = 512
	}
	offset := 0
	datas = make([]*dbop.UserTestEx, 0, limit)
	for {
		buf := util.Builder{}
		buf.Grow(32)
		buf.Write([]byte(" limit "))
		buf.WriteInt(limit)
		buf.WriteByte(',')
		buf.WriteInt(offset)
		rows, err := t.db.QueryContext(ctx, findSql+buf.String())
		if err != nil {
			return nil, fmt.Errorf("exec db_user.user_test all_data_ex failed, offset:%d limit:%d %w", offset, limit, err)
		}
		defer rows.Close()

		for rows.Next() {
			data, err := scanUserTestEx(rows)
			if err != nil {
				return nil, err
			}
			datas = append(datas, data)
		}
		if len(datas) < offset+limit {
			break
		}
		offset += limit
	}
	return
}

func (t *xUserTestOperation) DB() *sqlx.DB {
	return t.db
}

////////////////////////////////////////////////////////////////////////////////
// where stmt

type UserTestWhereStmt struct {
	buf           util.Builder
	limit, offset int
}

func (w *UserTestWhereStmt) Uid() *svc_db.IntSignedCondition[UserTestWhereStmt, int64] {
	return svc_db.NewIntSignedCondition[UserTestWhereStmt, int64](w, &w.buf, "uid")
}

func (w *UserTestWhereStmt) Xxx() *svc_db.IntSignedCondition[UserTestWhereStmt, int64] {
	return svc_db.NewIntSignedCondition[UserTestWhereStmt, int64](w, &w.buf, "xxx")
}

func (w *UserTestWhereStmt) State() *svc_db.IntSignedCondition[UserTestWhereStmt, int8] {
	return svc_db.NewIntSignedCondition[UserTestWhereStmt, int8](w, &w.buf, "state")
}

func (w *UserTestWhereStmt) Limit(limit, offset int) *UserTestWhereStmt {
	w.limit = limit
	w.offset = offset
	return w
}

func (w *UserTestWhereStmt) And() *UserTestWhereStmt {
	w.buf.Write([]byte(" and "))
	return w
}

func (w *UserTestWhereStmt) Or() *UserTestWhereStmt {
	w.buf.Write([]byte(" or "))
	return w
}

func (w *UserTestWhereStmt) Group(gf func(w *UserTestWhereStmt)) *UserTestWhereStmt {
	w.buf.WriteByte('(')
	gf(w)
	w.buf.WriteByte(')')
	return w
}

func (w *UserTestWhereStmt) Custom(f func(buf *util.Builder)) *UserTestWhereStmt {
	f(&w.buf)
	return w
}

func (w *UserTestWhereStmt) applyLimitAndOffset() {
	if w.limit == 0 && w.offset == 0 {
		return
	}
	w.buf.Write([]byte(" limit "))
	w.buf.WriteInt(w.limit)
	w.buf.WriteByte(',')
	w.buf.WriteInt(w.offset)
}

func (w *UserTestWhereStmt) String() string {
	return w.buf.String()
}

////////////////////////////////////////////////////////////////////////////////
// scan interface

func scanUserTest(rows *sql.Rows) (data *dbop.UserTest, err error) {
	var values [3]sql.RawBytes
	err = rows.Scan(&values[0], &values[1], &values[2])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_test scan failed, %w", err)
	}

	data = &dbop.UserTest{}
	data.Uid, err = UserTestUidUnamrshal(values[0])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_test scan uid failed, %w", err)
	}
	data.Xxx, err = UserTestXxxUnamrshal(values[1])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_test scan xxx failed, %w", err)
	}
	data.State, err = UserTestStateUnamrshal(values[2])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_test scan state failed, %w", err)
	}
	return data, nil
}

func scanUserTestEx(rows *sql.Rows) (data *dbop.UserTestEx, err error) {
	var values [3 + 2]sql.RawBytes
	err = rows.Scan(&values[0], &values[1], &values[2], &values[3], &values[3+1])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_test scan_ex failed, %w", err)
	}

	data = &dbop.UserTestEx{}
	data.Uid, err = UserTestUidUnamrshal(values[0])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_test scan_ex uid failed, %w", err)
	}
	data.Xxx, err = UserTestXxxUnamrshal(values[1])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_test scan_ex xxx failed, %w", err)
	}
	data.State, err = UserTestStateUnamrshal(values[2])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_test scan_ex state failed, %w", err)
	}
	data.ModifyStamp, err = svc_db.RawToStampInt64(values[3])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_test scan_ex modify_stamp failed, %w", err)
	}
	data.CreateStamp, err = svc_db.RawToStampInt64(values[3+1])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_test scan_ex create_stamp failed, %w", err)
	}
	return data, nil
}

////////////////////////////////////////////////////////////////////////////////
// named sql

type UserTestSQLWriter struct {
	buf util.Builder
}

func (x *UserTestSQLWriter) Select() *UserTestNamedSelect {
	x.buf.Write([]byte("select "))
	var v int
	return &UserTestNamedSelect{
		buf: &x.buf,
		n:   &v,
	}
}

func (x *UserTestSQLWriter) Update() *UserTestNamedUpdate {
	x.buf.Write([]byte("update user_test set "))
	var v int
	return &UserTestNamedUpdate{
		buf: &x.buf,
		n:   &v,
	}
}

func (x *UserTestSQLWriter) Insert() *UserTestNamedInsert {
	return &UserTestNamedInsert{
		buf: &x.buf,
	}
}

func (x *UserTestSQLWriter) Delete() *UserTestNamedWhere {
	x.buf.Write([]byte("delete user_test where "))
	return &UserTestNamedWhere{
		buf: &x.buf,
	}
}

type UserTestNamedInsert struct {
	buf          *util.Builder
	list, values []string
}

func (x *UserTestNamedInsert) Uid() *UserTestNamedInsert {
	x.list = append(x.list, "`uid`")
	x.values = append(x.values, ":uid")
	return x
}

func (x *UserTestNamedInsert) Xxx() *UserTestNamedInsert {
	x.list = append(x.list, "`xxx`")
	x.values = append(x.values, ":xxx")
	return x
}

func (x *UserTestNamedInsert) State() *UserTestNamedInsert {
	x.list = append(x.list, "`state`")
	x.values = append(x.values, ":state")
	return x
}

func (x *UserTestNamedInsert) ToSQL() string {
	x.buf.Write([]byte("insert user_test("))
	x.buf.WriteString(strings.Join(x.list, ","))
	x.buf.Write([]byte(") values("))
	x.buf.WriteString(strings.Join(x.values, ","))
	x.buf.Write([]byte(")"))
	return x.buf.String()
}

func (x *UserTestNamedInsert) ValuesToSQL() string {
	x.buf.Write([]byte(",values("))
	x.buf.WriteString(strings.Join(x.values, ","))
	x.buf.Write([]byte(")"))
	return x.buf.String()
}

type UserTestNamedUpdate struct {
	buf    *util.Builder
	n      *int
	values *bool
}

func (x *UserTestNamedUpdate) Uid() *UserTestNamedUpdate {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	if x.values != nil && *x.values {
		x.buf.Write([]byte("`uid`=values(`uid`)"))
	}
	x.buf.Write([]byte("`uid`=:uid"))
	*x.n++
	return x
}

func (x *UserTestNamedUpdate) Xxx() *UserTestNamedUpdate {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	if x.values != nil && *x.values {
		x.buf.Write([]byte("`xxx`=values(`xxx`)"))
	}
	x.buf.Write([]byte("`xxx`=:xxx"))
	*x.n++
	return x
}

func (x *UserTestNamedUpdate) State() *UserTestNamedUpdate {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	if x.values != nil && *x.values {
		x.buf.Write([]byte("`state`=values(`state`)"))
	}
	x.buf.Write([]byte("`state`=:state"))
	*x.n++
	return x
}

func (x *UserTestNamedUpdate) Where() *UserTestNamedWhere {
	if x.values != nil {
		panic("invalid where")
	}
	x.buf.Write([]byte(" where "))
	return &UserTestNamedWhere{
		buf: x.buf,
	}
}

func (x *UserTestNamedUpdate) ToSQL() string {
	return x.buf.String()
}

type UserTestNamedSelect struct {
	buf *util.Builder
	n   *int
}

func (x *UserTestNamedSelect) Uid() *UserTestNamedSelect {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`uid`"))
	*x.n++
	return x
}

func (x *UserTestNamedSelect) Xxx() *UserTestNamedSelect {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`xxx`"))
	*x.n++
	return x
}

func (x *UserTestNamedSelect) State() *UserTestNamedSelect {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`state`"))
	*x.n++
	return x
}

func (x *UserTestNamedSelect) Where() *UserTestNamedWhere {
	x.buf.Write([]byte(" from user_test where "))
	return &UserTestNamedWhere{
		buf: x.buf,
	}
}

func (x *UserTestNamedSelect) ToSQL() string {
	x.buf.Write([]byte(" from user_test"))
	return x.buf.String()
}

type UserTestNamedWhere struct {
	buf *util.Builder
}

func (x *UserTestNamedWhere) Uid() *UserTestNamedWhere {
	x.buf.Write([]byte("`uid` = :uid"))
	return x
}

func (x *UserTestNamedWhere) Xxx() *UserTestNamedWhere {
	x.buf.Write([]byte("`xxx` = :xxx"))
	return x
}

func (x *UserTestNamedWhere) State() *UserTestNamedWhere {
	x.buf.Write([]byte("`state` = :state"))
	return x
}

func (x *UserTestNamedWhere) Limit(limit, offset int) *UserTestNamedWhere {
	x.buf.Write([]byte(" limit "))
	x.buf.WriteInt(limit)
	x.buf.WriteByte(',')
	x.buf.WriteInt(offset)
	return x
}

func (x *UserTestNamedWhere) And() *UserTestNamedWhere {
	x.buf.Write([]byte(" and "))
	return x
}

func (x *UserTestNamedWhere) Or() *UserTestNamedWhere {
	x.buf.Write([]byte(" or "))
	return x
}

func (x *UserTestNamedWhere) Group(gf func(w *UserTestNamedWhere)) *UserTestNamedWhere {
	x.buf.WriteByte('(')
	gf(x)
	x.buf.WriteByte(')')
	return x
}

func (x *UserTestNamedWhere) Custom(f func(buf *util.Builder)) *UserTestNamedWhere {
	f(x.buf)
	return x
}

func (x *UserTestNamedWhere) OnDuplicateKeyUpdate() *UserTestNamedUpdate {
	x.buf.Write([]byte(" on duplicate key update "))
	var v int
	values := false
	return &UserTestNamedUpdate{
		buf:    x.buf,
		n:      &v,
		values: &values,
	}
}

func (x *UserTestNamedWhere) OnDuplicateKeyUpdateValues() *UserTestNamedUpdate {
	x.buf.Write([]byte(" on duplicate key update "))
	var v int
	values := true
	return &UserTestNamedUpdate{
		buf:    x.buf,
		n:      &v,
		values: &values,
	}
}

func (x *UserTestNamedWhere) ToSQL() string {
	return x.buf.String()
}

func (x *UserTestNamedWhere) OrderBy() *UserTestNamedOrderBy {
	x.buf.Write([]byte(" order by "))
	var v int
	return &UserTestNamedOrderBy{
		buf: x.buf,
		n:   &v,
	}
}

type UserTestNamedOrderBy struct {
	buf *util.Builder
	n   *int
}

func (x *UserTestNamedOrderBy) Uid() *UserTestNamedOrderAsc {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`uid`"))
	*x.n++
	return &UserTestNamedOrderAsc{
		buf: x.buf,
		n:   x.n,
	}
}

func (x *UserTestNamedOrderBy) Xxx() *UserTestNamedOrderAsc {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`xxx`"))
	*x.n++
	return &UserTestNamedOrderAsc{
		buf: x.buf,
		n:   x.n,
	}
}

func (x *UserTestNamedOrderBy) State() *UserTestNamedOrderAsc {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`state`"))
	*x.n++
	return &UserTestNamedOrderAsc{
		buf: x.buf,
		n:   x.n,
	}
}

func (x *UserTestNamedOrderBy) Limit(limit, offset int) *UserTestNamedOrderBy {
	x.buf.Write([]byte(" limit "))
	x.buf.WriteInt(limit)
	x.buf.WriteByte(',')
	x.buf.WriteInt(offset)
	return x
}

func (x *UserTestNamedOrderBy) ToSQL() string {
	return x.buf.String()
}

type UserTestNamedOrderAsc struct {
	buf *util.Builder
	n   *int
}

func (x *UserTestNamedOrderAsc) Asc() *UserTestNamedOrderBy {
	x.buf.Write([]byte(" asc "))
	return &UserTestNamedOrderBy{
		buf: x.buf,
		n:   x.n,
	}
}

func (x *UserTestNamedOrderAsc) Desc() *UserTestNamedOrderBy {
	x.buf.Write([]byte(" desc "))
	return &UserTestNamedOrderBy{
		buf: x.buf,
		n:   x.n,
	}
}

func (x *UserTestNamedOrderAsc) ToSQL() string {
	return x.buf.String()
}
