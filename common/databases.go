package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB{
 //open a db connection
 var err error

 DB, err = gorm.Open("mysql", "root:01117042116vero@/callperfect?charset=utf8&parseTime=True&loc=Local")

 if err != nil {
  panic("failed to connect database")
 }

return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}