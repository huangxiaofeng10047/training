package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"strings"
	"time"
)

// 数据库配置
const (
	userName    = "root"
	password    = "123456"
	ip          = "localhost"
	port        = "3309"
	dbName      = "test"
	slaveIp     = "localhost"
	slavePort   = "3310"
	slaveUser   = "root"
	slavePasswd = "123456"
)

// Db数据库连接池
var DB *sql.DB
var slaveDb *sql.DB

type User struct {
	id    int64
	name  string
	age   int8
	sex   int8
	phone string
}

// 注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入：_ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}
func InitSlaveDB() {
	path := strings.Join([]string{slaveUser, ":", slavePasswd, "@tcp(", slaveIp, ":", slavePort, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入：_ "github.com/go-sql-driver/mysql"
	slaveDb, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	slaveDb.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	slaveDb.SetMaxIdleConns(10)
	//验证连接
	if err := slaveDb.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("slave connnect success")
}

// 查询操作
func Query() {
	var user User
	rows, e := DB.Query("select * from user where name ='test2'")
	if e == nil {
		errors.New("query incur error")
	}
	for rows.Next() {
		e := rows.Scan(user.sex, user.phone, user.name, user.id, user.age)
		if e != nil {
			fmt.Println(json.Marshal(user))
		}
	}
	rows.Close()
	DB.QueryRow("select * from user where id=1").Scan(user.age, user.id, user.name, user.phone, user.sex)

	stmt, e := DB.Prepare("select * from user where id=?")
	query, e := stmt.Query(1)
	query.Scan()
}

// 查询从库
func QuerySlave() {
	var user User
	rows, e := slaveDb.Query("select * from user where name ='test2'")
	if e == nil {
		errors.New("query incur error")
	}
	for rows.Next() {
		e := rows.Scan(user.sex, user.phone, user.name, user.id, user.age)
		if e != nil {
			fmt.Println(json.Marshal(user))
		}
	}
	rows.Close()
	slaveDb.QueryRow("select * from user where id=1").Scan(user.age, user.id, user.name, user.phone, user.sex)

	stmt, e := slaveDb.Prepare("select * from user where id=?")
	query, e := stmt.Query(1)
	query.Scan()
}
func DeleteUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(user.id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}

func InsertUser(user User) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`name`, `phone`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	var str []string = []string{user.name}
	for i := 0; i < 900000; i++ {
		_, err = stmt.Exec(strings.Join(str, string(i)), user.phone)
		if err != nil {
			panic(err.Error())
		}
	}
	//res, err := stmt.Exec(user.name, user.phone)
	//if err != nil {
	//	fmt.Println("Exec fail")
	//	return false
	//}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	//fmt.Println(res.LastInsertId())
	return true
}

func main() {
	InitDB()
	InsertUser(User{
		name:  "test2",
		phone: "123456789",
		age:   18,
	})
	Query()
	defer DB.Close()

	//InitSlaveDB()
	//主线程 sleep 1s
	//time.Sleep(1 * time.Second)
	//QuerySlave()
	//fmt.Println("sleep 1")
	//主线程 sleep 1s
	//time.Sleep(1 * time.Millisecond)
	//QuerySlave()
	//defer slaveDb.Close()
}

/**
create table user
(
  id bigint(20) not null auto_increment,
  name varchar(255)      default '',
  age int(11)      not null     default 0,
  sex tinyint(3)      not null     default '0',
 phone varchar(45)    not null  default '',
 primary key (id)
) engine=InnoDB auto_increment=3
default charset=utf8;
*/
