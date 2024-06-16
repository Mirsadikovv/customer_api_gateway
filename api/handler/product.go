package handler

import (
	"customer-api-gateway/genproto/product_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /product/getall [GET]
// @Summary Get all categories
// @Description API for getting all categories
// @Tags product
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllProduct(c *gin.Context) {
	product := &product_service.GetListProductRequest{}

	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while parsing page")
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while parsing limit")
		return
	}

	product.Search = search
	product.Offset = int64(page)
	product.Limit = int64(limit)

	resp, err := h.grpcClient.ProductService().GetList(c.Request.Context(), product)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating product")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product/create [POST]
// @Summary Create product
// @Description API for creating categories
// @Tags product
// @Accept  json
// @Produce  json
// @Param		product body  product_service.CreateProduct true "product"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateProduct(c *gin.Context) {
	product := &product_service.CreateProduct{}
	if err := c.ShouldBindJSON(&product); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.ProductService().Create(c.Request.Context(), product)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating product")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product/update [PUT]
// @Summary Update product
// @Description API for Updating categories
// @Tags product
// @Accept  json
// @Produce  json
// @Param		product body  product_service.UpdateProduct true "product"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateProduct(c *gin.Context) {
	product := &product_service.UpdateProduct{}
	if err := c.ShouldBindJSON(&product); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.ProductService().Update(c.Request.Context(), product)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating product")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product/get/{id} [GET]
// @Summary Get product
// @Description API for getting product
// @Tags product
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetProductById(c *gin.Context) {
	id := c.Param("id")
	product := &product_service.ProductPrimaryKey{Id: id}

	resp, err := h.grpcClient.ProductService().GetByID(c.Request.Context(), product)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting product")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /product/delete/{id} [DELETE]
// @Summary Delete product
// @Description API for deleting product
// @Tags product
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	product := &product_service.ProductPrimaryKey{Id: id}

	resp, err := h.grpcClient.ProductService().Delete(c.Request.Context(), product)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting product")
		return
	}
	c.JSON(http.StatusOK, resp)
}
