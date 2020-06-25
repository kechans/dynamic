package dao

import (
	"database/sql"
	"dynamic/pkg/utils"
	"fmt"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

const (
	ServerConfPath = "conf/database.ini"
	ServerConfName = "database"
	sqlStr         = "select age from user where id = ? "
)

var dbsql *sql.DB

type Database struct {
	Type          string
	User          string
	Password      string
	Host          string
	DataBasesName string
	TablePrefix   string
}

var DatabaseSetting = &Database{}

type User struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	//Status     int    `json:"status" form:"status"` // 0 正常状态， 1删除
	Createtime int64 `json:"createtime" form:"createtime"`
}

func InitMysqlConfig() {
	utils.Setup(ServerConfPath, ServerConfName, DatabaseSetting)
	// user:password@/dbname
	// User:password@tcp(localhost:5555)/dbname?
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?", DatabaseSetting.User, DatabaseSetting.Password, DatabaseSetting.Host, DatabaseSetting.DataBasesName)
	log.Printf("mysql init %s", dsn)
	var err error
	dbsql, err = sql.Open("mysql", dsn)
	fmt.Println(dbsql)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	//Insert(dbsql)
	//Select(dbsql)
}

func FetchRow(sqlstr string, args ...interface{}) (map[string]string, error) {
	//dbRead, err := dbsql.GetRead()
	//if err != nil {
	//	return nil, err
	//}

	return readRow(dbsql, sqlstr, args...)
}

// 取多行数据
func FetchRows(sqlstr string, args ...interface{}) ([]map[string]string, error) {
	//dbRead, err := dbsql.GetRead()
	//if err != nil {
	//	return nil, err
	//}

	return readRows(dbsql, sqlstr, args...)
}

func readRow(dbsql *sql.DB, sqlstr string, args ...interface{}) (map[string]string, error) {
	rows, err := dbsql.Query(sqlstr, args...)

	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make(map[string]string)

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "" //把数据表中所有为null的地方改成“”
			} else {
				value = string(col)
			}

			ret[columns[i]] = value
		}

		break
	}

	rows.Close()

	return ret, err
}
func readRows(dbsql *sql.DB, sqlstr string, args ...interface{}) ([]map[string]string, error) {
	rows, err := dbsql.Query(sqlstr, args...)

	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	var rets = make([]map[string]string, 0)

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var ret = make(map[string]string) //这里要注意(对语法的理解)

		var value string
		for i, col := range values {
			if col == nil {
				value = "" //把数据表中所有为null的地方改成“”
			} else {
				value = string(col)
			}

			ret[columns[i]] = value
		}

		rets = append(rets, ret)
	}

	return rets, err
}

func Insert(dbWrite *sql.DB) int64 {
	sqlstr := "Insert INTO `user` (`username`, `password`, `email`,`age`,`sex`,`tel`,`addr`,`createtime`) VALUES (?,?,?,?,?,?,?,?)"
	args := make([]interface{}, 0)
	args = append(args, "aaaa")
	args = append(args, "32465cvb")
	args = append(args, "32588@qq.com")
	args = append(args, 18)
	args = append(args, "man")
	args = append(args, "17710296356")
	args = append(args, "beijing")
	args = append(args, "2020-05-10 12:23:34")
	result, err := dbWrite.Exec(sqlstr, args...)
	fmt.Println(result)
	if err == nil {
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		return id
	}
	return 0
}

// 获取一条数据
func Select() (map[string]string, error) {
	sqlstr := "SELECT * FROM `user` "
	args := make([]interface{}, 0)
	//args = append(args, "chenheng")
	rows, err := dbsql.Query(sqlstr, args...)

	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret := make(map[string]string)

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "" //把数据表中所有为null的地方改成“”
			} else {
				value = string(col)
			}

			ret[columns[i]] = value
		}

		break
	}

	rows.Close()
	return ret, err
}
