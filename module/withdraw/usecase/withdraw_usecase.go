package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/cpartogi/withdrawdeposit/constant"
	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/module/withdraw"
	"github.com/cpartogi/withdrawdeposit/pkg/helper"
	"github.com/cpartogi/withdrawdeposit/pkg/utils"
	"github.com/cpartogi/withdrawdeposit/schema/request"
	"github.com/cpartogi/withdrawdeposit/schema/response"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// AuthUsecase will create a usecase with its required repo
type WithdrawUsecase struct {
	withdrawRepo   withdraw.Repository
	contextTimeout time.Duration
}

// NewAuthUsecase will create new an contactUsecase object representation of auth.Usecase
func NewWithdrawUsecase(ar withdraw.Repository, timeout time.Duration) withdraw.Usecase {
	return &WithdrawUsecase{
		withdrawRepo:   ar,
		contextTimeout: timeout,
	}
}

func (u *WithdrawUsecase) getUserByMsisdn(ctx context.Context, msisdn sql.NullString) (user entity.User, err error) {
	msisdn.String = utils.Encrypt(msisdn.String, viper.GetString("encrypt.msisdn"))
	user, err = u.withdrawRepo.GetUserByMsisdn(ctx, msisdn)

	return
}

func (u *WithdrawUsecase) getUserByEmail(ctx context.Context, email string) (user entity.User, err error) {
	email = utils.Encrypt(email, viper.GetString("encrypt.email"))
	user, err = u.withdrawRepo.GetUserByEmail(ctx, email)

	return
}

// RegisterUser returns
func (u *WithdrawUsecase) RegisterUser(c context.Context, createUserParams entity.CreateUserParams) (user entity.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	pHash, err := utils.HashPassword(createUserParams.Password)
	if err != nil {
		return
	}
	createUserParams.Password = pHash

	verifEmailCode, err := helper.GenerateVerifyEmailToken(createUserParams.Email, createUserParams.UserTypeID)
	if err != nil {
		return
	}

	// encryption
	createUserParams.Email = utils.Encrypt(createUserParams.Email, viper.GetString("encrypt.email"))
	createUserParams.Msisdn.String = utils.Encrypt(createUserParams.Msisdn.String, viper.GetString("encrypt.msisdn"))

	createUserParams.ProfileImage = sql.NullString{String: u.generateProfileImage(), Valid: true}

	timeNow := time.Now().Unix()
	createUserVerificationParams := entity.CreateUserVerificationDataParams{
		VerificationCode: entity.VerificationCode{},
		EmailToken:       entity.EmailToken{Token: verifEmailCode},
		ResetPassword:    entity.ResetPassword{},
		CreatedAt:        sql.NullInt64{Int64: timeNow, Valid: true},
		UpdatedAt:        sql.NullInt64{Int64: timeNow, Valid: true},
	}

	user, err = u.withdrawRepo.RegisterTx(ctx, createUserParams, createUserVerificationParams)
	if err != nil {
		return
	}

	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))
	user.Msisdn.String = utils.Decrypt(user.Msisdn.String, viper.GetString("encrypt.msisdn"))

	// Send Email
	go u.sendVerification(user, verifEmailCode)

	return
}

// RegisterCompany returns
func (u *WithdrawUsecase) RegisterCompany(c context.Context, createUserParams entity.CreateUserParams) (user entity.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	pHash, err := utils.HashPassword(createUserParams.Password)
	if err != nil {
		return
	}
	createUserParams.Password = pHash

	verifEmailCode, err := helper.GenerateVerifyEmailToken(createUserParams.Email, createUserParams.UserTypeID)
	if err != nil {
		return
	}

	// encryption
	createUserParams.Email = utils.Encrypt(createUserParams.Email, viper.GetString("encrypt.email"))
	createUserParams.Msisdn.String = utils.Encrypt(createUserParams.Msisdn.String, viper.GetString("encrypt.msisdn"))

	createUserParams.ProfileImage = sql.NullString{String: u.generateProfileImage(), Valid: true}

	timeNow := time.Now().Unix()
	createUserVerificationParams := entity.CreateUserVerificationDataParams{
		VerificationCode: entity.VerificationCode{},
		ResetPassword:    entity.ResetPassword{},
		CreatedAt:        sql.NullInt64{Int64: timeNow, Valid: true},
		UpdatedAt:        sql.NullInt64{Int64: timeNow, Valid: true},
	}

	user, err = u.withdrawRepo.RegisterTx(ctx, createUserParams, createUserVerificationParams)
	if err != nil {
		return
	}

	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))
	user.Msisdn.String = utils.Decrypt(user.Msisdn.String, viper.GetString("encrypt.msisdn"))

	// Send Email
	go u.sendVerification(user, verifEmailCode)

	return
}

