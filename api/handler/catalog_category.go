package handler

import (
	"customer-api-gateway/genproto/catalog_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /v1/category/getall [GET]
// @Summary Get all categories
// @Description API for getting all categories
// @Tags category
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllCategory(c *gin.Context) {
	category := &catalog_service.GetListCategoryRequest{}

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

	category.Search = search
	category.Offset = int64(page)
	category.Limit = int64(limit)

	resp, err := h.grpcClient.CategoryService().GetList(c.Request.Context(), category)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating category")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/category/create [POST]
// @Summary Create category
// @Description API for creating categories
// @Tags category
// @Accept  json
// @Produce  json
// @Param		category body  catalog_service.CreateCategory true "category"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateCategory(c *gin.Context) {
	category := &catalog_service.CreateCategory{}
	if err := c.ShouldBindJSON(&category); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.CategoryService().Create(c.Request.Context(), category)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating category")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/category/update [PUT]
// @Summary Update category
// @Description API for Updating categories
// @Tags category
// @Accept  json
// @Produce  json
// @Param		category body  catalog_service.UpdateCategory true "category"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) UpdateCategory(c *gin.Context) {
	category := &catalog_service.UpdateCategory{}
	if err := c.ShouldBindJSON(&category); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.CategoryService().Update(c.Request.Context(), category)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating category")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/category/get/{id} [GET]
// @Summary Get category
// @Description API for getting category
// @Tags category
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	category := &catalog_service.CategoryPrimaryKey{Id: id}

	resp, err := h.grpcClient.CategoryService().GetByID(c.Request.Context(), category)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting category")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/category/delete/{id} [DELETE]
// @Summary Delete category
// @Description API for deleting category
// @Tags category
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	category := &catalog_service.CategoryPrimaryKey{Id: id}

	resp, err := h.grpcClient.CategoryService().Delete(c.Request.Context(), category)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting category")
		return
	}
	c.JSON(http.StatusOK, resp)
}
