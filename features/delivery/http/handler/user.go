package handler

import (
	"online-learning-platform/domain"
	"online-learning-platform/utils/middlewares"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var input domain.LoginRequest
		if err := c.Bind(&input); err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		res, err := uh.UserUsecase.GetUserLogin(c.Request().Context(), input)
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}

func (uh *UserHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var input domain.UserRequest
		if err := c.Bind(&input); err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		res, err := uh.UserUsecase.CreateUser(c.Request().Context(), input)
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}

func (uh *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		_, role, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(fasthttp.StatusUnauthorized, err.Error())
		}

		if role != "admin" {
			return c.JSON(fasthttp.StatusUnauthorized, "admin only")
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		err = uh.UserUsecase.DeleteUser(c.Request().Context(), int64(id))
		if err != nil {
			if strings.Contains(err.Error(), "no") {
				return c.JSON(fasthttp.StatusNotFound, err.Error())
			}
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, "success")
	}
}
