package mysql

import (
	"fmt"
	"hjfu/Wolverine/domain"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitMysqlX() (err error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/report"
	// 原生的 是open 这里直接用connect 就行了 里面包含了ping的操作
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Second * 10)
	// 设置最大的连接数 默认是无限制 如果超出限制了 就会排队等待
	db.SetMaxOpenConns(200)
	// 设置最大的空闲连接数 默认是无限制 业务量小的时候 可以把多余的连接释放掉，只保留一定数量的连接数
	db.SetMaxIdleConns(10)
	return nil
}

func QueryRowX() {
	sqlStr := "select user_id,password,login_name from sys_user where user_id=?"
	var u domain.SysUserInfo
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:", u)
}

func QueryUserRowsX() {
	sqlStr := "select user_id,user_name,password from sys_user where user_id>1"
	u := domain.SysUserInfo{}
	rows, err := db.Queryx(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err := rows.StructScan(&u)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", u)
	}
}

func InsertX() {
	sqlStr := "insert into sys_user(user_name,login_name) values(?,?)"
	ret, err := db.Exec(sqlStr, "李彦宏", "liyanhong")
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err2 := ret.LastInsertId()
	if err2 != nil {
		fmt.Println("get lastid error:", err2)
		return
	}
	fmt.Println("insert success,user_id: ", id)
}

func BetterInsert() {
	sqlStr := `insert into sys_user(user_name,login_name) values(:user_name,:login_name) `
	ret, err := db.NamedExec(sqlStr, map[string]interface{}{
		"user_name":  "大强子",
		"login_name": "DAQIAONGZI",
	})
	if err != nil {
		panic(err)
	}
	id, err2 := ret.LastInsertId()
	if err2 != nil {
		fmt.Println("get lastid error:", err2)
		return
	}
	fmt.Println("insert success,user_id: ", id)

}

func InsertUsersX(users []domain.SysUserInfo) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.UserName)
		valueArgs = append(valueArgs, u.LoginName)
	}
	// strings.Join(valueStrings, ",")  ===》(?, ?),(?, ?),(?, ?)
	// 自行拼接要执行的具体语句 形如：INSERT INTO user (name, age) VALUES (?, ?),(?, ?),(?, ?)
	stmt := fmt.Sprintf("INSERT INTO sys_user(user_name,login_name) VALUES %s",
		strings.Join(valueStrings, ","))
	fmt.Println(valueArgs)
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

func InsertMoreUsersX(users []domain.SysUserInfo) error {
	_, err := db.NamedExec("INSERT INTO sys_user(user_name,login_name) VALUES (:user_name, :login_name)", users)
	return err
}
