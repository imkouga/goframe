package mysql

import (
	"github.com/imkouga/gocore/cfg/conf"
	"github.com/imkouga/gocore/store/mysqlx"
)

func Init() error {
	return load()
}

func Reload() error {
	return load()
}

func load() error {

	host := conf.GetValueByStringCarryDefault("mysql", "host", "127.0.0.1")
	port := conf.GetValueByStringCarryDefault("mysql", "port", "3306")
	username := conf.GetValueByStringCarryDefault("mysql", "username", "")
	password := conf.GetValueByStringCarryDefault("mysql", "password", "")
	database := conf.GetValueByStringCarryDefault("mysql", "database", "")

	dsn := mysqlx.Dsn(
		mysqlx.WithHost(host),
		mysqlx.WithPort(port),
		mysqlx.WithUsername(username),
		mysqlx.WithPassword(password),
		mysqlx.WithDatabase(database),
	)
	if err := mysqlx.LoadMySQL(dsn); nil != err {
		return err
	}
	return nil
}
