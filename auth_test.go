package auth

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Test_DB(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@/ohoh?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	DB = &db
	t.Log(DB)
}

func Test_CreateTables(t *testing.T) {

}
