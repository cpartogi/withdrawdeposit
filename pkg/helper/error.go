package helper

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/cpartogi/withdrawdeposit/constant"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var pqErrorMap = map[string]int{
	"unique_violation": http.StatusConflict,
}

// PqError is
func PqError(err error) (int, error) {
	re := regexp.MustCompile("\\((.*?)\\)")
	if err, ok := err.(*pq.Error); ok {
		match := re.FindStringSubmatch(err.Detail)
		// Change Field Name
		switch match[1] {
		case "msisdn":
			match[1] = "phone number"
		}

		switch err.Code.Name() {
		case "unique_violation":
			return pqErrorMap["unique_violation"], fmt.Errorf("%s already exists", match[1])
		}
	}

	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}

var commonErrorMap = map[error]int{
	constant.ErrUserType:                http.StatusBadRequest,
	constant.ErrUserNotFound:            http.StatusNotFound,
	constant.ErrVerficationDataNotFound: http.StatusNotFound,
	constant.ErrWrongOTPCode:            http.StatusBadRequest,
	bcrypt.ErrMismatchedHashAndPassword: http.StatusBadRequest,
	constant.ErrOnHoldOTPInput:          http.StatusBadRequest,
	constant.ErrEmailHasVerified:        http.StatusBadRequest,
	constant.ErrCannotSendEmail:         http.StatusBadRequest,
	constant.ErrOnHoldSendEmail:         http.StatusBadRequest,
	constant.ErrCannotSendOTP:           http.StatusBadRequest,
	constant.ErrExpiredOTP:              http.StatusBadRequest,
	constant.ErrMsisdnHasVerified:       http.StatusBadRequest,
	constant.ErrEmailNotVerified:        http.StatusBadRequest,
	constant.ErrMsisdnNotVerified:       http.StatusBadRequest,
	constant.ErrCannotResetPassword:     http.StatusBadRequest,
	constant.ErrConflict:                http.StatusConflict,
}

// CommonError is
func CommonError(err error) (int, error) {
	switch err {
	case constant.ErrUserType:
		return commonErrorMap[constant.ErrUserType], constant.ErrUserType
	case constant.ErrUserNotFound:
		return commonErrorMap[constant.ErrUserNotFound], constant.ErrUserNotFound
	case constant.ErrVerficationDataNotFound:
		return commonErrorMap[constant.ErrVerficationDataNotFound], constant.ErrVerficationDataNotFound
	case constant.ErrWrongOTPCode:
		return commonErrorMap[constant.ErrWrongOTPCode], constant.ErrWrongOTPCode
	case bcrypt.ErrMismatchedHashAndPassword:
		return commonErrorMap[bcrypt.ErrMismatchedHashAndPassword], constant.ErrWrongPassword
	case constant.ErrOnHoldOTPInput:
		return commonErrorMap[constant.ErrOnHoldOTPInput], constant.ErrOnHoldOTPInput
	case constant.ErrEmailHasVerified:
		return commonErrorMap[constant.ErrEmailHasVerified], constant.ErrEmailHasVerified
	case constant.ErrCannotSendEmail:
		return commonErrorMap[constant.ErrCannotSendEmail], constant.ErrCannotSendEmail
	case constant.ErrOnHoldSendEmail:
		return commonErrorMap[constant.ErrOnHoldSendEmail], constant.ErrOnHoldSendEmail
	case constant.ErrCannotSendOTP:
		return commonErrorMap[constant.ErrCannotSendOTP], constant.ErrCannotSendOTP
	case constant.ErrExpiredOTP:
		return commonErrorMap[constant.ErrExpiredOTP], constant.ErrExpiredOTP
	case constant.ErrMsisdnHasVerified:
		return commonErrorMap[constant.ErrMsisdnHasVerified], constant.ErrMsisdnHasVerified
	case constant.ErrEmailNotVerified:
		return commonErrorMap[constant.ErrEmailNotVerified], constant.ErrEmailNotVerified
	case constant.ErrMsisdnNotVerified:
		return commonErrorMap[constant.ErrMsisdnNotVerified], constant.ErrMsisdnNotVerified
	case constant.ErrCannotResetPassword:
		return commonErrorMap[constant.ErrCannotResetPassword], constant.ErrCannotResetPassword
	case constant.ErrConflict:
		return commonErrorMap[constant.ErrConflict], constant.ErrConflict
	}
	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}
