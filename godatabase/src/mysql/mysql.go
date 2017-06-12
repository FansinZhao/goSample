package main

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//连接数据库
	db, err := sql.Open("mysql", "root:root@tcp(172.20.0.193:3306)/test?charset=utf8")
	checkErr(err)
	if db != nil {
		fmt.Println("连接数据库成功!")
	}
	defer db.Close()
	//插入sql语句
	stmt, e0 := db.Prepare("insert userinfo (username,departname,created) values (?,?,?)")
	checkErr(e0)
	//准备参数
	ret, e := stmt.Exec("赵锋", "交易平台", "2016-12-21")
	checkErr(e)
	ret, e = db.Exec("insert userinfo (username,departname,created) values (?,?,?)", "赵锋db", "交易平台", "2016-12-21")
	checkErr(e)
	if ret != nil {
		fmt.Println("插入成功!")
	}
	id, e := ret.LastInsertId()
	checkErr(e)
	fmt.Println("最后一条插入记录:", id)
	stmt, err = db.Prepare("update userinfo set username=? where uid = ?")
	checkErr(err)
	ret, e = stmt.Exec("zhaofeng", 1)
	checkErr(e)
	a, e := ret.RowsAffected()
	checkErr(e)
	fmt.Println("影响行数:", a)

	stmt, err = db.Prepare("select * from userinfo")
	checkErr(err)
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Println(uid, username, departname, created)
	}

	ret, err = db.Exec("delete from userinfo where uid =?", id)
	checkErr(err)
	a, err = ret.RowsAffected()
	fmt.Println("删除记录数:", a)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
