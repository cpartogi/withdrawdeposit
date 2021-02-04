package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/module/auth/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	timeNow := time.Now().UTC().Unix()
	user := entity.User{}
	mockRepo := new(mocks.Repository)
	mockCreateUser := entity.CreateUserParams{
		Name:         "Test",
		Email:        "test@example.com",
		Msisdn:       sql.NullString{String: "876123123", Valid: true},
		Password:     "password",
		UserStatusID: 1,
		ProfileImage: sql.NullString{String: "test/image.png", Valid: true},
		Gender:       entity.GenderM,
		UserTypeID:   1,
		CreatedAt:    sql.NullInt64{Int64: timeNow, Valid: true},
		UpdatedAt:    sql.NullInt64{Int64: timeNow, Valid: true},
	}

	// mockCreateUserVerification := entity.CreateUserVerificationDataParams{
	// 	VerificationCode: entity.VerificationCode{},
	// 	ResetPassword:    entity.ResetPassword{},
	// 	CreatedAt:        sql.NullInt64{Int64: timeNow, Valid: true},
	// 	UpdatedAt:        sql.NullInt64{Int64: timeNow, Valid: true},
	// }

	t.Run("success", func(t *testing.T) {
		tempCreateUser := mockCreateUser
		u := NewAuthUsecase(mockRepo, 2)

		mockRepo.On("RegisterTx", mock.Anything, mock.AnythingOfType("entity.CreateUserParams"), mock.AnythingOfType("entity.CreateUserVerificationDataParams")).
			Return(user, nil).Once()

		_, err := u.RegisterUser(context.TODO(), tempCreateUser)
		fmt.Println(mockCreateUser)
		assert.NoError(t, err)
		assert.Equal(t, mockCreateUser.Email, tempCreateUser.Email)
		mockRepo.AssertExpectations(t)
	})
}
