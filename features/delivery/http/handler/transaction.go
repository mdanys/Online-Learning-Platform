package handler

import (
	"online-learning-platform/domain"
	"online-learning-platform/utils/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

type TransactionHandler struct {
	TransactionUsecase domain.TransactionUsecase
}

func (th *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, _ := middlewares.ExtractToken(c)

		var input domain.TransactionRequest
		err = c.Bind(&input)
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		input.UserID = int64(id)

		res, err := th.TransactionUsecase.CreateTransaction(c.Request().Context(), input)
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}

func (th *TransactionHandler) GetTransactionByID() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(fasthttp.StatusBadRequest, err.Error())
		}

		res, err := th.TransactionUsecase.GetTransactionByID(c.Request().Context(), int64(id))
		if err != nil {
			return c.JSON(fasthttp.StatusInternalServerError, err.Error())
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
