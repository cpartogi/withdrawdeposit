package store

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	timeNow := time.Now().Unix()
	columns := []string{"id", "name", "email",
		"msisdn", "password", "user_status_id",
		"profile_image", "gender", "user_type_id",
		"is_msisdn_verified", "is_email_verified",
		"subscribe_id", "created_at", "created_by",
		"updated_at", "updated_by", "deleted_at", "deleted_by"}

	arg := entity.CreateUserParams{
		Name:         "test",
		Email:        "test@example.com",
		Msisdn:       sql.NullString{String: "876123123", Valid: true},
		Password:     "testpassword",
		UserStatusID: 1,
		ProfileImage: sql.NullString{String: "test", Valid: true},
		Gender:       entity.GenderF,
		UserTypeID:   1,
		CreatedAt:    sql.NullInt64{Int64: timeNow, Valid: true},
		UpdatedAt:    sql.NullInt64{Int64: timeNow, Valid: true},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectQuery(createUser).
		WithArgs(
			arg.Name,
			arg.Email,
			arg.Msisdn,
			arg.Password,
			arg.UserStatusID,
			arg.ProfileImage,
			arg.Gender,
			arg.UserTypeID,
			arg.CreatedAt,
			arg.CreatedBy,
			arg.UpdatedAt,
			arg.UpdatedBy,
		).WillReturnRows(sqlmock.NewRows(columns).AddRow("88ef9aca-68e3-420e-bf62-854ef58a09c9", "test2", "test@example.com",
		sql.NullString{String: "876123123", Valid: true}, "testpassword", 1, sql.NullString{String: "test", Valid: true}, entity.GenderF,
		1, false, false, nil, timeNow, nil, timeNow, nil, nil, nil))

	a := NewStore(db)
	user, err := a.CreateUser(context.TODO(), arg)
	assert.NoError(t, err)
	require.NotEmpty(t, user)
}
