package handler

import (
	ct "customer-api-gateway/genproto/catalog_service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /get_list_category [GET]
// @Summary Create On Demand Order
// @Description API for creating on demand order
// @Tags order
// @Accept  json
// @Produce  json
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetListCategory(c *gin.Context) {
	category := &ct.GetListCategoryRequest{}

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
	fmt.Println(page, limit, search, "hand")

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
