package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"

	"github.com/ryuseifujisaki/Gomysql/api/src/interfaces/database"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := "gomysql:password@tcp(gomysql-db:3306)/gomysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}

func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}

func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
	fmt.Println("asas")
	fmt.Println("asas"     )
	fmt.Println("asas")
	fmt.Println("asas")
	fmt.Println("asas")
}
//  colour
