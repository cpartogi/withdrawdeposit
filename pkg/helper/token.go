package helper

import (
	"time"

	log "go.uber.org/zap"

	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/schema/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// ApplicationName is for JWT Application Name
const ApplicationName = "TradaruAuth"

// JWTSigningMethod is JWT's signing method
var jwtSigningMethod = jwt.SigningMethodHS256

// GenerateToken will generate both access and refresh token
// for current user.
// Access Token will be expired in 15 Minutes
// Refresh Token will be expired in 6 Months
func GenerateToken(user entity.User) (token response.Token, err error) {
	jwtToken, err := GenerateJWT(user)
	if err != nil {
		return
	}

	refreshToken, e := generateRefresh(user)
	if e != nil {
		return
	}

	token = response.Token{
		Token:        jwtToken,
		RefreshToken: refreshToken,
	}

	return
}

// GenerateJWT is
func GenerateJWT(user entity.User) (signedToken string, err error) {
	exp := time.Now().UTC().Add(viper.GetDuration("auth.token_expiry"))
	claims := response.JWTToken{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ApplicationName,
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
		Email:        user.Email,
		Msisdn:       user.Msisdn.String,
		UserID:       user.ID,
		UserStatusID: user.UserStatusID,
		UserType:     user.UserTypeID,
		Name:         user.Name,
		OldUserID:    user.ExternalUserID.Int64,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err = token.SignedString([]byte(viper.GetString("auth.private_key")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

func generateRefresh(user entity.User) (signedToken string, err error) {
	exp := time.Now().UTC().Add(viper.GetDuration("auth.refresh_token_expiry"))
	claims := response.JWTRefreshToken{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ApplicationName,
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err = token.SignedString([]byte(viper.GetString("auth.private_key")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

// GenerateVerifyEmailToken will generate reset password token
// Token will be expired in 15 Minutes
func GenerateVerifyEmailToken(email string, userType int64) (signedToken string, err error) {
	exp := time.Now().UTC().Add(viper.GetDuration("auth.verif_email_expiry"))
	claims := response.JWTVerifyEmail{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ApplicationName,
			ExpiresAt: exp.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
		},
		Email:    email,
		UserType: userType,
	}

	token := jwt.NewWithClaims(
		jwtSigningMethod,
		claims,
	)

	signedToken, err = token.SignedString([]byte(viper.GetString("auth.private_key")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}

// func GetEmailTokenData(token *jwt.Token) (string, error) {

// }

// ExtractClaims is
func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := viper.GetString("auth.private_key")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	log.S().Error("Invalid JWT Token")
	return nil, false

}