// LoginUser will
func (u *WithdrawUsecase) LoginUser(c context.Context, eUser entity.User) (token response.Token, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.getUserByMsisdn(ctx, eUser.Msisdn)
	if err != nil {
		return
	}

	if !utils.IsUserType(user.UserTypeID) {
		err = constant.ErrUserType
		return
	}

	if !user.IsMsisdnVerified {
		err = constant.ErrMsisdnNotVerified
		return
	}

	err = utils.CheckPasswordHash(eUser.Password, user.Password)
	if err != nil {
		return
	}

	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))
	user.Msisdn.String = utils.Decrypt(user.Msisdn.String, viper.GetString("encrypt.msisdn"))

	token, err = helper.GenerateToken(user)
	if err != nil {
		return
	}

	return
}

// LoginOTP will
func (u *WithdrawUsecase) LoginOTP(c context.Context, eUser entity.User, otp string) (token response.Token, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.getUserByMsisdn(ctx, eUser.Msisdn)
	if err != nil {
		return
	}

	if !utils.IsUserType(user.UserTypeID) {
		err = constant.ErrUserType
		return
	}

	verificationCode, err := u.withdrawRepo.GetVerificationCodeByUserID(ctx, user.ID)
	if err != nil {
		return
	}

	verificationCode, err = u.checkHoldInputOTP(ctx, verificationCode, user.ID)
	if err != nil {
		return
	}

	if otp != verificationCode.Otp {
		err = u.addFailedOTP(ctx, verificationCode, user.ID)
		if err != nil {
			return
		}
		err = constant.ErrWrongOTPCode
		return
	}

	if ok := u.checkExpiredOtp(ctx, verificationCode, user.ID); !ok {
		verificationCode.Otp = ""
		err = u.addFailedOTP(ctx, verificationCode, user.ID)
		if err != nil {
			return
		}
		err = constant.ErrExpiredOTP
		return
	}

	verificationCode.FailedCount = 0
	verificationCode.HoldUntil = sql.NullInt64{Int64: 0, Valid: false}
	verificationCode.Otp = ""
	updateVerificationCodeParams := entity.UpdateVerificationCodeByUserIDParams{
		UserID:           user.ID,
		VerificationCode: verificationCode,
		UpdatedAt:        sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
		UpdatedBy:        user.ID,
	}

	updateMsisdnParams := entity.UpdateUserToVerifiedStatusParams{
		ID:               user.ID,
		IsMsisdnVerified: true,
		UserStatusID:     constant.StatusVerified,
	}

	user, err = u.withdrawRepo.LoginOTPTx(ctx, updateVerificationCodeParams, updateMsisdnParams)
	if err != nil {
		return
	}

	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))
	user.Msisdn.String = utils.Decrypt(user.Msisdn.String, viper.GetString("encrypt.msisdn"))

	token, err = helper.GenerateToken(user)
	if err != nil {
		return
	}

	return
}

// LoginCompany will
func (u *WithdrawUsecase) LoginCompany(c context.Context, eUser entity.User) (token response.Token, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.getUserByEmail(ctx, eUser.Email)
	if err != nil {
		return
	}

	if !utils.IsCompanyType(user.UserTypeID) {
		err = constant.ErrCompanyType
		return
	}

	if !user.IsEmailVerified {
		err = constant.ErrEmailNotVerified
		return
	}

	err = utils.CheckPasswordHash(eUser.Password, user.Password)
	if err != nil {
		return
	}

	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))
	user.Msisdn.String = utils.Decrypt(user.Msisdn.String, viper.GetString("encrypt.msisdn"))

	token, err = helper.GenerateToken(user)
	if err != nil {
		return
	}

	return
}

