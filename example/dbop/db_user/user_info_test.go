package db_user_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/walleframe/svc_db/example/dbop"
	"github.com/walleframe/svc_db/example/dbop/db_user"
	"github.com/walleframe/svc_db/mysqlcfg"
)

func TestUserInfoSQL(t *testing.T) {
	insert := db_user.UserInfoNamedSQL(128).Insert()
	t.Log(insert.Uid().Email().Name().ToSQL())
	t.Log(db_user.UserInfoNamedSQL(128).Delete().Email().And().Name().ToSQL())
	t.Log(db_user.UserInfoNamedSQL(128).Update().Name().Email().Where().Uid().ToSQL())
	t.Log(db_user.UserInfoNamedSQL(128).Select().Uid().Name().Where().Uid().And().Email().Limit(10, 0).ToSQL())
	t.Log(db_user.UserInfoNamedSQL(128).Update().Email().Where().And().Email().ToSQL())
}

func TestUserInfo(t *testing.T) {

	var checkRes = func(res sql.Result, err error) {
		if err != nil {
			panic(err)
		}

		t.Log(res.RowsAffected())
		t.Log(res.LastInsertId())
	}

	var checkError = func(_ any, err error) {
		if err != nil {
			panic(err)
		}
	}
	var checkList = func(datas []*dbop.UserInfo, err error) {
		if err != nil {
			panic(err)
		}
	}

	ctx := context.Background()
	cfg := mysqlcfg.NewConfig("x")
	cfg.DB = "test2024"
	db, err := mysqlcfg.NewMysqlClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	err = db_user.SyncUserInfoDBTable(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	user, err := db_user.NewUserInfoOperation(db)
	if err != nil {
		t.Fatal(err)
	}

	_, err = user.DeleteMany(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	lists, err := user.AllData(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 0, len(lists), "find all interface")

	checkRes(user.Insert(ctx, &dbop.UserInfo{
		Uid:   0,
		Name:  "u1",
		Email: "u1",
	}))

	checkRes(user.DeleteMany(ctx, nil))

	checkRes(user.Insert(ctx, &dbop.UserInfo{
		Uid:   100,
		Name:  "u100",
		Email: "email100",
	}))

	checkRes(user.Insert(ctx, &dbop.UserInfo{
		Uid:   101,
		Name:  "u101",
		Email: "email101",
	}))

	v, err := user.Find(ctx, 100)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, "u100", v.Name, "user name")

	checkRes(user.InsertMany(ctx, []*dbop.UserInfo{
		{
			Name:  "xx1",
			Email: "x",
		},
		{
			Name:  "xx2",
			Email: "x2",
		},
	}))

	checkRes(user.Update(ctx, &dbop.UserInfo{
		Uid:  100,
		Name: "update100",
	}))

	checkRes(user.Upsert(ctx, &dbop.UserInfo{
		Uid:  100,
		Name: "xxx",
	}))

	checkRes(user.UpsertMany(ctx, []*dbop.UserInfo{
		{Uid: 10, Name: "10", Email: "10"},
		{Uid: 100, Name: "100", Email: "100"},
	}))

	v, err = user.Find(ctx, 10)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, "10", v.Name, "upsert values")
	v2, err := user.FindEx(ctx, 100)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, "100", v2.Name, "second")

	list, err := user.FindByKeyArray(ctx, []int64{10, 100})
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 2, len(list), "mul find")

	checkError(user.FindExByKeyArray(ctx, []int64{10, 100}))

	list, err = user.FindByIndexName(ctx, "10", 1, 0)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 1, len(list), "find by index")

	if len(list) > 0 {
		assert.EqualValues(t, 10, list[0].Uid, "find index uid")
	}

	checkError(user.FindExByIndexName(ctx, "10", 1, 0))
	cnt, err := user.CountByIndexName(ctx, "10")
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 1, cnt, "count by index")

	checkError(user.Select(ctx, user.Where(128).Uid().In(10, 100)))
	checkError(user.SelectEx(ctx, user.Where(128).Uid().In(10, 100)))

	err = user.RangeAll(ctx, nil, func(ctx context.Context, row *dbop.UserInfo) bool {
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	err = user.RangeAllEx(ctx, nil, func(ctx context.Context, row *dbop.UserInfoEx) bool {
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	checkError(user.AllDataEx(ctx, nil))

	checkRes(user.Delete(ctx, 100))

	checkList(user.AllData(ctx, nil))

	checkRes(user.DeleteMany(ctx, nil))
}
