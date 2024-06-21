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
// @Router        /v1/shop/create [post]
// @Summary       Create shop
// @Description   API for creating shop
// @Tags          shop
// @Accept        json
// @Produce       json
// @Param         shop body user_service.CreateShop true "shop"
// @Success 200   {object} user_service.Shop
// @Failure 404   {object} models.ResponseError
// @Failure 500   {object} models.ResponseError
func (h *handler) CreateShop(c *gin.Context) {
	var (
		req  user_service.CreateShop
		resp *user_service.Shop
		err  error
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	if err := helpers.ValidatePhone(req.Phone); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while validating phone number "+req.Phone)
		return
	}

	resp, err = h.grpcClient.ShopService().Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "failed to create shop")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /v1/shop/getlist [GET]
// @Summary        Get list of shops
// @Description    API for getting list of shops
// @Tags           shop
// @Accept         json
// @Produce        json
// @Param		   shop query string false "shops"
// @Param		   page query int false "page"
// @Param		   limit query int false "limit"
// @Success 200    {object} user_service.GetListShopResponse
// @Failure 404    {object} models.ResponseError
// @Failure 500    {object} models.ResponseError
func (h *handler) GetListShop(c *gin.Context) {

	var (
		req  user_service.GetListShopRequest
		resp *user_service.GetListShopResponse
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

	resp, err = h.grpcClient.ShopService().GetList(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /v1/shop/getbyid/{id} [GET]
// @Summary        Get a single shop by ID
// @Description    API for getting a single shop by ID
// @Tags           shop
// @Accept         json
// @Produce        json
// @Param          id path string true "shop ID"
// @Success        200 {object} user_service.Shop
// @Failure        404 {object} models.ResponseError
// @Failure        500 {object} models.ResponseError
func (h *handler) GetShopByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *user_service.Shop
		err  error
	)

	req := &user_service.ShopPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.ShopService().GetByID(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router          /v1/shop/update/{id} [PUT]
// @Summary         Update a shop by ID
// @Description     API for updating a shop by ID
// @Tags            shop
// @Accept          json
// @Produce         json
// @Param           id path string true "shop ID"
// @Param           shop body user_service.UpdateShop true "shop"
// @Success         200 {object} user_service.Shop
// @Failure         404 {object} models.ResponseError
// @Failure         500 {object} models.ResponseError
func (h *handler) UpdateShop(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  user_service.UpdateShop
		resp *user_service.Shop
		err  error
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "invalid request body")
		return
	}

	if err := helpers.ValidatePhone(req.Phone); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while validating phone number "+req.Phone)
		return
	}

	req.Id = id
	resp, err = h.grpcClient.ShopService().Update(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router        /v1/shop/delete/{id} [DELETE]
// @Summary       Delete a shop by ID
// @Description   API for deleting a shop by ID
// @Tags          shop
// @Accept        json
// @Produce       json
// @Param         id path string true "shop ID"
// @Success       200 {object} models.Response
// @Failure       404 {object} models.ResponseError
// @Failure       500 {object} models.ResponseError
func (h *handler) DeleteShop(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *emptypb.Empty
	)

	req := &user_service.ShopPrimaryKey{
		Id: id,
	}

	resp, err = h.grpcClient.ShopService().Delete(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