// SendEmailVerification will
func (u *WithdrawUsecase) SendEmailVerification(c context.Context, email string) (user entity.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	timeNow := time.Now().UTC()
	timeNowUnix := timeNow.Unix()

	user, err = u.getUserByEmail(ctx, email)
	if err != nil {
		return
	}

	if user.IsEmailVerified {
		err = constant.ErrEmailHasVerified
		return
	}

	emailToken, err := u.withdrawRepo.GetEmailTokenByUserID(ctx, user.ID)
	if err != nil {
		return
	}

	emailToken, err = u.checkHoldEmail(ctx, emailToken, user.ID)
	if err != nil {
		return
	}

	lastRequestDatetime := time.Unix(emailToken.LastRequestedAt.Int64, 0).UTC()
	if lastRequestDatetime.Add(viper.GetDuration("email.throttle_per_request")).After(timeNow) {
		err = constant.ErrCannotSendEmail
		return
	}

	tokenJWT, ok := helper.ExtractClaims(emailToken.Token)
	tokenExp := int64(0)
	if ok {
		tokenExp = int64(tokenJWT["exp"].(float64))
	}

	updateEmailTokenParams := entity.UpdateEmailTokenByUserIDParams{
		UserID: user.ID,
		EmailToken: entity.EmailToken{
			Token:           emailToken.Token,
			LastRequestedAt: sql.NullInt64{Int64: timeNowUnix, Valid: true},
			RequestCount:    emailToken.RequestCount + 1,
		},
		UpdatedAt: sql.NullInt64{Int64: timeNowUnix, Valid: true},
		UpdatedBy: user.ID,
	}

	if updateEmailTokenParams.EmailToken.RequestCount >= viper.GetInt("email.request_limit") {
		updateEmailTokenParams.EmailToken.HoldUntil = sql.NullInt64{
			Int64: timeNow.Add(viper.GetDuration("email.hold_time")).Unix(),
			Valid: true,
		}
	}

	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))

	if timeNowUnix > tokenExp {
		var verifEmailCode string
		verifEmailCode, err = helper.GenerateVerifyEmailToken(user.Email, user.UserTypeID)
		if err != nil {
			return
		}
		updateEmailTokenParams.EmailToken.Token = verifEmailCode
	}

	emailToken, err = u.withdrawRepo.UpdateEmailTokenByUserID(ctx, updateEmailTokenParams)
	if err != nil {
		return
	}

	// Send Email
	go u.sendVerification(user, emailToken.Token)

	return
}

// SendMisscallOTP will
func (u *WithdrawUsecase) SendMisscallOTP(c context.Context, msisdn string) (misscallResponse response.MisscallOTP, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	timeNow := time.Now().UTC()
	timeNowUnix := timeNow.Unix()

	user, err := u.getUserByMsisdn(ctx, sql.NullString{String: msisdn, Valid: true})
	if err != nil {
		return
	}

	if !utils.IsUserType(user.UserTypeID) {
		err = constant.ErrUserType
		return
	}

	verificationCode, err := u.withdrawRepo.GetVerificationCodeByUserID(ctx, user.ID)
	if err != nil {
		return
	}

	verificationCode, err = u.checkHold(ctx, verificationCode, user.ID)
	if err != nil {
		return
	}

	lastRequestDatetime := time.Unix(verificationCode.LastRequestedAt.Int64, 0).UTC()
	if lastRequestDatetime.Add(viper.GetDuration("otp.throttle_per_request")).After(timeNow) {
		err = constant.ErrCannotSendOTP
		return
	}

	user.Msisdn.String = utils.Decrypt(user.Msisdn.String, viper.GetString("encrypt.msisdn"))

	req := request.CitcallMisscall{
		Msisdn:  user.Msisdn.String,
		Gateway: viper.GetString("misscall_otp.gateway"),
	}

	requestByte, err := json.Marshal(req)
	if err != nil {
		return
	}

	test := helper.APICall{
		URL:       viper.GetString("misscall_otp.url"),
		Method:    http.MethodPost,
		Header:    map[string]string{"Authorization": fmt.Sprintf("Apikey %s", viper.GetString("misscall_otp.apikey"))},
		FormParam: string(requestByte),
	}

	res, err := test.Call()
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(res.Body), &misscallResponse)
	if err != nil {
		return
	}

	if misscallResponse.RC != constant.StatusOKCitcall {
		err = constant.ErrCannotSendOTP
		return
	}

	token := misscallResponse.Token[len(misscallResponse.Token)-4:]
	verificationCode.Otp = token
	verificationCode.TokenCreatedAt = sql.NullInt64{Int64: timeNowUnix, Valid: true}

	err = u.addRequestOTP(ctx, verificationCode, user.ID)
	if err != nil {
		return
	}

	return
}

