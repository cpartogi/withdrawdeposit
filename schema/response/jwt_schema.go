package response

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// JWTVerifyEmail is
type JWTVerifyEmail struct {
	Email    string `json:"email"`
	UserType int64  `json:"user_type"`
	jwt.StandardClaims
}

// JWTToken is
type JWTToken struct {
	UserID       uuid.UUID `json:"user_id"`
	OldUserID    int64     `json:"old_user_id"`
	Email        string    `json:"email"`
	Msisdn       string    `json:"msisdn"`
	UserStatusID int32     `json:"user_status_id"`
	UserType     int64     `json:"user_type"`
	Name         string    `json:"name"`
	jwt.StandardClaims
}

// JWTRefreshToken is
type JWTRefreshToken struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}
