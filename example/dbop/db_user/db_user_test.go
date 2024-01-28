package db_user_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/walleframe/svc_db"
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

	cnt, err := user.Count(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 2, cnt, "count")

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
	cnt, err = user.CountByIndexName(ctx, "10")
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

func TestUserFriend(t *testing.T) {

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

	ctx := context.Background()
	cfg := mysqlcfg.NewConfig("x")
	cfg.DB = "test2024"
	db, err := mysqlcfg.NewMysqlClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	err = db_user.SyncUserFriendDBTable(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	friend, err := db_user.NewUserFriendOperation(db)
	if err != nil {
		t.Fatal(err)
	}

	_, err = friend.DeleteMany(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	lists, err := friend.AllData(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 0, len(lists), "find all interface")

	checkRes(friend.Insert(ctx, &dbop.UserFriend{
		Uid:   1,
		Fid:   2,
		State: 1,
	}))

	checkRes(friend.DeleteMany(ctx, nil))

	checkRes(friend.Insert(ctx, &dbop.UserFriend{
		Uid:   100,
		Fid:   2,
		State: 1,
	}))

	checkRes(friend.Insert(ctx, &dbop.UserFriend{
		Uid:   101,
		Fid:   2,
		State: 1,
	}))

	cnt, err := friend.Count(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 2, cnt, "count")

	v, err := friend.Find(ctx, 100, 2)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 1, v.State, "user name")

	checkRes(friend.InsertMany(ctx, []*dbop.UserFriend{
		{
			Uid:   1,
			Fid:   2,
			State: 1,
		},
		{
			Uid:   101,
			Fid:   23,
			State: 1,
		},
	}))

	checkRes(friend.Update(ctx, &dbop.UserFriend{
		Uid:   1,
		Fid:   2,
		State: 3,
	}))

	checkRes(friend.Upsert(ctx, &dbop.UserFriend{
		Uid:   101,
		Fid:   2,
		State: 8,
	}))

	checkRes(friend.UpsertMany(ctx, []*dbop.UserFriend{
		{
			Uid:   101,
			Fid:   2,
			State: 2,
		},
		{
			Uid:   1001,
			Fid:   2,
			State: 100,
		},
	}))

	v, err = friend.Find(ctx, 1001, 2)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 100, v.State, "upsert values")
	v2, err := friend.FindEx(ctx, 1001, 2)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 100, v2.State, "second")

	list, err := friend.FindByKeyArray(ctx, []db_user.UserFriendKey{{101, 2}, {1001, 2}})
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 2, len(list), "mul find")

	checkError(friend.FindExByKeyArray(ctx, []db_user.UserFriendKey{{101, 2}, {1001, 2}}))

	checkError(friend.Select(ctx, friend.Where(128).Uid().In(10, 100)))
	checkError(friend.SelectEx(ctx, friend.Where(128).Uid().In(10, 100)))

	err = friend.RangeAll(ctx, nil, func(ctx context.Context, row *dbop.UserFriend) bool {
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	err = friend.RangeAllEx(ctx, nil, func(ctx context.Context, row *dbop.UserFriendEx) bool {
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	checkError(friend.AllDataEx(ctx, nil))

	checkRes(friend.Delete(ctx, 101, 2))

	checkError(friend.AllData(ctx, nil))

	checkRes(friend.DeleteMany(ctx, nil))
}

func TestUserTest(t *testing.T) {

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

	ctx := context.Background()
	cfg := mysqlcfg.NewConfig("x")
	cfg.DB = "test2024"
	db, err := mysqlcfg.NewMysqlClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	err = db_user.SyncUserTestDBTable(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	user, err := db_user.NewUserTestOperation(db)
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

	checkRes(user.Insert(ctx, &dbop.UserTest{
		Uid:   1,
		Xxx:   2,
		State: 1,
	}))

	checkRes(user.DeleteMany(ctx, nil))

	checkRes(user.Insert(ctx, &dbop.UserTest{
		Uid:   100,
		Xxx:   2,
		State: 1,
	}))

	checkRes(user.Insert(ctx, &dbop.UserTest{
		Uid:   101,
		Xxx:   2,
		State: 1,
	}))

	cnt, err := user.Count(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 2, cnt, "count")

	checkRes(user.InsertMany(ctx, []*dbop.UserTest{
		{
			Uid:   3,
			Xxx:   2,
			State: 1,
		},
		{
			Uid:   102,
			Xxx:   23,
			State: 1,
		},
	}))

	checkError(user.Select(ctx, user.Where(128).Uid().In(10, 100)))
	checkError(user.SelectEx(ctx, user.Where(128).Uid().In(10, 100)))

	err = user.RangeAll(ctx, nil, func(ctx context.Context, row *dbop.UserTest) bool {
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	err = user.RangeAllEx(ctx, nil, func(ctx context.Context, row *dbop.UserTestEx) bool {
		return true
	})
	if err != nil {
		t.Fatal(err)
	}

	checkError(user.FindByIndexXxx(ctx, 101, 2, 1, 0))
	checkError(user.FindExByIndexXxx(ctx, 101, 2, 1, 0))
	checkError(user.CountByIndexXxx(ctx, 101, 2))
	checkError(user.DeleteByIndexXxx(ctx, 101, 2))

	checkError(user.AllDataEx(ctx, nil))

	checkError(user.AllData(ctx, nil))

	checkRes(user.DeleteMany(ctx, nil))
}

func TestUserInfoSql(t *testing.T) {
	cfg := mysqlcfg.NewConfig("mysql.local")
	cfg.DB = "test"
	// db连接
	db, err := mysqlcfg.NewMysqlClient(cfg)
	if err != nil {
		t.Fatalf("create mysql client err: %s", err)
	}

	// sync table columns
	err = svc_db.SyncTableColumns(context.Background(), db, "user_info", db_user.UserInfoSQL_Create, db_user.UserInfoSQL_TableColumns)
	assert.NoErrorf(t, err, "swap db_user.user_info pointer, sync columns failed")

	// 插入
	t.Run("test.insert", func(t *testing.T) {
		insertSql := db_user.UserInfoNamedSQL(128).Insert().Uid().Email().Name().ToSQL()

		datas := []dbop.UserInfo{
			{
				Uid:   1001,
				Name:  "named1001",
				Email: "1001@xx.mail",
			},
			{
				Uid:   1002,
				Name:  "named1002",
				Email: "1002@xx.mail",
			},
			{
				Uid:   1003,
				Name:  "named1003",
				Email: "1003@xx.mail",
			},
		}

		for _, v := range datas {
			_, err = db.NamedExec(insertSql, v)
			assert.NoError(t, err, "test.insert err")
		}

	})

	// 查询
	t.Run("test.select", func(t *testing.T) {
		selectSql := db_user.UserInfoNamedSQL(128).Select().Uid().Email().Name().ToSQL()

		row, err := db.Query(selectSql)
		assert.NoError(t, err, "test.select query err")

		for row.Next() {
			var userInfo dbop.UserInfo
			assert.NoError(t, row.Scan(&userInfo.Uid, &userInfo.Email, &userInfo.Name), "test.select scan err")

			t.Logf("test.select ret: %v", userInfo)
		}
	})

	// 更新
	t.Run("test.update", func(t *testing.T) {
		updateSql := db_user.UserInfoNamedSQL(128).Update().Name().Where().Uid().ToSQL()

		_, err = db.NamedExec(updateSql, dbop.UserInfo{Uid: 1003, Name: "named1003xx"})
		assert.NoError(t, err, "test.update err")
	})

	// 删除
	t.Run("test.delete", func(t *testing.T) {
		deleteSql := db_user.UserInfoNamedSQL(128).Delete().Uid().ToSQL()

		_, err = db.NamedExec(deleteSql, dbop.UserInfo{Uid: 1003})
		assert.NoError(t, err, "test.delete err")
	})

}
