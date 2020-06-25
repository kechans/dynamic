package service

import (
	"dynamic/models/dao"
)

func GetUserdata() []map[string]string {

	sqlstr := "SELECT * FROM `user` "
	args := make([]interface{}, 0)
	res, err := dao.FetchRows(sqlstr, args...)
	//log.Printf("data is %v", res)
	if err != nil {
		return nil
	}
	return res
}
