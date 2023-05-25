package main

import (
	"fmt"
	"online-learning-platform/config"
	"online-learning-platform/features/delivery/http"
	"online-learning-platform/features/repository/mysql"
	"online-learning-platform/features/usecase"

	"online-learning-platform/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	dbConn := database.InitDatabase(cfg)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	categoryRepo := mysql.NewMySQLCategoryRepository(&dbConn)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	http.RouteAPI(e, categoryUsecase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ServerPort)))
}
