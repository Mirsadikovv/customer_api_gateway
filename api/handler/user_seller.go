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
// @Router        /v1/seller/create [post]
// @Summary       Create seller
// @Description   API for creating seller
// @Tags          seller
// @Accept        json
// @Produce       json
// @Param         seller body user_service.CreateSeller true "seller"
// @Success 200   {object} user_service.CreateSeller
// @Failure 404   {object} models.ResponseError
// @Failure 500   {object} models.ResponseError
func (h *handler) CreateSeller(c *gin.Context) {
	var (
		req  user_service.CreateSeller
		resp *user_service.Seller
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

	resp, err = h.grpcClient.SellerService().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "failed to create customer")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /v1/seller/getlist [GET]
// @Summary        Get List seller
// @Description    API for getting list seller
// @Tags           seller
// @Accept         json
// @Produce        json
// @Param		   seller query string false "sellers"
// @Param		   page query int false "page"
// @Param		   limit query int false "limit"
// @Success 200    {object} user_service.GetListSellerResponse
// @Failure 404    {object} models.ResponseError
// @Failure 500    {object} models.ResponseError
func (h *handler) GetListSeller(c *gin.Context) {
	var (
		req  user_service.GetListSellerRequest
		resp *user_service.GetListSellerResponse
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

	resp, err = h.grpcClient.SellerService().GetList(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /v1/seller/getbyid/{id} [GET]
// @Summary        Get a single seller by ID
// @Description    API for getting a single seller by ID
// @Tags           seller
// @Accept         json
// @Produce        json
// @Param          id path string true "seller ID"
// @Success        200 {object} user_service.Seller
// @Failure        404 {object} models.ResponseError
// @Failure        500 {object} models.ResponseError
func (h *handler) GetSellerByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *user_service.Seller
		err  error
	)

	req := &user_service.SellerPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.SellerService().GetByID(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router          /v1/seller/update/{id} [PUT]
// @Summary         Update a seller by ID
// @Description     API for updating a seller by ID
// @Tags            seller
// @Accept          json
// @Produce         json
// @Param           id path string true "seller ID"
// @Param           seller body user_service.UpdateSeller true "seller"
// @Success         200 {object} user_service.UpdateSeller
// @Failure         404 {object} models.ResponseError
// @Failure         500 {object} models.ResponseError
func (h *handler) UpdateSeller(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  user_service.UpdateSeller
		resp *user_service.Seller
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
	resp, err = h.grpcClient.SellerService().Update(c, &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router        /v1/seller/delete/{id} [DELETE]
// @Summary       Delete a seller by ID
// @Description   API for deleting a seller by ID
// @Tags          seller
// @Accept        json
// @Produce       json
// @Param         id path string true "seller ID"
// @Success       200 {object} models.ResponseError
// @Failure       404 {object} models.ResponseError
// @Failure       500 {object} models.ResponseError
func (h *handler) DeleteSeller(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *emptypb.Empty
	)

	req := &user_service.SellerPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.SellerService().Delete(c, req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
