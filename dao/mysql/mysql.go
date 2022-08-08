package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"webapp.io/settings"
)

var (
	db *sqlx.DB // 数据库实例
)

func Init(conf *settings.MySQLConf) (err error) {
	dsn := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True`,
		conf.User, conf.Password,
		conf.Host, conf.Port, conf.Dbname)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.open_max_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.idel_max_conns"))
	return
}

func Close() {
	_ = db.Close()
}
