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
// @Router        /v1/user/create [post]
// @Summary       Create user
// @Description   API for creating user
// @Tags          user
// @Accept        json
// @Produce       json
// @Param order   body     user_service.CreateUs true "user"
// @Success 200   {object} user_service.CreateUs
// @Failure 404   {object} models.ResponseError
// @Failure 500   {object} models.ResponseError
func (h *handler) CreateUser(c *gin.Context) {
	var (
		req  user_service.CreateUs
		resp *user_service.Us
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

	if err := helpers.ValidateEmailAddress(req.Gmail); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "erroe while validating email"+req.Gmail)
	}

	System_user := h.grpcClient.SystemUserService()
	if System_user == nil {
		handleGrpcErrWithDescription(c, h.log, err, "system service client not initialized")
		return
	}

	resp, err = System_user.Create(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "failed to create customer")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /v1/user/getlist [GET]
// @Summary        Get List user
// @Description    API for getting list user
// @Tags           user
// @Accept         json
// @Produce        json
// @Param		   system_user query string false "system_users"
// @Param		   page query int false "page"
// @Param		   limit query int false "limit"
// @Success 200    {object} user_service.GetListUsResponse
// @Failure 404    {object} models.ResponseError
// @Failure 500    {object} models.ResponseError
func (h *handler) GetListUser(c *gin.Context) {
	var (
		req  user_service.GetListUsRequest
		resp *user_service.GetListUsResponse
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

	System_user := h.grpcClient.SystemUserService()
	if System_user == nil {
		handleGrpcErrWithDescription(c, h.log, err, "system service client not initialized")
		return
	}

	resp, err = System_user.GetList(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router         /v1/user/getbyid/{id} [GET]
// @Summary        Get a single user by ID
// @Description    API for getting a single user by ID
// @Tags           user
// @Accept         json
// @Produce        json
// @Param          id path string true "user ID"
// @Success        200 {object} models.ResponseSuccess
// @Failure        404 {object} models.ResponseError
// @Failure        500 {object} models.ResponseError
func (h *handler) GetUserByID(c *gin.Context) {
	var (
		id   = c.Param("id")
		resp *user_service.Us
		err  error
	)

	req := &user_service.UsPrimaryKey{
		Id: id,
	}

	System_user := h.grpcClient.SystemUserService()
	if System_user == nil {
		handleGrpcErrWithDescription(c, h.log, err, "system service client not initialized")
		return
	}

	resp, err = System_user.GetByID(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router          /v1/user/update/{id} [PUT]
// @Summary         Update a user by ID
// @Description     API for updating a user by ID
// @Tags            user
// @Accept          json
// @Produce         json
// @Param           id path string true "user ID"
// @Param           category body user_service.UpdateUs true "user"
// @Success         200 {object} user_service.Us
// @Failure         404 {object} models.ResponseError
// @Failure         500 {object} models.ResponseError
func (h *handler) UpdateUser(c *gin.Context) {
	var (
		id   = c.Param("id")
		req  user_service.UpdateUs
		resp *user_service.Us
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

	if err := helpers.ValidateEmailAddress(req.Gmail); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "erroe while validating email"+req.Gmail)
	}

	System_user := h.grpcClient.SystemUserService()
	if System_user == nil {
		handleGrpcErrWithDescription(c, h.log, err, "system service client not initialized")
		return
	}

	req.Id = id
	resp, err = System_user.Update(c.Request.Context(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router        /v1/user/delete/{id} [DELETE]
// @Summary       Delete a user by ID
// @Description   API for deleting a user by ID
// @Tags          user
// @Accept        json
// @Produce       json
// @Param         id path string true "user ID"
// @Success       200 {object} models.ResponseError
// @Failure       404 {object} models.ResponseError
// @Failure       500 {object} models.ResponseError
func (h *handler) DeleteUser(c *gin.Context) {
	var (
		id   = c.Param("id")
		err  error
		resp *emptypb.Empty
	)

	req := &user_service.UsPrimaryKey{
		Id: id,
	}

	System_user := h.grpcClient.SystemUserService()
	if System_user == nil {
		handleGrpcErrWithDescription(c, h.log, err, "system service client not initialized")
		return
	}

	resp, err = System_user.Delete(c.Request.Context(), req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}
