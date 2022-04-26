package infrastructure 

import (
	controllers "api/src/interfaces/api"
	"net/http"

	"github.com/labstack/echo"
)

func Init() {
	e := echo.New()
	userController := controllers.NewUserController(NewSqlHandler())

	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUser()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})

	e.GET("/users", func(c echo.Context) error {
		userController.Create(c)
		return c.string(http.StatusOK, "Create")
	})

	e.Delete("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		userController.Delete(id)
		return c.String(http.StatusOK, "Deleted")
	})

	e.Logger.Fatal(e.Start(":1323"))
}