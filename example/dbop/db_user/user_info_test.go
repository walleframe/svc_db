package db_user_test

import (
	"testing"

	"github.com/walleframe/svc_db/example/dbop/db_user"
)

func TestUserInfoSQL(t *testing.T) {
	insert := db_user.UserInfoNamedSQL(128).Insert()
	t.Log(insert.Uid().Email().Name().ToSQL())
	t.Log(db_user.UserInfoNamedSQL(128).Delete().Email().And().Name().ToSQL())
	t.Log(db_user.UserInfoNamedSQL(128).Update().Name().Email().Where().Uid().ToSQL())
	t.Log(db_user.UserInfoNamedSQL(128).Select().Uid().Name().Where().Uid().And().Email().Limit(10, 0).ToSQL())
}
