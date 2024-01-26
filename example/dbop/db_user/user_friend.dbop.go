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

type UserFriendKey struct {
	Uid int64
	Fid int64
}

type UserFriendOperation interface {
	Insert(ctx context.Context, data *dbop.UserFriend) (res sql.Result, err error)
	InsertMany(ctx context.Context, datas []*dbop.UserFriend) (res sql.Result, err error)

	Update(ctx context.Context, data *dbop.UserFriend) (res sql.Result, err error)
	Upsert(ctx context.Context, data *dbop.UserFriend) (res sql.Result, err error)
	UpsertMany(ctx context.Context, datas []*dbop.UserFriend) (res sql.Result, err error)

	Find(ctx context.Context, uid int64, fid int64) (data *dbop.UserFriend, err error)
	FindEx(ctx context.Context, uid int64, fid int64) (data *dbop.UserFriendEx, err error)
	Delete(ctx context.Context, uid int64, fid int64) (res sql.Result, err error)

	FindByKey(ctx context.Context, id UserFriendKey) (data *dbop.UserFriend, err error)
	FindExByKey(ctx context.Context, id UserFriendKey) (data *dbop.UserFriendEx, err error)
	DeleteByKey(ctx context.Context, id UserFriendKey) (res sql.Result, err error)

	FindByKeyArray(ctx context.Context, ids []UserFriendKey) (datas []*dbop.UserFriend, err error)
	FindExByKeyArray(ctx context.Context, ids []UserFriendKey) (datas []*dbop.UserFriendEx, err error)
	DeleteByKeyArray(ctx context.Context, ids []UserFriendKey) (res sql.Result, err error)

	FindByIndexUid(ctx context.Context, uid int64, limit, offset int) (datas []*dbop.UserFriend, err error)
	FindExByIndexUid(ctx context.Context, uid int64, limit, offset int) (datas []*dbop.UserFriendEx, err error)
	CountByIndexUid(ctx context.Context, uid int64) (count int, err error)
	DeleteByIndexUid(ctx context.Context, uid int64) (res sql.Result, err error)

	Where(bufSize int) *UserFriendWhereStmt
	Select(ctx context.Context, where *UserFriendWhereStmt) (datas []*dbop.UserFriend, err error)
	SelectEx(ctx context.Context, where *UserFriendWhereStmt) (datas []*dbop.UserFriendEx, err error)

	DeleteMany(ctx context.Context, where *UserFriendWhereStmt) (res sql.Result, err error)

	RangeAll(ctx context.Context, where *UserFriendWhereStmt, f func(ctx context.Context, row *dbop.UserFriend) bool) error
	RangeAllEx(ctx context.Context, where *UserFriendWhereStmt, f func(ctx context.Context, row *dbop.UserFriendEx) bool) error
	AllData(ctx context.Context, where *UserFriendWhereStmt) (datas []*dbop.UserFriend, err error)
	AllDataEx(ctx context.Context, where *UserFriendWhereStmt) (datas []*dbop.UserFriendEx, err error)

	// use for custom named sql
	DB() *sqlx.DB
}

var (
	UserFriendUidUnamrshal   = svc_db.RawToInt64
	UserFriendFidUnamrshal   = svc_db.RawToInt64
	UserFriendStateUnamrshal = svc_db.RawToInt8
)

var (
	globalUserFriendOP atomic.Pointer[xUserFriendOperation]
)

func init() {
	svc_db.RegisterDB("mysql", "db_user", "user_friend", func(db *sqlx.DB) error {
		// sync table columns
		err := svc_db.SyncTableColumns(context.Background(), db, "user_friend", UserFriendSQL_Create, UserFriendSQL_TableColumns)
		if err != nil {
			return fmt.Errorf("swap db_user.user_friend pointer, sync columns failed, %w", err)
		}
		// sync table index
		err = svc_db.SyncTableIndex(context.Background(), db, "user_friend", UserFriendSQL_TableIndex)
		if err != nil {
			return fmt.Errorf("swap db_user.user_friend pointer, sync index failed, %w", err)
		}
		//
		tableOP, err := NewUserFriendOperation(db)
		if err != nil {
			return fmt.Errorf("swap db_user.user_friend pointer, new table operation failed, %w", err)
		}

		globalUserFriendOP.Store(tableOP)
		return nil
	})
}

var UserFriendOP = func() UserFriendOperation {
	return globalUserFriendOP.Load()
}

func UserFriendNamedSQL(bufSize int) *UserFriendSQLWriter {
	sql := &UserFriendSQLWriter{}
	sql.buf.Grow(bufSize)
	return sql
}

