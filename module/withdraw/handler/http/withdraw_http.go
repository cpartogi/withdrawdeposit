package http

import (
	"github.com/cpartogi/withdrawdeposit/constant"
	"github.com/cpartogi/withdrawdeposit/middleware"
	"github.com/cpartogi/withdrawdeposit/module/withdraw"
	"github.com/cpartogi/withdrawdeposit/pkg/utils"
	"github.com/cpartogi/withdrawdeposit/schema/request"
	"github.com/cpartogi/withdrawdeposit/schema/response"
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
	router := e.Group("/v2")

	router.POST("/register-user", handler.RegisterUser)
	router.POST("/register-company", handler.RegisterCompany)
	router.POST("/login-user", handler.LoginUser)
	router.POST("/login-company", handler.LoginCompany)
	router.POST("/login-otp", handler.LoginOTP)
	router.POST("/validate/reset-password", handler.ValidateResetPassword)

	send := router.Group("/send")
	send.POST("/email-verification", handler.SendEmailVerification)
	send.POST("/misscall", handler.SendMisscallOTP)
	send.POST("/otp", handler.SendSMSOTP)
	send.POST("/reset-password", handler.SendResetPassword)

	legacy := router.Group("/legacy")
	legacy.POST("/register-user", handler.RegisterUserLegacy, middleware.ApikeyMiddleware)
	legacy.POST("/register-company", handler.RegisterCompanyLegacy, middleware.ApikeyMiddleware)

	withdrawrouter := e.Group("/v1")
	withdrawrouter.GET("/deposit/balance/:seller_id", handler.DepositBalance)
}