// SendSMSOTP will
func (u *WithdrawUsecase) SendSMSOTP(c context.Context, msisdn, smsType, signature string) (smsResponse response.SMSViro, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	timeNow := time.Now().UTC()
	timeNowUnix := timeNow.Unix()

	user, err := u.getUserByMsisdn(ctx, sql.NullString{String: msisdn, Valid: true})
	if err != nil {
		return
	}

	if !utils.IsUserType(user.UserTypeID) {
		err = constant.ErrUserType
		return
	}

	verificationCode, err := u.withdrawRepo.GetVerificationCodeByUserID(ctx, user.ID)
	if err != nil {
		return
	}

	verificationCode, err = u.checkHold(ctx, verificationCode, user.ID)
	if err != nil {
		return
	}

	lastRequestDatetime := time.Unix(verificationCode.LastRequestedAt.Int64, 0).UTC()
	if lastRequestDatetime.Add(viper.GetDuration("otp.throttle_per_request")).After(timeNow) {
		err = constant.ErrCannotSendOTP
		return
	}

	user.Msisdn.String = utils.Decrypt(user.Msisdn.String, viper.GetString("encrypt.msisdn"))

	otp := verificationCode.Otp
	if ok := u.checkExpiredOtp(ctx, verificationCode, user.ID); !ok {
		otp = fmt.Sprintf("%d", randOTP())
		verificationCode.Otp = otp
		verificationCode.TokenCreatedAt = sql.NullInt64{Int64: timeNowUnix, Valid: true}
	}

	req := request.SmsViro{
		From: viper.GetString("sms_otp.viro.from"),
		To:   user.Msisdn.String,
		Text: fmt.Sprintf("%s %s \n%s", constant.SMSOTPText, otp, signature),
	}

	requestByte, err := json.Marshal(req)
	if err != nil {
		return
	}

	test := helper.APICall{
		URL:       viper.GetString("sms_otp.viro.url"),
		Method:    http.MethodPost,
		Header:    map[string]string{"Authorization": fmt.Sprintf("Basic %s", viper.GetString("sms_otp.viro.auth_key"))},
		FormParam: string(requestByte),
	}

	res, err := test.Call()
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = constant.ErrCannotSendOTP
		return
	}

	err = json.Unmarshal([]byte(res.Body), &smsResponse)
	if err != nil {
		return
	}

	err = u.addRequestOTP(ctx, verificationCode, user.ID)
	if err != nil {
		return
	}

	if smsType == constant.OTPTypeRegister {
		go u.sendOTPEmail(user, verificationCode.Otp)
		return
	}

	if user.IsEmailVerified {
		go u.sendOTPEmail(user, verificationCode.Otp)
	}

	return
}

// SendResetPassword for
func (u *WithdrawUsecase) SendResetPassword(c context.Context, email string) (user entity.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	timeNow := time.Now().UTC()
	timeNowUnix := timeNow.Unix()

	user, err = u.getUserByEmail(ctx, email)
	if err != nil {
		return
	}

	resetPassword, err := u.withdrawRepo.GetResetPasswordByUserID(ctx, user.ID)
	if err != nil {
		return
	}

	resetPassword, err = u.checkHoldResetPassword(ctx, resetPassword, user.ID)
	if err != nil {
		return
	}

	lastRequestDatetime := time.Unix(resetPassword.LastRequestedAt.Int64, 0).UTC()
	if lastRequestDatetime.Add(viper.GetDuration("reset_password.throttle_per_request")).After(timeNow) {
		err = constant.ErrCannotResetPassword
		return
	}

	if resetPassword.IsValidated {
		resetPassword.Otp = ""
		resetPassword.IsValidated = false
		resetPassword.TokenCreatedAt = sql.NullInt64{Int64: 0, Valid: false}
	}

	if ok := u.checkResetPasswordExpiredOtp(ctx, resetPassword, user.ID); !ok {
		otp := fmt.Sprintf("%d", randOTP())
		resetPassword.Otp = otp
		resetPassword.TokenCreatedAt = sql.NullInt64{Int64: timeNowUnix, Valid: true}
	}

	err = u.addRequestResetPassword(ctx, resetPassword, user.ID)
	if err != nil {
		return
	}

	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))

	go u.sendEmailResetPassword(user, resetPassword.Otp)

	return
}

