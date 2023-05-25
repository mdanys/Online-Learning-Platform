package http

import (
	"online-learning-platform/domain"
	"online-learning-platform/features/delivery/http/handler"

	"github.com/labstack/echo/v4"
)

func RouteAPI(e *echo.Echo, category domain.CategoryUsecase, user domain.UserUsecase) {
	handlerCategory := handler.CategoryHandler{CategoryUsecase: category}
	handlerUser := handler.UserHandler{UserUsecase: user}

	// Category
	e.POST("/category", handlerCategory.CreateCategory())
	e.GET("/category", handlerCategory.GetCategories())
	e.GET("/category/:id", handlerCategory.GetCategory())
	e.PUT("/category/:id", handlerCategory.UpdateCategory())

	// User
	e.POST("/login", handlerUser.Login())
	e.POST("/user", handlerUser.CreateUser())
	e.DELETE("/user/:id", handlerUser.DeleteUser())
}
