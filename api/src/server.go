package main 

import (
	//"github.com/ryuseifujisaki/Gomysql/api/src/domain"
	"github.com/ryuseifujisaki/Gomysql/api/src/infrastructure"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
	"net/http"
)

var (
	db *gorm.DB
	err error
	dsn = "gomysql:password@tcp(gomysql-db:3306)/gomysql?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	dbinit()
	infrastructure.Init()
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods:    []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Logger.Fatal(e.Start(":1323"))
}

func dbinit() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	}
	//db.Migrator().CreateTable(domain.User{})
}