// ValidateResetPassword for
func (u *WithdrawUsecase) ValidateResetPassword(c context.Context, email, otp string) (user entity.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err = u.getUserByEmail(ctx, email)
	if err != nil {
		return
	}

	resetPassword, err := u.withdrawRepo.GetResetPasswordByUserID(ctx, user.ID)
	if err != nil {
		return
	}

	if resetPassword.IsValidated {
		resetPassword.Otp = ""
		resetPassword.IsValidated = false
		resetPassword.TokenCreatedAt = sql.NullInt64{Int64: 0, Valid: false}
		err = u.addFailedResetPasswordOTP(ctx, resetPassword, user.ID)
		if err != nil {
			return
		}
		err = constant.ErrExpiredOTP
		return
	}

	resetPassword, err = u.checkHoldResetPassword(ctx, resetPassword, user.ID)
	if err != nil {
		return
	}

	if otp != resetPassword.Otp {
		err = u.addFailedResetPasswordOTP(ctx, resetPassword, user.ID)
		if err != nil {
			return
		}
		err = constant.ErrWrongOTPCode
		return
	}

	if ok := u.checkResetPasswordExpiredOtp(ctx, resetPassword, user.ID); !ok {
		resetPassword.Otp = ""
		err = u.addFailedResetPasswordOTP(ctx, resetPassword, user.ID)
		if err != nil {
			return
		}
		err = constant.ErrExpiredOTP
		return
	}

	resetPassword.FailedCount = 0
	resetPassword.HoldUntil = sql.NullInt64{Int64: 0, Valid: false}
	resetPassword.IsValidated = true
	updateResetPasswordParams := entity.UpdateResetPasswordByUserIDParams{
		UserID:        user.ID,
		ResetPassword: resetPassword,
		UpdatedAt:     sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
		UpdatedBy:     user.ID,
	}

	_, err = u.withdrawRepo.UpdateResetPasswordByUserID(ctx, updateResetPasswordParams)

	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))

	return

}

// RegisterUserLegacy returns
func (u *WithdrawUsecase) RegisterUserLegacy(c context.Context, createUserLegacyParams entity.CreateUserLegacyParams) (user entity.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	pHash, err := utils.HashPassword(createUserLegacyParams.Password)
	if err != nil {
		return
	}
	createUserLegacyParams.Password = pHash

	// encryption
	createUserLegacyParams.Email = utils.Encrypt(createUserLegacyParams.Email, viper.GetString("encrypt.email"))
	createUserLegacyParams.Msisdn.String = utils.Encrypt(createUserLegacyParams.Msisdn.String, viper.GetString("encrypt.msisdn"))

	createUserLegacyParams.ProfileImage = sql.NullString{String: u.generateProfileImage(), Valid: true}

	timeNow := time.Now().Unix()
	createUserVerificationParams := entity.CreateUserVerificationDataParams{
		VerificationCode: entity.VerificationCode{},
		EmailToken:       entity.EmailToken{},
		ResetPassword:    entity.ResetPassword{},
		CreatedAt:        sql.NullInt64{Int64: timeNow, Valid: true},
		UpdatedAt:        sql.NullInt64{Int64: timeNow, Valid: true},
	}

	user, err = u.withdrawRepo.RegisterLegacyTx(ctx, createUserLegacyParams, createUserVerificationParams)
	if err != nil {
		return
	}

	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))
	user.Msisdn.String = utils.Decrypt(user.Msisdn.String, viper.GetString("encrypt.msisdn"))

	return
}

// RegisterCompanyLegacy returns
func (u *WithdrawUsecase) RegisterCompanyLegacy(c context.Context, createUserLegacyParams entity.CreateUserLegacyParams) (user entity.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	pHash, err := utils.HashPassword(createUserLegacyParams.Password)
	if err != nil {
		return
	}
	createUserLegacyParams.Password = pHash

	// encryption
	createUserLegacyParams.Email = utils.Encrypt(createUserLegacyParams.Email, viper.GetString("encrypt.email"))
	createUserLegacyParams.Msisdn.String = utils.Encrypt(createUserLegacyParams.Msisdn.String, viper.GetString("encrypt.msisdn"))

	createUserLegacyParams.ProfileImage = sql.NullString{String: u.generateProfileImage(), Valid: true}

	timeNow := time.Now().Unix()
	createUserVerificationParams := entity.CreateUserVerificationDataParams{
		VerificationCode: entity.VerificationCode{},
		EmailToken:       entity.EmailToken{},
		ResetPassword:    entity.ResetPassword{},
		CreatedAt:        sql.NullInt64{Int64: timeNow, Valid: true},
		UpdatedAt:        sql.NullInt64{Int64: timeNow, Valid: true},
	}

	user, err = u.withdrawRepo.RegisterLegacyTx(ctx, createUserLegacyParams, createUserVerificationParams)
	if err != nil {
		return
	}

	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))
	user.Msisdn.String = utils.Decrypt(user.Msisdn.String, viper.GetString("encrypt.msisdn"))

	return
}