func UserFriendToPrimaryKeys(rows []*dbop.UserFriend) (ids []UserFriendKey) {
	ids = make([]UserFriendKey, 0, len(rows))
	for _, v := range rows {
		ids = append(ids, UserFriendKey{
			Uid: v.Uid,
			Fid: v.Fid,
		})
	}
	return
}

func UserFriendExToPrimaryKeysEx(rows []*dbop.UserFriendEx) (ids []UserFriendKey) {
	ids = make([]UserFriendKey, 0, len(rows))
	for _, v := range rows {
		ids = append(ids, UserFriendKey{
			Uid: v.Uid,
			Fid: v.Fid,
		})
	}
	return
}

////////////////////////////////////////////////////////////////////////////////
// sql statement

const (
	UserFriendSQL_Insert        = "insert user_friend(`uid`,`fid`,`state`) values(?,?,?)"
	UserFriendSQL_InsertValues  = ",values(?,?,?)"
	UserFriendSQL_InsertValues2 = ",values(?,?,?)"
	UserFriendSQL_Where1        = " where (`uid`=?,`fid`=?)"
	UserFriendSQL_Where2        = " or (`uid`=?,`fid`=?)"
	UserFriendSQL_Upsert        = "insert user_friend(`uid`,`fid`,`state`) values(?,?,?)"
	UserFriendSQL_UpsertUpdate  = " on duplicate key update `uid`=values(`uid`),`fid`=values(`fid`),`state`=values(`state`)"
	UserFriendSQL_Update        = "update user_friend set `state`=? where `uid`=?,`fid`=?"
	UserFriendSQL_Delete        = "delete from user_friend"
	UserFriendSQL_Find          = "select user_friend(`uid`,`fid`,`state`) from user_friend"
	UserFriendSQL_FindRow       = "select user_friend(`uid`,`fid`,`state`,`modify_stamp`,`create_stamp`) from user_friend"
	UserFriendSQL_Count         = "select count(id) from user_friend"
	UserFriendSQL_Create        = "create table user_friend (" +
		"`uid` bigint not null default 0," +
		"`fid` bigint not null default 0," +
		"`state` tinyint not null default 0," +
		"`modify_stamp` timestamp default current_timestamp on update current_timestamp," +
		"`create_stamp` timestamp default current_timestamp" + "," +
		"PRIMARY KEY ( `uid`,`fid`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;"
)

var (
	UserFriendSQL_TableColumns = map[string]string{
		"uid":   "alter table user_friend add `uid` bigint not null default 0;",
		"fid":   "alter table user_friend add `fid` bigint not null default 0;",
		"state": "alter table user_friend add `state` tinyint not null default 0;",
	}
	UserFriendSQL_TableIndex = map[string]string{
		"user_friend_uid": "create index user_friend_uid on user_friend(`uid`)",
	}
)

////////////////////////////////////////////////////////////////////////////////
// TblTestOperation impl

type xUserFriendOperation struct {
	db           *sqlx.DB
	insert       *sql.Stmt
	update       *sql.Stmt
	upsert       *sql.Stmt
	delete       *sql.Stmt
	find         *sql.Stmt
	findRow      *sql.Stmt
	idxUidFind   *sql.Stmt
	idxUidFindEx *sql.Stmt
	idxUidCount  *sql.Stmt
	idxUidDelete *sql.Stmt
}

func NewUserFriendOperation(db *sqlx.DB) (_ *xUserFriendOperation, err error) {
	t := &xUserFriendOperation{
		db: db,
	}

	t.insert, err = db.Prepare(UserFriendSQL_Insert)
	if err != nil {
		return nil, fmt.Errorf("prepare db_user.user_friend insert failed,%w", err)
	}
	t.update, err = db.Prepare(UserFriendSQL_Update)
	if err != nil {
		return nil, fmt.Errorf("prepare db_user.user_friend update failed,%w", err)
	}
	t.upsert, err = db.Prepare(UserFriendSQL_Upsert + UserFriendSQL_UpsertUpdate)
	if err != nil {
		return nil, fmt.Errorf("prepare db_user.user_friend upsert failed,%w", err)
	}
	t.delete, err = db.Prepare(UserFriendSQL_Delete + " where `uid`=?,`fid`=?")
	if err != nil {
		return nil, fmt.Errorf("prepare db_user.user_friend delete failed,%w", err)
	}
	t.find, err = db.Prepare(UserFriendSQL_Find + " where `uid`=?,`fid`=?")
	if err != nil {
		return nil, fmt.Errorf("prepare db_user.user_friend find failed,%w", err)
	}
	t.findRow, err = db.Prepare(UserFriendSQL_FindRow + " where `uid`=?,`fid`=?")
	if err != nil {
		return nil, fmt.Errorf("prepare db_user.user_friend findex failed,%w", err)
	}

	return t, nil
}