// RegisterUser godoc
// @Summary Register new user account
// @Description Register new user account
// @Tags Register
// @Accept  json
// @Produce  json
// @Param request body request.UserRegistration true "Request Body"
// @Success 201 {object} response.SwaggerUserRegister
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/register-user [post]
// Register handles HTTP request to create new account
func (h *WithdrawHandler) RegisterUser(c echo.Context) error {
	req := request.UserRegistration{}
	res := response.Register{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	createUserParams := req.CreateUserParams()

	user, err := h.withdrawUsecase.RegisterUser(ctx, createUserParams)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(user)

	return utils.CreatedResponse(c, "Success register new user", res)
}

// RegisterCompany godoc
// @Summary Register new company account
// @Description Register new company account
// @Tags Register
// @Accept  json
// @Produce  json
// @Param request body request.CompanyRegistration true "Request Body"
// @Success 201 {object} response.SwaggerCompanyRegister
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/register-company [post]
// Register handles HTTP request to create new account
func (h *WithdrawHandler) RegisterCompany(c echo.Context) error {
	req := request.CompanyRegistration{}
	res := response.Register{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	createUserParams := req.CreateUserParams()

	user, err := h.withdrawUsecase.RegisterCompany(ctx, createUserParams)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(user)

	return utils.CreatedResponse(c, "Success register new company", res)
}

// LoginUser godoc
// @Summary Login for user
// @Description Login for user
// @Tags Login
// @Accept  json
// @Produce  json
// @Param request body request.UserLogin true "Request Body"
// @Success 200 {object} response.SwaggerUserLogin
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/login-user [post]
// LoginUser handles HTTP request for user login
func (h *WithdrawHandler) LoginUser(c echo.Context) error {
	req := request.UserLogin{}
	res := response.Login{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	user := req.User()

	token, err := h.withdrawUsecase.LoginUser(ctx, user)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(token)

	return utils.SuccessResponse(c, constant.SuccessLogin, res)
}

// LoginCompany godoc
// @Summary Login for company
// @Description Login for company
// @Tags Login
// @Accept  json
// @Produce  json
// @Param request body request.CompanyLogin true "Request Body"
// @Success 200 {object} response.SwaggerCompanyLogin
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/login-company [post]
// LoginCompany handles HTTP request for company login
func (h *WithdrawHandler) LoginCompany(c echo.Context) error {
	req := request.CompanyLogin{}
	res := response.Login{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	user := req.User()

	token, err := h.withdrawUsecase.LoginCompany(ctx, user)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(token)

	return utils.SuccessResponse(c, constant.SuccessLogin, res)
}

// LoginOTP godoc
// @Summary Login with otp code for user
// @Description Login with otp code for user
// @Tags Login
// @Accept  json
// @Produce  json
// @Param request body request.OTPLogin true "Request Body"
// @Success 200 {object} response.SwaggerOTPLogin
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/login-otp [post]
// LoginOTP handles HTTP request for login with otp code for user
func (h *WithdrawHandler) LoginOTP(c echo.Context) error {
	req := request.OTPLogin{}
	res := response.Login{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	user := req.User()

	token, err := h.withdrawUsecase.LoginOTP(ctx, user, req.OTP)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(token)

	return utils.SuccessResponse(c, constant.SuccessLogin, res)
}

// SendEmailVerification godoc
// @Summary Send email verification
// @Description Send email verification to user or company
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param request body request.EmailVerification true "Request Body"
// @Success 200 {object} response.SwaggerEmailVerification
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/send/email-verification [post]
// SendEmailVerification handles HTTP request for Send email verification
func (h *WithdrawHandler) SendEmailVerification(c echo.Context) error {
	req := request.EmailVerification{}
	res := response.EmailVerification{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	user, err := h.withdrawUsecase.SendEmailVerification(ctx, req.Email)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(user)

	return utils.SuccessResponse(c, constant.SuccessSendEmail, res)
}

// SendMisscallOTP godoc
// @Summary Send miscall otp
// @Description Send miscall otp for user
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param request body request.MisscallOTP true "Request Body"
// @Success 200 {object} response.SwaggerMisscallOTP
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/send/misscall [post]
// SendMisscallOTP handles HTTP request for Send miscall otp
func (h *WithdrawHandler) SendMisscallOTP(c echo.Context) error {
	req := request.MisscallOTP{}
	res := response.MisscallOTPNumber{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	misscallResponse, err := h.withdrawUsecase.SendMisscallOTP(ctx, req.Msisdn)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(misscallResponse)

	return utils.SuccessResponse(c, constant.SuccesRequestMisscall, res)
}

// SendSMSOTP godoc
// @Summary Send sms otp
// @Description Send sms otp for user
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param request body request.SMSOTP true "Request Body"
// @Success 200 {object} response.SwaggerSMSOTP
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/send/otp [post]
// SendSMSOTP handles HTTP request for Send miscall otp
func (h *WithdrawHandler) SendSMSOTP(c echo.Context) error {
	req := request.SMSOTP{}
	res := response.SMSOTP{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	smsResponse, err := h.withdrawUsecase.SendSMSOTP(ctx, req.Msisdn, req.Type, req.Signature)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(smsResponse, req.Type)

	return utils.SuccessResponse(c, constant.SuccesRequestOTP, res)
}

// SendResetPassword godoc
// @Summary Send reset password otp
// @Description Send reset password otp
// @Tags Reset Password
// @Accept  json
// @Produce  json
// @Param request body request.ResetPassword true "Request Body"
// @Success 200 {object} response.SwaggerResetPassword
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/send/reset-password [post]
// SendResetPassword handles HTTP request for Send reset password otp
func (h *WithdrawHandler) SendResetPassword(c echo.Context) error {
	req := request.ResetPassword{}
	res := response.ResetPassword{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	user, err := h.withdrawUsecase.SendResetPassword(ctx, req.Email)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(&user)

	return utils.SuccessResponse(c, constant.SuccessSendResetPassword, res)
}

// ValidateResetPassword godoc
// @Summary Validate reset password otp
// @Description Validate reset password otp
// @Tags Reset Password
// @Accept  json
// @Produce  json
// @Param request body request.ValidateResetPassword true "Request Body"
// @Success 200 {object} response.SwaggerResetPassword
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/validate/reset-password [post]
// ValidateResetPassword handles HTTP request for validate reset password otp
func (h *WithdrawHandler) ValidateResetPassword(c echo.Context) error {
	req := request.ValidateResetPassword{}
	res := response.ValidateResetPassword{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	user, err := h.withdrawUsecase.ValidateResetPassword(ctx, req.Email, req.OTP)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(&user, req.OTP)

	return utils.SuccessResponse(c, constant.SuccessValidate, res)
}

// RegisterUserLegacy godoc
// @Summary Register new user account (legacy)
// @Description Register new user account  (legacy)
// @Tags Legacy
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Apikey"
// @Param request body request.UserRegistrationLegacy true "Request Body"
// @Success 201 {object} response.SwaggerUserRegister
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/legacy/register-user [post]
// Register handles HTTP request to create new account
func (h *WithdrawHandler) RegisterUserLegacy(c echo.Context) error {
	req := request.UserRegistrationLegacy{}
	res := response.Register{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	createUserLegacyParams := req.CreateUserLegacyParams()

	user, err := h.withdrawUsecase.RegisterUserLegacy(ctx, createUserLegacyParams)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(user)

	return utils.CreatedResponse(c, "Success register new user", res)
}

// RegisterCompanyLegacy godoc
// @Summary Register new company account (legacy)
// @Description Register new company account (legacy)
// @Tags Legacy
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Apikey"
// @Param request body request.CompanyRegistrationLegacy true "Request Body"
// @Success 201 {object} response.SwaggerCompanyRegister
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v2/legacy/register-company [post]
// Register handles HTTP request to create new account
func (h *WithdrawHandler) RegisterCompanyLegacy(c echo.Context) error {
	req := request.CompanyRegistrationLegacy{}
	res := response.Register{}
	ctx := c.Request().Context()

	// parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, res)
	}

	// validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, res)
	}

	createUserLegacyParams := req.CreateUserLegacyParams()

	user, err := h.withdrawUsecase.RegisterCompanyLegacy(ctx, createUserLegacyParams)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	res.Format(user)

	return utils.CreatedResponse(c, "Success register new company", res)
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

	res := response.Balance{}

	bal, err := h.withdrawUsecase.DepositBalance(ctx, sellerId)
	if err != nil {
		return utils.ErrorResponse(c, err, res)
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, bal)
}
