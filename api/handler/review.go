package handler

import (
	"customer-api-gateway/genproto/review_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /review/getall/{id} [GET]
// @Summary Get all reviews
// @Description API for getting all reviews
// @Tags review
// @Accept  json
// @Produce  json
// @Param		search query string false "search"
// @Param		page query int false "page"
// @Param		limit query int false "limit"
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) GetAllReview(c *gin.Context) {
	review := &review_service.GetListReviewRequest{}
	id := c.Param("id")

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
	review.ProductId = id
	review.Search = search
	review.Offset = int64(page)
	review.Limit = int64(limit)

	resp, err := h.grpcClient.ReviewService().GetList(c.Request.Context(), review)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating review")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /review/create [POST]
// @Summary Create review
// @Description API for creating reviews
// @Tags review
// @Accept  json
// @Produce  json
// @Param		review body  review_service.CreateReview true "review"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) CreateReview(c *gin.Context) {
	review := &review_service.CreateReview{}
	if err := c.ShouldBindJSON(&review); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while reading body")
		return
	}

	resp, err := h.grpcClient.ReviewService().Create(c.Request.Context(), review)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating review")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /review/delete/{id} [DELETE]
// @Summary Delete review
// @Description API for deleting review
// @Tags review
// @Accept  json
// @Produce  json
// @Param 		id path string true "id"
// @Success		200  {object}  models.ResponseError
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) DeleteReview(c *gin.Context) {
	id := c.Param("id")
	review := &review_service.ReviewPrimaryKey{Id: id}

	resp, err := h.grpcClient.ReviewService().Delete(c.Request.Context(), review)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting review")
		return
	}
	c.JSON(http.StatusOK, resp)
}
