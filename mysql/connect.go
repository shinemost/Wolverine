package mysql

import (
	"database/sql"
	"fmt"
	"hjfu/Wolverine/domain"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() (db *sql.DB, err error) {
	dsn := "root:password@tcp(127.0.0.1:3306)/report"
	db, err = sql.Open("mysql", dsn)
	// defer db.Close() 这个地方defer 不能写在error判断的前面
	//  因为 如果open失败 db一般都是nil nil的close一般就报错了 而且打印不出关键的数据库
	//  失败信息
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil

}

func QueryUser(db *sql.DB) {
	sqlStr := "select user_id,password,login_name from sys_user where user_id=?"
	var u domain.SysUserInfo
	// row := db.QueryRow(sqlStr, 1)
	//queryRow 之后 一定要执行scan  否则 持有的数据库连接 不会释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.UserId, &u.Password, &u.LoginName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:", u)
}

func QueryUserRows(db *sql.DB) {
	sqlStr := "select user_id,user_name,password from sys_user where user_id>?"
	rows, err := db.Query(sqlStr, 1)
	if err != nil {
		fmt.Println(err)
	}
	// 这个关闭连接的操作 一定不要忘记
	defer rows.Close()
	// 实际上 next 函数 到最后也会释放连接的，但是有时候我们for循环可能会跳出
	// 所以 为了保险 我们是用defer 来保证 连接一定会被释放
	for rows.Next() {
		var u domain.SysUserInfo
		err := rows.Scan(&u.UserId, &u.UserName, &u.Password)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(u)
	}
}

func InsertUser(db *sql.DB) {
	sqlStr := "insert into sys_user(user_name,login_name) values(?,?)"
	ret, err := db.Exec(sqlStr, "马云", "马爸爸")
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err2 := ret.LastInsertId()
	if err2 != nil {
		fmt.Println("get lastid error:", err2)
		return
	}
	fmt.Println("insert success,id: ", id)
}

func PrepareTest(db *sql.DB) {
	sqlStr := "select user_id,user_name,password from sys_user where user_id>?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	rows, err2 := stmt.Query(0)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u domain.SysUserInfo
		err := rows.Scan(&u.UserId, &u.UserName, &u.Password)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(u)
	}
}

func TransactionDemo(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Println("begin trans failed")
		return
	}

	sqlstr1 := "update sys_user set login_name='我来改一下' where user_id=?"
	_, err2 := tx.Exec(sqlstr1, 103)
	if err2 != nil {
		tx.Rollback()
		fmt.Println(err2)
		return
	}

	sqlstr2 := "update sys_user set email='supertain147@163.com' where user_id=?"
	_, err3 := tx.Exec(sqlstr2, 109)
	if err3 != nil {
		tx.Rollback()
		fmt.Println(err3)
		return
	}

	err4 := tx.Commit()
	if err4 != nil {
		tx.Rollback()
		fmt.Println(err4)
		return
	}

}
