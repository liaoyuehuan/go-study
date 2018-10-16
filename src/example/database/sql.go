package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"strings"
	"log"
)

//定义连接mysql参数
const (
	username    = "xiaoliao"
	password    = "123456"
	host        = "192.168.50.74"
	port        = "3306"
	dbName      = "test"
	maxCons     = 20
	maxIdleCons = 5
)

func ShowRow(rows *sql.Rows) {
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	fmt.Println(columns)

}

func GetDbConnect() (*sql.DB,error){
	//打开数据库,一般不用关闭
	dataSourceNmae := strings.Join([]string{username, ":", password, "@", "tcp(", host, ":", port, ")", "/", dbName}, "")
	db, err := sql.Open("mysql", dataSourceNmae)

	if err != nil {
		return nil,err
	}
	//设置数据库最大连接数目
	db.SetMaxOpenConns(maxCons)

	//设置数据库最大限制连接数目
	db.SetMaxIdleConns(maxIdleCons)

	if err != nil {
		return nil,err
	}
	//ping
	if err = db.Ping(); err != nil {
		return nil,err
	}
	return db,err
}

func testMysql() {
	//打开数据库,一般不用关闭
	dataSourceNmae := strings.Join([]string{username, ":", password, "@", "tcp(", host, ":", port, ")", "/", dbName}, "")
	db, err := sql.Open("mysql", dataSourceNmae)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect success!")

	//设置数据库最大连接数目
	db.SetMaxOpenConns(maxCons)

	//设置数据库最大限制连接数目
	db.SetMaxIdleConns(maxIdleCons)

	//ping
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("ping success!")
	}

	//查询
	rows, err := db.Query("select source,`time` from hhd_news limit 2")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//查询单条
	var v_source string
	err = db.QueryRow("select source from  hhd_news where id = ? ", 3).Scan(&v_source)
	switch err {
	case sql.ErrNoRows:
		log.Println("no record")
	case nil:
		fmt.Printf("source : %s\n", v_source)
	default:
		log.Fatal(err)
	}

	//显示结果
	for rows.Next() {
		var v_source string
		var v_time string
		err = rows.Scan(&v_source, &v_time)
		if err != nil {
			panic(err)
		}
		fmt.Println(v_source, v_time)
	}

	//关闭数据库
	defer db.Close()
}

func testStmt()  {
	db,err := GetDbConnect()
	if err != nil{
		panic(err)
	}
	stmt,err := db.Prepare("select source from  hhd_news where id =  ? ")
	if err != nil{
		panic(err)
	}
	var source string
	err = stmt.QueryRow(3).Scan(&source)
	if err != nil{
		panic(err)
	}
	fmt.Printf("source is %s\n",source)

}

func main() {
	testStmt()
}
