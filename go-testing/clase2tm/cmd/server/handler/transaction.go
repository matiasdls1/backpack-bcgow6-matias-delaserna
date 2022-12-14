package handler

/*
Se debe generar la estructura request
Se debe generar la estructura del controlador que tenga como campo el servicio
Se debe generar la función que retorne el controlador
Se deben generar todos los métodos correspondientes a los endpoints
*/

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-testing/clase2tm/internal/transactions"
	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-testing/clase2tm/pkg/web"
)

type request struct {
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(tx transactions.Service) *Transaction {
	return &Transaction{
		service: tx,
	}
}

// ListTransactions godoc
// @Summary List transactions
// @Tags Transactions
// @Description get transactions
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions [get]
func (t *Transaction) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}
		tx, err := t.service.GetAll()
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, tx, ""))
	}
}

// StoreTransactions godoc
// @Summary List transactions
// @Tags Transactions
// @Description store transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param transaction body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Router /transactions [post]
func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}
		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		tx, err := t.service.Store(req.Code, req.Currency, req.Amount, req.Sender, req.Receiver, req.Date)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, tx, ""))
	}
}

func (t *Transaction) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "invalid id"))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Code == "" {
			c.JSON(400, web.NewResponse(400, nil, "code is required"))
			return
		}
		if req.Currency == "" {
			c.JSON(400, web.NewResponse(400, nil, "currency is required"))
			return
		}
		if req.Amount == 0 {
			c.JSON(400, web.NewResponse(400, nil, "amount is required"))
			return
		}
		if req.Sender == "" {
			c.JSON(400, web.NewResponse(400, nil, "sender is required"))
			return
		}
		if req.Receiver == "" {
			c.JSON(400, web.NewResponse(400, nil, "receiver is required"))
			return
		}
		if req.Date == "" {
			c.JSON(400, web.NewResponse(400, nil, "date is required"))
			return
		}
		tx, err := t.service.Update(int(id), req.Code, req.Currency, req.Amount, req.Sender, req.Receiver, req.Date)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, tx, ""))
	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "invalid id"))
			return
		}
		err = t.service.Delete(int(id))
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, fmt.Sprintf("Transaction %d was deleted.", id), ""))
	}
}

func (t *Transaction) UpdateCodeAmount() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "invalid id"))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if req.Code == "" {
			c.JSON(400, web.NewResponse(400, nil, "code is required"))
			return
		}
		if req.Amount == 0 {
			c.JSON(400, web.NewResponse(400, nil, "amount is required"))
			return
		}
		tx, err := t.service.UpdateCodeAmount(int(id), req.Code, req.Amount)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, tx, ""))
	}
}
