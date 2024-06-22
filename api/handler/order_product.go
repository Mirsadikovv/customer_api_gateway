package handler

import (
	"customer-api-gateway/genproto/order_product"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/order [post]
// @Summary Create Order_product
// @Description API for creating a order_product
// @Tags  order_product
// @Accept       json
// @Produce      json
// @Param        order body order_product.CreateOrderProducts true "Order_Product"
// @Success      201 {object} models.ResponseSuccess
// @Failure      404 {object} models.ResponseError
// @Failure      500 {object} models.ResponseError
func (h *handler) CreateOrderProducts(c *gin.Context) {

	order := &order_product.CreateOrderProducts{}

	if err := c.ShouldBindJSON(order); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading request body")
		return
	}
	fmt.Println(order)
	resp, err := h.grpcClient.ProducOrderService().Create(c.Request.Context(), order)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating order")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// @Router		/v1/order/{id} [get]
// @Summary		Get by id a order
// @Description	This api get a order by id
// @Tags		order_product
// @Produce		json
// @Param 		id path order_product.OrderProductsPrimaryKey true "order_product.OrderProductsPrimaryKey"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *handler) OrderProductetById(c *gin.Context) {

	id := &order_product.OrderProductsPrimaryKey{}
	id.Id = c.Param("id")
	fmt.Println(id, "id______________________")
	data, err := h.grpcClient.ProducOrderService().GetByID(c.Request.Context(), id)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading request id")
		return
	}

	c.JSON(http.StatusCreated, data)
}
