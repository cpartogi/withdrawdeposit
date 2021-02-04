package utils

import "github.com/cpartogi/withdrawdeposit/constant"

// IsUserType returns
func IsUserType(userType int64) bool {
	return userType == constant.TypeUser
}

// IsCompanyType returns
func IsCompanyType(userType int64) bool {
	return userType == constant.TypeCompany
}
