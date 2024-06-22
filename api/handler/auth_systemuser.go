package handler

import (
	as "customer-api-gateway/genproto/auth_service"
	"customer-api-gateway/pkg/validator"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router		/v1/systemuser/register [post]
// @Summary		Register a systemuser
// @Description	This api register a systemuser
// @Tags		SystemUser
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserGmailCheckRequest true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) RegisterSystemUser(c *gin.Context) {
	req := &as.SystemUserGmailCheckRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	if !validator.ValidateGmail(req.Gmail) {
		handleGrpcErrWithDescription(c, h.log, errors.New("wrong gmail"), "error while validating gmail")
		return
	}

	resp, err := h.grpcClient.AuthSystemUserService().SystemUserRegisterByMail(c.Request.Context(), req)
	if err != nil {
		fmt.Println("error while register systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while register systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/registercomfirm [post]
// @Summary		Register Confirm a systemuser
// @Description	This api register confirm a systemuser
// @Tags		SystemUser
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserRConfirm true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) RegisterConfirmSystemUser(c *gin.Context) {
	req := &as.SystemUserRConfirm{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.AuthSystemUserService().SystemUserRegisterByMailConfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Println("error while register confirm systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while register confirm systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/login [post]
// @Summary		Login  a systemuser
// @Description	This api login a systemuser
// @Tags		SystemUser
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserGmailCheckRequest true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) LoginSystemUser(c *gin.Context) {
	req := &as.SystemUserGmailCheckRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.AuthSystemUserService().SystemUserLoginByGmail(c.Request.Context(), req)
	if err != nil {
		fmt.Println("error while login systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/loginconfirm [post]
// @Summary		Login confirm a systemuser
// @Description	This api login confirm a systemuser
// @Tags		SystemUser
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.LoginByGmailRequest true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) SystemUserLoginConfirm(c *gin.Context) {
	req := &as.LoginByGmailRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.AuthSystemUserService().SystemUserLoginByGmailComfirm(c.Request.Context(), (*as.SystemUserLoginByGmailRequest)(req))
	if err != nil {
		fmt.Println("error while login confirm systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login confirm systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/resetpassword [put]
// @Summary		Reset a password a systemuser
// @Description	This api reset a password a systemuser
// @Tags		SystemUser
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserGmailCheckRequest true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) ResetPasswordSystemUser(c *gin.Context) {
	req := &as.SystemUserGmailCheckRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.AuthSystemUserService().SystemUserResetPassword(c.Request.Context(), req)
	if err != nil {
		fmt.Println("error while reset password systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/resetconfirm [put]
// @Summary		Confirm reset password a systemuser
// @Description	This api reset password confirm a systemuser
// @Tags		SystemUser
// @Accept		json
// @Produce		json
// @Param		systemuser body auth_service.SystemUserPasswordConfirm true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) SystemUserResetPasswordConfirm(c *gin.Context) {
	req := &as.SystemUserPasswordConfirm{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.AuthSystemUserService().SystemUserResetPasswordConfirm(c.Request.Context(), req)
	if err != nil {
		fmt.Println("error while reset password confirm systemuser", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while reset password confirm systemuser")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router		/v1/systemuser/loginpassword [post]
// @Summary		Login  a systemuser
// @Description	This api login a systemuser
// @Tags		SystemUser
// @Accept		json
// @Produce		json
// @Param		customer body auth_service.SystemUserLoginRequest true "systemuser"
// @Success		200  {object}  models.ResponseSuccess
// @Failure		400  {object}  models.ResponseError
// @Failure		404  {object}  models.ResponseError
// @Failure		500  {object}  models.ResponseError
func (h *handler) LoginPassworsystem(c *gin.Context) {
	req := &as.SystemUserLoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.AuthSystemUserService().SystemUserLoginByPassword(c.Request.Context(), req)
	if err != nil {
		fmt.Println("error while login customer", err)
		handleGrpcErrWithDescription(c, h.log, err, "error while login customer")
		return
	}

	c.JSON(http.StatusOK, resp)
}
