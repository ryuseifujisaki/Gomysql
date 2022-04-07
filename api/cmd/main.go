package main

import (
	"net/http"
	_ "github.com/pkg/errors"
	"github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	
	//インスタンスの作成
	e := echo.New()

	//ログ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//ルーティング
	e.GET("/",hello)

	e.Logger.Fatal(e.Start(":1323"))
}

//ハンドラーを定義
func hello (c echo.Context) error {
	return c.String(http.StatusOK, "Hello,World")
}