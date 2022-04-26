package  infrastructure

import (
	"gorm.io/driver/mysql"
    "gorm.io/gorm"

    "api/src/interfaces/database"
)

type sqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	dsn := "root:password@tcp(127.0.0.1:3306)/gomysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.config{})
	if err != nil	{
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db =db
	return sqlHandler
}

func (handler *sqlHandler) Create(obj interface{}){
	handler.db.Create(obj)
}

func (handler *sqlHandler) FindAll(obj interface{}){
	handler.db.Find(obj)
}

func (handler *sqlHandler) DeleteById(obj interface{}, id string){
	handler.db.Delete(obj, id)
}
