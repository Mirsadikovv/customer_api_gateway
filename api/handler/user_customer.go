package handler

import (
	"customer-api-gateway/api/helpers"
	"customer-api-gateway/genproto/user_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

// @Security ApiKeyAuth
// @Router        /v1/customer/create [post]
// @Summary       Create customer
// @Description   API for creating customer
// @Tags          customer
// @Accept        json
// @Produce       json
// @Param order   body     user_service.CreateCustomer true "customer"
// @Success 200   {object} user_service.CreateCustomer
// @Failure 404   {object} models.ResponseError
// @Failure 500   {object} models.ResponseError
func (h *handler) CreateCustomer(c *gin.Context) {

	var (
		req  user_service.CreateCustomer
		resp *user_service.Customer
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	if err := helpers.ValidatePhone(req.Phone); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while validating phone number"+req.Phone)
		return
	}

	if err := helpers.ValidateEmailAddress(req.Email); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "erroe while validating email"+req.Email)
	}

	resp, err = h.grpcClient.UserService().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "failed to create customer")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /v1/customer/getlist [GET]
// @Summary        Get List Customer
// @Description    API for getting list customer
// @Tags           customer
// @Accept         json
// @Produce        json
// @Param		   customer query string false "customers"
// @Param		   page query int false "page"
// @Param		   limit query int false "limit"
// @Success 200    {object} user_service.GetListCustomerResponse
// @Failure 404    {object} models.ResponseError
// @Failure 500    {object} models.ResponseError
func (h *handler) GetListCustomer(c *gin.Context) {
	var (
		req  user_service.GetListCustomerRequest
		resp *user_service.GetListCustomerResponse
		err  error
	)

	req.Search = c.Query("search")

	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while parsing page")
		return
	}

	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while parsing limit")
		return
	}

	req.Page = page
	req.Limit = limit

	resp, err = h.grpcClient.UserService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /v1/customer/getbyid/{id} [GET]
// @Summary        Get a single customer by ID
// @Description    API for getting a single category by ID
// @Tags           customer
// @Accept         json
// @Produce        json
// @Param          id path string true "customer ID"
// @Success        200 {object} user_service.Customer
// @Failure        404 {object} models.ResponseError
// @Failure        500 {object} models.ResponseError
func (h *handler) GetCustomerByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *user_service.Customer
		err  error
	)

	req := &user_service.CustomerPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.UserService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router          /v1/customer/update/{id} [PUT]
// @Summary         Update a customer by ID
// @Description     API for updating a customer by ID
// @Tags            customer
// @Accept          json
// @Produce         json
// @Param           id path string true "customer ID"
// @Param           category body user_service.UpdateCustomer true "customer"
// @Success         200 {object} user_service.Customer
// @Failure         404 {object} models.ResponseError
// @Failure         500 {object} models.ResponseError
func (h *handler) UpdateCustomer(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  user_service.UpdateCustomer
		resp *user_service.Customer
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	if err := helpers.ValidatePhone(req.Phone); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while validating phone number"+req.Phone)
		return
	}

	if err := helpers.ValidateEmailAddress(req.Email); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "erroe while validating email"+req.Email)
	}

	req.Id = id
	resp, err = h.grpcClient.UserService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router        /v1/customer/delete/{id} [DELETE]
// @Summary       Delete a customer by ID
// @Description   API for deleting a customer by ID
// @Tags          customer
// @Accept        json
// @Produce       json
// @Param         id path string true "customer ID"
// @Success       200 {object} models.ResponseError
// @Failure       404 {object} models.ResponseError
// @Failure       500 {object} models.ResponseError
func (h *handler) DeleteCustomer(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *emptypb.Empty
	)

	req := &user_service.CustomerPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.UserService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
