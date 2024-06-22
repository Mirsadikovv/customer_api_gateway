package handler

import (
	"customer-api-gateway/api/helpers"
	"customer-api-gateway/genproto/order_status_notes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/order-status [POST]
// @Summary		Update order status by ID
// @Description	This API create the status of an order
// @Tags		order_status
// @Produce		json
// @Param        order body order_status_notes.CreateStatusRequest true "Order Status"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h *handler) CreateOrderstatus(c *gin.Context) {

	order := &order_status_notes.CreateStatusRequest{}

	if err := c.ShouldBindJSON(order); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading request body")
		return
	}

	if err := helpers.Validstatus(order.Status); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "some error input status err ")
		return
	}

	fmt.Println(order)
	resp, err := h.grpcClient.OrderStatus().Create(c.Request.Context(), order)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating order")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// @Router		/v1/order-status/{id} [GET]
// @Summary		Get by id a orderstatus
// @Description	This api get a order by id
// @Tags		order_status
// @Produce		json
// @Param 		id path order_status_notes.OrderPrimaryKeyRequest true "Order Status"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h *handler) OrderStatusById(c *gin.Context) {

	id := &order_status_notes.OrderPrimaryKeyRequest{}
	id.Id = c.Param("id")
	fmt.Println(id, "id______________________")
	data, err := h.grpcClient.OrderStatus().GetByID(c.Request.Context(), id)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading request id")
		return
	}

	c.JSON(http.StatusCreated, data)
}

// @Security ApiKeyAuth
// @Router     /v1/order-status [PATCH]
// @Summary    Update order status by ID
// @Description This API updates the status of an order by ID
// @Tags       order_status
// @Produce    json
// @Param      orders body order_status_notes.OrderPrimaryStatusKeyRequest true "Order Status"
// @Success    200 {object} models.Response
// @Failure    400 {object} models.Response
// @Failure    404 {object} models.Response
// @Failure    500 {object} models.Response
func (h *handler) OrderStatusPatch(c *gin.Context) {
	order := &order_status_notes.OrderPrimaryStatusKeyRequest{}

	if err := c.ShouldBindJSON(order); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading request body")
		return
	}
	data, err := h.grpcClient.OrderStatus().GetStatusByID(c.Request.Context(), order)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while fetching current order status")
		return
	}

	fmt.Println("Current status:", data.Status)
	fmt.Println("New status:", order.Status)

	if err := helpers.Validstatusorder(data.Status, order.Status); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid status transition")
		return
	}
	resp, err := h.grpcClient.OrderStatus().PUTCH(c.Request.Context(), order)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating order status")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router		/v1/order-statusp/{id} [GET]
// @Summary		Get by id a orderstatusputch
// @Description	This API gets an order by ID
// @Tags		order_statusp
// @Produce		json
// @Param 		id path string true "Order Status ID"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h *handler) OrderStatusputch(c *gin.Context) {
	order := &order_status_notes.OrderPrimaryStatusKeyRequest{}

	order.OrderId = c.Param("id")

	data, err := h.grpcClient.OrderStatus().GetStatusByID(c.Request.Context(), order)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading request id")
		return
	}

	c.JSON(http.StatusOK, data)
}