func (u *WithdrawUsecase) sendOTPEmail(user entity.User, otp string) {
	user.Email = utils.Decrypt(user.Email, viper.GetString("encrypt.email"))

	htmlContent, err := helper.GenerateSendOTP(user.Name, otp)
	if err != nil {
		return
	}

	helper.SendMail(user.Email, constant.SendOTPEmailSubject, htmlContent, constant.SendOTPType)
}

func (u *WithdrawUsecase) sendVerification(user entity.User, verifEmailCode string) {
	htmlContent, err := helper.GenerateVerification(user.Name, fmt.Sprintf("%s/verify/email?token=%s",
		viper.GetString("frontend.base_url"), verifEmailCode))
	if err != nil {
		return
	}

	helper.SendMail(user.Email, constant.VerifEmailSubject, htmlContent, constant.VerifEmailType)
}

func (u *WithdrawUsecase) sendEmailResetPassword(user entity.User, otpCode string) {
	htmlContent, err := helper.GenerateResetPassword(user.Name, otpCode)
	if err != nil {
		return
	}

	helper.SendMail(user.Email, constant.SendResetPasswordSubject, htmlContent, constant.ResetPasswordType)
}

func (u *WithdrawUsecase) generateProfileImage() string {
	min := 1
	max := 6
	rand.Seed(time.Now().Unix())
	picnum := strconv.Itoa(rand.Intn(max-min) + min)
	return "default/bird" + picnum + ".png"
}

func (u *WithdrawUsecase) addFailedOTP(ctx context.Context, verificationCode entity.VerificationCode, userID uuid.UUID) (err error) {
	timeNow := time.Now().UTC()
	holdTime := timeNow.Add(viper.GetDuration("otp.input_hold_time")).Unix()

	verificationCode.FailedCount++
	if verificationCode.FailedCount >= viper.GetInt("otp.input_limit") {
		verificationCode.HoldUntil = sql.NullInt64{Int64: holdTime, Valid: true}
	}

	updateVerificationCodeParams := entity.UpdateVerificationCodeByUserIDParams{
		UserID:           userID,
		VerificationCode: verificationCode,
		UpdatedAt:        sql.NullInt64{Int64: timeNow.Unix(), Valid: true},
		UpdatedBy:        userID,
	}
	_, err = u.withdrawRepo.UpdateVerificationCodeByUserID(ctx, updateVerificationCodeParams)
	return
}

func (u *WithdrawUsecase) checkHold(ctx context.Context, verificationCode entity.VerificationCode, userID uuid.UUID) (entity.VerificationCode, error) {
	timeNow := time.Now().UTC()
	holdTimeUTC := time.Unix(verificationCode.HoldUntil.Int64, 0).UTC()

	if holdTimeUTC.After(timeNow) {
		return verificationCode, constant.ErrOnHoldOTPInput
	}

	if holdTimeUTC.Before(timeNow) && verificationCode.FailedCount >= viper.GetInt("otp.input_limit") {
		return u.resetFailedCount(ctx, verificationCode, userID)
	}

	if holdTimeUTC.Before(timeNow) && verificationCode.RequestCount >= viper.GetInt("otp.request_limit") {
		return u.resetRequestCount(ctx, verificationCode, userID)
	}

	return verificationCode, nil
}

func (u *WithdrawUsecase) checkHoldInputOTP(ctx context.Context, verificationCode entity.VerificationCode, userID uuid.UUID) (entity.VerificationCode, error) {
	timeNow := time.Now().UTC()
	holdTimeUTC := time.Unix(verificationCode.HoldUntil.Int64, 0).UTC()

	if holdTimeUTC.After(timeNow) && verificationCode.FailedCount >= viper.GetInt("otp.input_limit") {
		return verificationCode, constant.ErrOnHoldOTPInput
	}

	if holdTimeUTC.Before(timeNow) && verificationCode.FailedCount >= viper.GetInt("otp.input_limit") {
		return u.resetFailedCount(ctx, verificationCode, userID)
	}

	return verificationCode, nil
}

func (u *WithdrawUsecase) checkExpiredOtp(ctx context.Context, verificationCode entity.VerificationCode, userID uuid.UUID) bool {
	timeNow := time.Now().UTC()
	tokenCreatedAt := time.Unix(verificationCode.TokenCreatedAt.Int64, 0).UTC()

	if len(verificationCode.Otp) == constant.MisscallOTPLength {
		return timeNow.Before(tokenCreatedAt.Add(viper.GetDuration("otp.misscall_otp_expiry")))
	}

	if len(verificationCode.Otp) == constant.SmsOtpLength {
		return timeNow.Before(tokenCreatedAt.Add(viper.GetDuration("otp.sms_otp_expiry")))
	}

	return false
}

