package routes

import (
	"net/http"

	"github.com/hafif/echoFramework/controllers"
	"github.com/hafif/echoFramework/middleware"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world from echo!")
	})

	//buku routing
	e.GET("buku", controllers.FetchAllData, middleware.IsAuthenticated)

	e.POST("buku", controllers.StoreData)

	e.PUT("buku", controllers.UpdateData)

	e.DELETE("buku", controllers.DeleteData)

	// login and register routing
	e.POST("register", controllers.RegisterUser)
	e.POST("login", controllers.LoginUser)
	return e

}
