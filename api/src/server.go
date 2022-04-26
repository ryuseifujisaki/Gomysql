package main 

import (
	"github.com/ryuseifujisaki/Gomysql/api/src/domain"
	"github.com/ryuseifujisaki/Gomysql/api/src/infrastructure"

	"github.com/labstack/echo/v4"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
	dsn := "root:password@tcp(127.0.0.1:3306)/gomysql?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	dbinit()
	infrastructure.Init()
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}

func dbinit() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.config{})
	if err != nil {
	}
	db.Migrator().CreateTable(domain.User{})
}