func (u *WithdrawUsecase) resetFailedCount(ctx context.Context, verificationCode entity.VerificationCode, userID uuid.UUID) (entity.VerificationCode, error) {
	timeNow := time.Now().UTC()
	verificationCode.FailedCount = 0
	verificationCode.HoldUntil = sql.NullInt64{Int64: 0, Valid: false}
	updateVerificationCodeParams := entity.UpdateVerificationCodeByUserIDParams{
		UserID:           userID,
		VerificationCode: verificationCode,
		UpdatedAt:        sql.NullInt64{Int64: timeNow.Unix(), Valid: true},
		UpdatedBy:        userID,
	}
	verificationCodeData, err := u.withdrawRepo.UpdateVerificationCodeByUserID(ctx, updateVerificationCodeParams)
	return verificationCodeData, err
}

func (u *WithdrawUsecase) checkHoldEmail(ctx context.Context, emailToken entity.EmailToken, userID uuid.UUID) (entity.EmailToken, error) {
	timeNow := time.Now().UTC()
	holdTimeUTC := time.Unix(emailToken.HoldUntil.Int64, 0).UTC()

	if holdTimeUTC.After(timeNow) {
		return emailToken, constant.ErrOnHoldSendEmail
	}

	if holdTimeUTC.Before(timeNow) && emailToken.RequestCount >= viper.GetInt("email.request_limit") {
		return u.resetEmailRequestCount(ctx, emailToken, userID)
	}

	return emailToken, nil
}

func (u *WithdrawUsecase) resetEmailRequestCount(ctx context.Context, emailToken entity.EmailToken, userID uuid.UUID) (entity.EmailToken, error) {
	timeNow := time.Now().UTC()
	emailToken.RequestCount = 0
	emailToken.HoldUntil = sql.NullInt64{Int64: 0, Valid: false}
	updateEmailTokenParams := entity.UpdateEmailTokenByUserIDParams{
		UserID:     userID,
		EmailToken: emailToken,
		UpdatedAt:  sql.NullInt64{Int64: timeNow.Unix(), Valid: true},
		UpdatedBy:  userID,
	}
	emailToken, err := u.withdrawRepo.UpdateEmailTokenByUserID(ctx, updateEmailTokenParams)
	return emailToken, err
}

func (u *WithdrawUsecase) addRequestOTP(ctx context.Context, verificationCode entity.VerificationCode, userID uuid.UUID) (err error) {
	timeNow := time.Now().UTC()
	holdTime := timeNow.Add(viper.GetDuration("otp.request_hold_time")).Unix()

	verificationCode.RequestCount++
	if verificationCode.RequestCount >= viper.GetInt("otp.request_limit") {
		verificationCode.HoldUntil = sql.NullInt64{Int64: holdTime, Valid: true}
	}

	verificationCode.LastRequestedAt = sql.NullInt64{Int64: timeNow.Unix(), Valid: true}
	updateVerificationCodeParams := entity.UpdateVerificationCodeByUserIDParams{
		UserID:           userID,
		VerificationCode: verificationCode,
		UpdatedAt:        sql.NullInt64{Int64: timeNow.Unix(), Valid: true},
		UpdatedBy:        userID,
	}
	_, err = u.withdrawRepo.UpdateVerificationCodeByUserID(ctx, updateVerificationCodeParams)
	return
}

func (u *WithdrawUsecase) resetRequestCount(ctx context.Context, verificationCode entity.VerificationCode, userID uuid.UUID) (entity.VerificationCode, error) {
	timeNow := time.Now().UTC()
	verificationCode.RequestCount = 0
	verificationCode.HoldUntil = sql.NullInt64{Int64: 0, Valid: false}

	updateVerificationCodeParams := entity.UpdateVerificationCodeByUserIDParams{
		UserID:           userID,
		VerificationCode: verificationCode,
		UpdatedAt:        sql.NullInt64{Int64: timeNow.Unix(), Valid: true},
		UpdatedBy:        userID,
	}
	verificationCodeData, err := u.withdrawRepo.UpdateVerificationCodeByUserID(ctx, updateVerificationCodeParams)
	return verificationCodeData, err
}

func randOTP() int64 {
	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 99999
	otp := rand.Intn(max-min+1) + min
	return int64(otp)
}