func (t *xUserFriendOperation) Insert(ctx context.Context, data *dbop.UserFriend) (res sql.Result, err error) {
	res, err = t.insert.ExecContext(ctx, data.Uid, data.Fid, data.State)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend insert failed,%w", err)
	}

	return
}

func (t *xUserFriendOperation) InsertMany(ctx context.Context, datas []*dbop.UserFriend) (res sql.Result, err error) {
	switch len(datas) {
	case 0:
		return svc_db.EmptyResult{}, nil
	case 1:
		return t.Insert(ctx, datas[0])
	}

	buf := util.Builder{}
	buf.Grow(len(UserFriendSQL_Insert) + (len(datas)-1)*len(UserFriendSQL_InsertValues))
	buf.Write([]byte(UserFriendSQL_Insert))
	for i := 0; i < len(datas)-1; i++ {
		buf.Write([]byte(UserFriendSQL_InsertValues))
	}
	args := make([]any, 0, len(datas)*3)
	for i := 0; i < len(datas); i++ {
		data := datas[i]
		args = append(args, data.Uid, data.Fid, data.State)
	}
	res, err = t.db.DB.ExecContext(ctx, buf.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend insert_many failed,%w", err)
	}
	return
}

func (t *xUserFriendOperation) Update(ctx context.Context, data *dbop.UserFriend) (res sql.Result, err error) {
	res, err = t.update.ExecContext(ctx, data.State, data.Uid, data.Fid)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend update failed,%w", err)
	}

	return
}

func (t *xUserFriendOperation) Upsert(ctx context.Context, data *dbop.UserFriend) (res sql.Result, err error) {

	res, err = t.upsert.ExecContext(ctx, data.Uid, data.Fid, data.State)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend upsert failed,%w", err)
	}

	return
}

func (t *xUserFriendOperation) UpsertMany(ctx context.Context, datas []*dbop.UserFriend) (res sql.Result, err error) {
	switch len(datas) {
	case 0:
		return svc_db.EmptyResult{}, nil
	case 1:
		return t.Upsert(ctx, datas[0])
	}

	buf := util.Builder{}
	buf.Grow(len(UserFriendSQL_Upsert) + (len(datas)-1)*len(UserFriendSQL_InsertValues2) + len(UserFriendSQL_UpsertUpdate))
	buf.Write([]byte(UserFriendSQL_Upsert))
	for i := 0; i < len(datas)-1; i++ {
		buf.Write([]byte(UserFriendSQL_InsertValues2))
	}
	buf.Write([]byte(UserFriendSQL_UpsertUpdate))
	args := make([]any, 0, len(datas)*3)
	for i := 0; i < len(datas); i++ {
		data := datas[i]
		args = append(args, data.Uid, data.Fid, data.State)
	}
	res, err = t.db.DB.ExecContext(ctx, buf.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend upsert_many failed,%w", err)
	}
	return
}

// find by primary key
func (t *xUserFriendOperation) Find(ctx context.Context, uid int64, fid int64) (data *dbop.UserFriend, err error) {
	rows, err := t.find.QueryContext(ctx, uid, fid)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend find failed,%w", err)
	}
	defer rows.Close()
	for rows.Next() {
		return scanUserFriend(rows)
	}
	return
}

func (t *xUserFriendOperation) FindEx(ctx context.Context, uid int64, fid int64) (data *dbop.UserFriendEx, err error) {
	rows, err := t.findRow.QueryContext(ctx, uid, fid)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend findex failed,%w", err)
	}
	defer rows.Close()
	for rows.Next() {
		return scanUserFriendEx(rows)
	}
	return
}

func (t *xUserFriendOperation) Delete(ctx context.Context, uid int64, fid int64) (res sql.Result, err error) {
	res, err = t.delete.ExecContext(ctx, uid, fid)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend delete failed,%w", err)
	}

	return
}

// find by primary key
func (t *xUserFriendOperation) FindByKey(ctx context.Context, id UserFriendKey) (data *dbop.UserFriend, err error) {
	rows, err := t.find.QueryContext(ctx, id.Uid, id.Fid)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend find_by_key failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		return scanUserFriend(rows)
	}
	return
}

func (t *xUserFriendOperation) FindExByKey(ctx context.Context, id UserFriendKey) (data *dbop.UserFriendEx, err error) {
	rows, err := t.findRow.QueryContext(ctx, id.Uid, id.Fid)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend findex_by_key failed,%w", err)
	}
	defer rows.Close()
	data = &dbop.UserFriendEx{}
	for rows.Next() {
		return scanUserFriendEx(rows)
	}
	return
}

