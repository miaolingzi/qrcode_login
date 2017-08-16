package qrdb

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dbHelper *sql.DB

func init() {
	fmt.Println("database start to connect")
	db, err := sql.Open("mysql", "root:qazwsxedc@tcp(localhost:3306)/user?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
	fmt.Println("database connect successfully")
	dbHelper = db
}

func checkDBNull() {
	if dbHelper == nil {
		fmt.Errorf("mysql not connected")
		return
	}
}

/*************************************************AppID相关操作*******************************************************/
func isValidAppID(appid string) bool {
	checkDBNull()
	query, _ := dbHelper.Query("SELECT * FROM APP WHERE appid = ?", appid)
	defer query.Close()
	columns, _ := query.Columns()
	if len(columns) != 1 {
		//不存在appid，参数错误
		fmt.Errorf("appid %s not exists", appid)
		return false
	}
	if len(columns) == 1 {
		scanArgs := make([]interface{}, len(columns))
		values := make([][]byte, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		for query.Next() {
			query.Scan(scanArgs...)
			row := make(map[string]string)
			for k, v := range values {
				key := columns[k]
				row[key] = string(v)
			}
			if row["appid"] == "" {
				return false
			}
		}
	}
	return true
}

/*************************************************UUID相关操作*******************************************************/
func isValidUUID(uuid string) {

}

func addUUID(uuid string) {

}

func deleteUUID(uuid string) {

}

func bindUUID(uuid string, user_id string) {

}