func (u *WithdrawUsecase) checkHoldResetPassword(ctx context.Context, resetPassword entity.ResetPassword, userID uuid.UUID) (entity.ResetPassword, error) {
	timeNow := time.Now().UTC()
	holdTimeUTC := time.Unix(resetPassword.HoldUntil.Int64, 0).UTC()

	if holdTimeUTC.After(timeNow) {
		return resetPassword, constant.ErrOnHoldOTPInput
	}

	if holdTimeUTC.Before(timeNow) && resetPassword.FailedCount >= viper.GetInt("reset_password.input_limit") {
		return u.resetRequestCountResetPassword(ctx, resetPassword, userID)
	}

	if holdTimeUTC.Before(timeNow) && resetPassword.RequestCount >= viper.GetInt("reset_password.request_limit") {
		return u.resetRequestCountResetPassword(ctx, resetPassword, userID)
	}

	return resetPassword, nil
}

func (u *WithdrawUsecase) resetRequestCountResetPassword(ctx context.Context, resetPassword entity.ResetPassword, userID uuid.UUID) (entity.ResetPassword, error) {
	timeNow := time.Now().UTC()
	resetPassword.RequestCount = 0
	resetPassword.HoldUntil = sql.NullInt64{Int64: 0, Valid: false}

	updateResetPasswordParams := entity.UpdateResetPasswordByUserIDParams{
		UserID:        userID,
		ResetPassword: resetPassword,
		UpdatedAt:     sql.NullInt64{Int64: timeNow.Unix(), Valid: true},
		UpdatedBy:     userID,
	}
	resetPasswordData, err := u.withdrawRepo.UpdateResetPasswordByUserID(ctx, updateResetPasswordParams)
	return resetPasswordData, err
}

func (u *WithdrawUsecase) checkResetPasswordExpiredOtp(ctx context.Context, resetPassword entity.ResetPassword, userID uuid.UUID) bool {
	timeNow := time.Now().UTC()
	tokenCreatedAt := time.Unix(resetPassword.TokenCreatedAt.Int64, 0).UTC()

	return timeNow.Before(tokenCreatedAt.Add(viper.GetDuration("reset_password.expiry")))

}

func (u *WithdrawUsecase) addRequestResetPassword(ctx context.Context, resetPassword entity.ResetPassword, userID uuid.UUID) (err error) {
	timeNow := time.Now().UTC()
	holdTime := timeNow.Add(viper.GetDuration("reset_password.request_hold_time")).Unix()

	resetPassword.RequestCount++
	if resetPassword.RequestCount >= viper.GetInt("reset_password.request_limit") {
		resetPassword.HoldUntil = sql.NullInt64{Int64: holdTime, Valid: true}
	}

	resetPassword.LastRequestedAt = sql.NullInt64{Int64: timeNow.Unix(), Valid: true}
	updateResetPasswordParams := entity.UpdateResetPasswordByUserIDParams{
		UserID:        userID,
		ResetPassword: resetPassword,
		UpdatedAt:     sql.NullInt64{Int64: timeNow.Unix(), Valid: true},
		UpdatedBy:     userID,
	}
	_, err = u.withdrawRepo.UpdateResetPasswordByUserID(ctx, updateResetPasswordParams)
	return
}

func (u *WithdrawUsecase) addFailedResetPasswordOTP(ctx context.Context, resetPassword entity.ResetPassword, userID uuid.UUID) (err error) {
	timeNow := time.Now().UTC()
	holdTime := timeNow.Add(viper.GetDuration("reset_password.input_hold_time")).Unix()

	resetPassword.FailedCount++
	if resetPassword.FailedCount >= viper.GetInt("reset_password.input_limit") {
		resetPassword.HoldUntil = sql.NullInt64{Int64: holdTime, Valid: true}
	}

	updateResetPasswordParams := entity.UpdateResetPasswordByUserIDParams{
		UserID:        userID,
		ResetPassword: resetPassword,
		UpdatedAt:     sql.NullInt64{Int64: timeNow.Unix(), Valid: true},
		UpdatedBy:     userID,
	}
	_, err = u.withdrawRepo.UpdateResetPasswordByUserID(ctx, updateResetPasswordParams)
	return
}

func (u *WithdrawUsecase) DepositBalance(ctx context.Context, seller_id string) (bal response.Balance, err error) {
	resp := response.Balance{}

	deposit, err := u.withdrawRepo.DepositBalance(ctx, seller_id)

	if err != nil {
		return resp, err
	}

	return deposit, err
}
