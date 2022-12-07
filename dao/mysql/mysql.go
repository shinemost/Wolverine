package mysql

import (
	"fmt"
	"hjfu/Wolverine/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Close() {
	_ = db.Close()
}

func Init(config *domain.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password,
		config.Host, config.Port,
		config.DbName,
	)
	// 原生的 是open 这里直接用connect 就行了 里面包含了ping的操作
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))
		return
	}

	// 设置最大的连接数 默认是无限制 如果超出限制了 就会排队等待
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_connection"))
	// 设置最大的空闲连接数 默认是无限制 业务量小的时候 可以把多余的连接释放掉，只保留一定数量的连接数
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_connection"))
	return
}
