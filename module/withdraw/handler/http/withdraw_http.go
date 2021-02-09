package http

import (
	"github.com/cpartogi/withdrawdeposit/constant"
	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/module/withdraw"
	"github.com/cpartogi/withdrawdeposit/pkg/utils"
	"github.com/labstack/echo/v4"
)

// AuthHandler  represent the httphandler for auth
type WithdrawHandler struct {
	withdrawUsecase withdraw.Usecase
}

// NewAuthHandler will initialize the contact/ resources endpoint
func NewWithdrawHandler(e *echo.Echo, us withdraw.Usecase) {
	handler := &WithdrawHandler{
		withdrawUsecase: us,
	}

	withdrawrouter := e.Group("/v1")
	withdrawrouter.GET("/deposit/balance/:seller_id", handler.DepositBalance)
	withdrawrouter.GET("/deposit/log", handler.DepositBalanceLog)
	withdrawrouter.POST("/deposit/register", handler.DepositRegister)
	withdrawrouter.POST("/seller/register", handler.SellerRegister)
}

// DepositBalance godoc
// @Summary Seller deposit balance
// @Description Seller deposit balance
// @Tags Deposit
// @Accept  json
// @Produce  json
// @Param seller_id path string true "seller id"
// @Success 200 {object} response.SwaggerDepositBalance
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/deposit/balance/{seller_id} [get]
// DepositBalance handles HTTP request for deposit balance
func (h *WithdrawHandler) DepositBalance(c echo.Context) error {

	ctx := c.Request().Context()
	sellerId := c.Param("seller_id")

	bal, err := h.withdrawUsecase.DepositBalance(ctx, sellerId)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, bal)
}

// DepositBalanceLog godoc
// @Summary Seller deposit balance history
// @Description Seller deposit balance history
// @Tags Deposit
// @Accept  json
// @Produce  json
// @Param seller_id query string true "seller id"
// @Param date_from query string true "format YYYY-MM-DD"
// @Param date_to query string true "format YYYY-MM-DD"
// @Success 200 {object} response.SwaggerDepositBalanceLog
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/deposit/log [get]
// DepositBalancelog handles HTTP request for deposit balance history
func (h *WithdrawHandler) DepositBalanceLog(c echo.Context) error {
	ctx := c.Request().Context()

	queryValues := c.Request().URL.Query()
	sellerId := queryValues.Get("seller_id")
	dateFrom := queryValues.Get("date_from")
	dateTo := queryValues.Get("date_to")

	bal, err := h.withdrawUsecase.DepositBalanceLog(ctx, sellerId, dateFrom, dateTo)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, bal)
}

// DepositRegister godoc
// @Summary Register seller deposit
// @Description Register seller deposit
// @Tags Deposit
// @Accept  json
// @Produce  json
// @Param request body request.DepositRegistration true "Request Body"
// @Success 201 {object} response.SwaggerDepositRegister
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/deposit/register [post]
// DepositRegister handles HTTP request for seller deposit registration
func (h *WithdrawHandler) DepositRegister(c echo.Context) error {
	ctx := c.Request().Context()
	req := entity.Balance{}

	//parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, map[string]interface{}{})
	}

	//validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, map[string]interface{}{})
	}

	reg, err := h.withdrawUsecase.DepositRegister(ctx, req)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.CreatedResponse(c, "Succes register seller deposit", reg)
}

// SellerRegister godoc
// @Summary Register seller data
// @Description Register seller data
// @Tags Seller
// @Accept  json
// @Produce  json
// @Param request body entity.Seller true "Request Body"
// @Success 201 {object} response.SwaggerSellerRegister
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/seller/register [post]
// DepositRegister handles HTTP request for seller deposit registration
func (h *WithdrawHandler) SellerRegister(c echo.Context) error {
	ctx := c.Request().Context()
	req := entity.Seller{}

	//parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, map[string]interface{}{})
	}

	//validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, map[string]interface{}{})
	}

	reg, err := h.withdrawUsecase.SellerRegister(ctx, req)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.CreatedResponse(c, "Succes register seller data", reg)
}