func (t *xUserFriendOperation) DeleteByKey(ctx context.Context, id UserFriendKey) (res sql.Result, err error) {
	res, err = t.delete.ExecContext(ctx, id.Uid, id.Fid)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend del_by_key failed,%w", err)
	}

	return
}

// find by primary key
func (t *xUserFriendOperation) FindByKeyArray(ctx context.Context, ids []UserFriendKey) (datas []*dbop.UserFriend, err error) {
	switch len(ids) {
	case 0:
		return nil, nil
	case 1:
		data, err := t.FindByKey(ctx, ids[0])
		if err != nil {
			return nil, err
		}
		return []*dbop.UserFriend{data}, nil
	}
	buf := util.Builder{}
	buf.Grow(len(UserFriendSQL_Find) + len(UserFriendSQL_Where1) + (len(ids)-1)*len(UserFriendSQL_Where2))
	buf.Write([]byte(UserFriendSQL_Find))
	buf.Write([]byte(UserFriendSQL_Where1))
	for i := 0; i < len(ids)-1; i++ {
		buf.Write([]byte(UserFriendSQL_Where2))
	}

	args := make([]any, 0, len(ids)*2)
	for i := 0; i < len(ids); i++ {
		args = append(args, ids[i].Uid, ids[i].Fid)
	}
	rows, err := t.db.DB.QueryContext(ctx, buf.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend find_by_key_array failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		data, err := scanUserFriend(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return
}

func (t *xUserFriendOperation) FindExByKeyArray(ctx context.Context, ids []UserFriendKey) (datas []*dbop.UserFriendEx, err error) {
	switch len(ids) {
	case 0:
		return nil, nil
	case 1:
		data, err := t.FindExByKey(ctx, ids[0])
		if err != nil {
			return nil, err
		}
		return []*dbop.UserFriendEx{data}, nil
	}
	buf := util.Builder{}
	buf.Grow(len(UserFriendSQL_Find) + len(UserFriendSQL_Where1) + (len(ids)-1)*len(UserFriendSQL_Where2))
	buf.Write([]byte(UserFriendSQL_Find))
	buf.Write([]byte(UserFriendSQL_Where1))
	for i := 0; i < len(ids)-1; i++ {
		buf.Write([]byte(UserFriendSQL_Where2))
	}

	args := make([]any, 0, len(ids)*2)
	for i := 0; i < len(ids); i++ {
		args = append(args, ids[i].Uid, ids[i].Fid)
	}
	rows, err := t.db.DB.QueryContext(ctx, buf.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend findex_by_key_array failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		data, err := scanUserFriendEx(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return
}

func (t *xUserFriendOperation) DeleteByKeyArray(ctx context.Context, ids []UserFriendKey) (res sql.Result, err error) {
	switch len(ids) {
	case 0:
		return svc_db.EmptyResult{}, nil
	case 1:
		return t.DeleteByKey(ctx, ids[0])
	}
	buf := util.Builder{}
	buf.Grow(len(UserFriendSQL_Delete) + len(UserFriendSQL_Where1) + (len(ids)-1)*len(UserFriendSQL_Where2))
	buf.Write([]byte(UserFriendSQL_Delete))
	buf.Write([]byte(UserFriendSQL_Where1))
	for i := 0; i < len(ids)-1; i++ {
		buf.Write([]byte(UserFriendSQL_Where2))
	}

	args := make([]any, 0, len(ids)*2)
	for i := 0; i < len(ids); i++ {
		args = append(args, ids[i].Uid, ids[i].Fid)
	}
	res, err = t.db.DB.ExecContext(ctx, buf.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend del_by_key_array failed,%w", err)
	}
	return
}

func (t *xUserFriendOperation) FindByIndexUid(ctx context.Context, uid int64, limit, offset int) (datas []*dbop.UserFriend, err error) {
	if t.idxUidFind == nil {
		t.idxUidFind, err = t.db.PrepareContext(ctx, UserFriendSQL_Find+"where `uid`=? limit ?,?")
		if err != nil {
			return nil, fmt.Errorf("prepare db_user.user_friend find_by_index_uid failed,%w", err)
		}
	}
	rows, err := t.idxUidFind.QueryContext(ctx, uid, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend find_by_index_uid failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		data, err := scanUserFriend(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return
}
func (t *xUserFriendOperation) FindExByIndexUid(ctx context.Context, uid int64, limit, offset int) (datas []*dbop.UserFriendEx, err error) {
	if t.idxUidFindEx == nil {
		t.idxUidFindEx, err = t.db.PrepareContext(ctx, UserFriendSQL_FindRow+"where `uid`=? limit ?,?")
		if err != nil {
			return nil, fmt.Errorf("prepare db_user.user_friend findex_by_index_uid failed,%w", err)
		}
	}
	rows, err := t.idxUidFindEx.QueryContext(ctx, uid, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend findex_by_index_uid failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		data, err := scanUserFriendEx(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return
}
func (t *xUserFriendOperation) CountByIndexUid(ctx context.Context, uid int64) (count int, err error) {
	if t.idxUidCount == nil {
		t.idxUidCount, err = t.db.PrepareContext(ctx, UserFriendSQL_Count+"where `uid`=?")
		if err != nil {
			return 0, fmt.Errorf("prepare db_user.user_friend count_by_index_uid failed,%w", err)
		}
	}
	err = t.idxUidCount.QueryRowContext(ctx, uid).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("exec db_user.user_friend count_by_index_uid failed,%w", err)
	}
	return
}

func (t *xUserFriendOperation) DeleteByIndexUid(ctx context.Context, uid int64) (res sql.Result, err error) {
	if t.idxUidDelete == nil {
		t.idxUidDelete, err = t.db.PrepareContext(ctx, UserFriendSQL_Count+"where `uid`=?")
		if err != nil {
			return nil, fmt.Errorf("prepare db_user.user_friend delete_by_index_uid failed,%w", err)
		}
	}
	res, err = t.idxUidDelete.ExecContext(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend delete_by_index_uid failed,%w", err)
	}
	return
}

func (t *xUserFriendOperation) Where(bufSize int) *UserFriendWhereStmt {
	w := &UserFriendWhereStmt{}
	w.buf.Grow(bufSize)
	return w
}

func (t *xUserFriendOperation) Select(ctx context.Context, where *UserFriendWhereStmt) (datas []*dbop.UserFriend, err error) {
	where.applyLimitAndOffset()
	var findSql = UserFriendSQL_Find + where.String()
	rows, err := t.db.QueryContext(ctx, findSql)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend select failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {

		data, err := scanUserFriend(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return
}
func (t *xUserFriendOperation) SelectEx(ctx context.Context, where *UserFriendWhereStmt) (datas []*dbop.UserFriendEx, err error) {
	where.applyLimitAndOffset()
	var findSql = UserFriendSQL_FindRow + where.String()
	rows, err := t.db.QueryContext(ctx, findSql)
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend selectex failed,%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		data, err := scanUserFriendEx(rows)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return
}

func (t *xUserFriendOperation) DeleteMany(ctx context.Context, where *UserFriendWhereStmt) (res sql.Result, err error) {
	w := where.String()
	buf := util.Builder{}
	buf.Grow(len(UserFriendSQL_Delete) + len(w))
	buf.Write([]byte(UserFriendSQL_Delete))
	buf.WriteString(w)
	res, err = t.db.ExecContext(ctx, buf.String())
	if err != nil {
		return nil, fmt.Errorf("exec db_user.user_friend delete_many failed,%w", err)
	}

	return
}

func (t *xUserFriendOperation) RangeAll(ctx context.Context, where *UserFriendWhereStmt, f func(ctx context.Context, row *dbop.UserFriend) bool) error {
	var findSql = UserFriendSQL_Find + where.String()
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
			return fmt.Errorf("exec db_user.user_friend range_all failed, offset:%d limit:%d %w", offset, limit, err)
		}
		defer rows.Close()
		count = 0
		for rows.Next() {
			data, err := scanUserFriend(rows)
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

func (t *xUserFriendOperation) RangeAllEx(ctx context.Context, where *UserFriendWhereStmt, f func(ctx context.Context, row *dbop.UserFriendEx) bool) error {
	var findSql = UserFriendSQL_FindRow + where.String()
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
			return fmt.Errorf("exec db_user.user_friend range_all failed, offset:%d limit:%d %w", offset, limit, err)
		}
		defer rows.Close()
		count = 0
		for rows.Next() {
			data, err := scanUserFriendEx(rows)
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

func (t *xUserFriendOperation) AllData(ctx context.Context, where *UserFriendWhereStmt) (datas []*dbop.UserFriend, err error) {
	var findSql = UserFriendSQL_Find + where.String()
	limit := where.limit
	if limit == 0 {
		limit = 512
	}
	offset := 0
	datas = make([]*dbop.UserFriend, 0, limit)
	for {
		buf := util.Builder{}
		buf.Grow(32)
		buf.Write([]byte(" limit "))
		buf.WriteInt(limit)
		buf.WriteByte(',')
		buf.WriteInt(offset)
		rows, err := t.db.QueryContext(ctx, findSql+buf.String())
		if err != nil {
			return nil, fmt.Errorf("exec db_user.user_friend all_data failed, offset:%d limit:%d %w", offset, limit, err)
		}
		defer rows.Close()

		for rows.Next() {
			data, err := scanUserFriend(rows)
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

func (t *xUserFriendOperation) AllDataEx(ctx context.Context, where *UserFriendWhereStmt) (datas []*dbop.UserFriendEx, err error) {
	var findSql = UserFriendSQL_FindRow + where.String()
	limit := where.limit
	if limit == 0 {
		limit = 512
	}
	offset := 0
	datas = make([]*dbop.UserFriendEx, 0, limit)
	for {
		buf := util.Builder{}
		buf.Grow(32)
		buf.Write([]byte(" limit "))
		buf.WriteInt(limit)
		buf.WriteByte(',')
		buf.WriteInt(offset)
		rows, err := t.db.QueryContext(ctx, findSql+buf.String())
		if err != nil {
			return nil, fmt.Errorf("exec db_user.user_friend all_data_ex failed, offset:%d limit:%d %w", offset, limit, err)
		}
		defer rows.Close()

		for rows.Next() {
			data, err := scanUserFriendEx(rows)
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

func (t *xUserFriendOperation) DB() *sqlx.DB {
	return t.db
}

////////////////////////////////////////////////////////////////////////////////
// where stmt

type UserFriendWhereStmt struct {
	buf           util.Builder
	limit, offset int
}

func (w *UserFriendWhereStmt) Uid() *svc_db.IntSignedCondition[UserFriendWhereStmt, int64] {
	return svc_db.NewIntSignedCondition[UserFriendWhereStmt, int64](w, &w.buf, "uid")
}

func (w *UserFriendWhereStmt) Fid() *svc_db.IntSignedCondition[UserFriendWhereStmt, int64] {
	return svc_db.NewIntSignedCondition[UserFriendWhereStmt, int64](w, &w.buf, "fid")
}

func (w *UserFriendWhereStmt) State() *svc_db.IntSignedCondition[UserFriendWhereStmt, int8] {
	return svc_db.NewIntSignedCondition[UserFriendWhereStmt, int8](w, &w.buf, "state")
}

func (w *UserFriendWhereStmt) Limit(limit, offset int) *UserFriendWhereStmt {
	w.limit = limit
	w.offset = offset
	return w
}

func (w *UserFriendWhereStmt) And() *UserFriendWhereStmt {
	w.buf.Write([]byte(" and "))
	return w
}

func (w *UserFriendWhereStmt) Or() *UserFriendWhereStmt {
	w.buf.Write([]byte(" or "))
	return w
}

func (w *UserFriendWhereStmt) Group(gf func(w *UserFriendWhereStmt)) *UserFriendWhereStmt {
	w.buf.WriteByte('(')
	gf(w)
	w.buf.WriteByte(')')
	return w
}

func (w *UserFriendWhereStmt) Custom(f func(buf *util.Builder)) *UserFriendWhereStmt {
	f(&w.buf)
	return w
}

func (w *UserFriendWhereStmt) applyLimitAndOffset() {
	if w.limit == 0 && w.offset == 0 {
		return
	}
	w.buf.Write([]byte(" limit "))
	w.buf.WriteInt(w.limit)
	w.buf.WriteByte(',')
	w.buf.WriteInt(w.offset)
}

func (w *UserFriendWhereStmt) String() string {
	return w.buf.String()
}

////////////////////////////////////////////////////////////////////////////////
// scan interface

func scanUserFriend(rows *sql.Rows) (data *dbop.UserFriend, err error) {
	var values [3]sql.RawBytes
	err = rows.Scan(&values[0], &values[1], &values[2])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_friend scan failed, %w", err)
	}

	data = &dbop.UserFriend{}
	data.Uid, err = UserFriendUidUnamrshal(values[0])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_friend scan uid failed, %w", err)
	}
	data.Fid, err = UserFriendFidUnamrshal(values[1])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_friend scan fid failed, %w", err)
	}
	data.State, err = UserFriendStateUnamrshal(values[2])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_friend scan state failed, %w", err)
	}
	return data, nil
}

func scanUserFriendEx(rows *sql.Rows) (data *dbop.UserFriendEx, err error) {
	var values [3 + 2]sql.RawBytes
	err = rows.Scan(&values[0], &values[1], &values[2], &values[3], &values[3+1])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_friend scan_ex failed, %w", err)
	}

	data = &dbop.UserFriendEx{}
	data.Uid, err = UserFriendUidUnamrshal(values[0])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_friend scan_ex uid failed, %w", err)
	}
	data.Fid, err = UserFriendFidUnamrshal(values[1])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_friend scan_ex fid failed, %w", err)
	}
	data.State, err = UserFriendStateUnamrshal(values[2])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_friend scan_ex state failed, %w", err)
	}
	data.ModifyStamp, err = svc_db.RawToStampInt64(values[3])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_friend scan_ex modify_stamp failed, %w", err)
	}
	data.CreateStamp, err = svc_db.RawToStampInt64(values[3+1])
	if err != nil {
		return nil, fmt.Errorf("unmarshal db_user.user_friend scan_ex create_stamp failed, %w", err)
	}
	return data, nil
}

////////////////////////////////////////////////////////////////////////////////
// named sql

type UserFriendSQLWriter struct {
	buf util.Builder
}

func (x *UserFriendSQLWriter) Select() *UserFriendNamedSelect {
	x.buf.Write([]byte("select "))
	var v int
	return &UserFriendNamedSelect{
		buf: &x.buf,
		n:   &v,
	}
}

func (x *UserFriendSQLWriter) Update() *UserFriendNamedUpdate {
	x.buf.Write([]byte("update user_friend set "))
	var v int
	return &UserFriendNamedUpdate{
		buf: &x.buf,
		n:   &v,
	}
}

func (x *UserFriendSQLWriter) Insert() *UserFriendNamedInsert {
	return &UserFriendNamedInsert{
		buf: &x.buf,
	}
}

func (x *UserFriendSQLWriter) Delete() *UserFriendNamedWhere {
	x.buf.Write([]byte("delete user_friend where "))
	return &UserFriendNamedWhere{
		buf: &x.buf,
	}
}

type UserFriendNamedInsert struct {
	buf          *util.Builder
	list, values []string
}

func (x *UserFriendNamedInsert) Uid() *UserFriendNamedInsert {
	x.list = append(x.list, "`uid`")
	x.values = append(x.values, ":uid")
	return x
}

func (x *UserFriendNamedInsert) Fid() *UserFriendNamedInsert {
	x.list = append(x.list, "`fid`")
	x.values = append(x.values, ":fid")
	return x
}

func (x *UserFriendNamedInsert) State() *UserFriendNamedInsert {
	x.list = append(x.list, "`state`")
	x.values = append(x.values, ":state")
	return x
}

func (x *UserFriendNamedInsert) ToSQL() string {
	x.buf.Write([]byte("insert user_friend("))
	x.buf.WriteString(strings.Join(x.list, ","))
	x.buf.Write([]byte(") values("))
	x.buf.WriteString(strings.Join(x.values, ","))
	x.buf.Write([]byte(")"))
	return x.buf.String()
}

func (x *UserFriendNamedInsert) ValuesToSQL() string {
	x.buf.Write([]byte(",values("))
	x.buf.WriteString(strings.Join(x.values, ","))
	x.buf.Write([]byte(")"))
	return x.buf.String()
}

type UserFriendNamedUpdate struct {
	buf    *util.Builder
	n      *int
	values *bool
}

func (x *UserFriendNamedUpdate) Uid() *UserFriendNamedUpdate {
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

func (x *UserFriendNamedUpdate) Fid() *UserFriendNamedUpdate {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	if x.values != nil && *x.values {
		x.buf.Write([]byte("`fid`=values(`fid`)"))
	}
	x.buf.Write([]byte("`fid`=:fid"))
	*x.n++
	return x
}

func (x *UserFriendNamedUpdate) State() *UserFriendNamedUpdate {
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

func (x *UserFriendNamedUpdate) Where() *UserFriendNamedWhere {
	if x.values != nil {
		panic("invalid where")
	}
	x.buf.Write([]byte(" where "))
	return &UserFriendNamedWhere{
		buf: x.buf,
	}
}

func (x *UserFriendNamedUpdate) ToSQL() string {
	return x.buf.String()
}

type UserFriendNamedSelect struct {
	buf *util.Builder
	n   *int
}

func (x *UserFriendNamedSelect) Uid() *UserFriendNamedSelect {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`uid`"))
	*x.n++
	return x
}

func (x *UserFriendNamedSelect) Fid() *UserFriendNamedSelect {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`fid`"))
	*x.n++
	return x
}

func (x *UserFriendNamedSelect) State() *UserFriendNamedSelect {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`state`"))
	*x.n++
	return x
}

func (x *UserFriendNamedSelect) Where() *UserFriendNamedWhere {
	x.buf.Write([]byte(" from user_friend where "))
	return &UserFriendNamedWhere{
		buf: x.buf,
	}
}

func (x *UserFriendNamedSelect) ToSQL() string {
	x.buf.Write([]byte(" from user_friend"))
	return x.buf.String()
}

type UserFriendNamedWhere struct {
	buf *util.Builder
}

func (x *UserFriendNamedWhere) Uid() *UserFriendNamedWhere {
	x.buf.Write([]byte("`uid` = :uid"))
	return x
}

func (x *UserFriendNamedWhere) Fid() *UserFriendNamedWhere {
	x.buf.Write([]byte("`fid` = :fid"))
	return x
}

func (x *UserFriendNamedWhere) State() *UserFriendNamedWhere {
	x.buf.Write([]byte("`state` = :state"))
	return x
}

func (x *UserFriendNamedWhere) Limit(limit, offset int) *UserFriendNamedWhere {
	x.buf.Write([]byte(" limit "))
	x.buf.WriteInt(limit)
	x.buf.WriteByte(',')
	x.buf.WriteInt(offset)
	return x
}

func (x *UserFriendNamedWhere) And() *UserFriendNamedWhere {
	x.buf.Write([]byte(" and "))
	return x
}

func (x *UserFriendNamedWhere) Or() *UserFriendNamedWhere {
	x.buf.Write([]byte(" or "))
	return x
}

func (x *UserFriendNamedWhere) Group(gf func(w *UserFriendNamedWhere)) *UserFriendNamedWhere {
	x.buf.WriteByte('(')
	gf(x)
	x.buf.WriteByte(')')
	return x
}

func (x *UserFriendNamedWhere) Custom(f func(buf *util.Builder)) *UserFriendNamedWhere {
	f(x.buf)
	return x
}

func (x *UserFriendNamedWhere) OnDuplicateKeyUpdate() *UserFriendNamedUpdate {
	x.buf.Write([]byte(" on duplicate key update "))
	var v int
	values := false
	return &UserFriendNamedUpdate{
		buf:    x.buf,
		n:      &v,
		values: &values,
	}
}

func (x *UserFriendNamedWhere) OnDuplicateKeyUpdateValues() *UserFriendNamedUpdate {
	x.buf.Write([]byte(" on duplicate key update "))
	var v int
	values := true
	return &UserFriendNamedUpdate{
		buf:    x.buf,
		n:      &v,
		values: &values,
	}
}

func (x *UserFriendNamedWhere) ToSQL() string {
	return x.buf.String()
}

func (x *UserFriendNamedWhere) OrderBy() *UserFriendNamedOrderBy {
	x.buf.Write([]byte(" order by "))
	var v int
	return &UserFriendNamedOrderBy{
		buf: x.buf,
		n:   &v,
	}
}

type UserFriendNamedOrderBy struct {
	buf *util.Builder
	n   *int
}

func (x *UserFriendNamedOrderBy) Uid() *UserFriendNamedOrderAsc {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`uid`"))
	*x.n++
	return &UserFriendNamedOrderAsc{
		buf: x.buf,
		n:   x.n,
	}
}

func (x *UserFriendNamedOrderBy) Fid() *UserFriendNamedOrderAsc {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`fid`"))
	*x.n++
	return &UserFriendNamedOrderAsc{
		buf: x.buf,
		n:   x.n,
	}
}

func (x *UserFriendNamedOrderBy) State() *UserFriendNamedOrderAsc {
	if *x.n > 0 {
		x.buf.WriteByte(',')
	}
	x.buf.Write([]byte("`state`"))
	*x.n++
	return &UserFriendNamedOrderAsc{
		buf: x.buf,
		n:   x.n,
	}
}

func (x *UserFriendNamedOrderBy) Limit(limit, offset int) *UserFriendNamedOrderBy {
	x.buf.Write([]byte(" limit "))
	x.buf.WriteInt(limit)
	x.buf.WriteByte(',')
	x.buf.WriteInt(offset)
	return x
}

func (x *UserFriendNamedOrderBy) ToSQL() string {
	return x.buf.String()
}

type UserFriendNamedOrderAsc struct {
	buf *util.Builder
	n   *int
}

func (x *UserFriendNamedOrderAsc) Asc() *UserFriendNamedOrderBy {
	x.buf.Write([]byte(" asc "))
	return &UserFriendNamedOrderBy{
		buf: x.buf,
		n:   x.n,
	}
}

func (x *UserFriendNamedOrderAsc) Desc() *UserFriendNamedOrderBy {
	x.buf.Write([]byte(" desc "))
	return &UserFriendNamedOrderBy{
		buf: x.buf,
		n:   x.n,
	}
}

func (x *UserFriendNamedOrderAsc) ToSQL() string {
	return x.buf.String()
}
