package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/module/auth/mocks"
	"github.com/cpartogi/withdrawdeposit/schema/request"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterUserErrorParsing(t *testing.T) {
	mockRegister := map[string]interface{}{
		"email":    "test@gmail.com",
		"msisdn":   "897772332",
		"name":     "Test Data",
		"password": "PasswordTest",
	}

	errorParsingEmailMock := mockRegister
	errorParsingEmailMock["email"] = 0

	errorParsingMsisdnMock := mockRegister
	errorParsingMsisdnMock["msisdn"] = 0

	errorParsingNameMock := mockRegister
	errorParsingNameMock["name"] = 0

	errorParsingPasswordMock := mockRegister
	errorParsingPasswordMock["password"] = 0

	testCases := []struct {
		name            string
		reqUserRegister map[string]interface{}
		buildStubs      func(usecase *mocks.Usecase)
		checkResponse   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:            "Email",
			reqUserRegister: errorParsingEmailMock,
			buildStubs: func(usecase *mocks.Usecase) {
				usecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("entity.CreateUserParams")).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			name:            "Msisdn",
			reqUserRegister: errorParsingMsisdnMock,
			buildStubs: func(usecase *mocks.Usecase) {
				usecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("entity.CreateUserParams")).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			name:            "Name",
			reqUserRegister: errorParsingNameMock,
			buildStubs: func(usecase *mocks.Usecase) {
				usecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("entity.CreateUserParams")).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
		{
			name:            "Password",
			reqUserRegister: errorParsingPasswordMock,
			buildStubs: func(usecase *mocks.Usecase) {
				usecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("entity.CreateUserParams")).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			j, err := json.Marshal(tc.reqUserRegister)
			assert.NoError(t, err)

			mockUCase := new(mocks.Usecase)
			tc.buildStubs(mockUCase)

			e := echo.New()
			req, err := http.NewRequest(echo.POST, "/v2/register-user", strings.NewReader(string(j)))
			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v2/register-user")

			handler := AuthHandler{
				authUsecase: mockUCase,
			}

			err = handler.RegisterUser(c)
			require.NoError(t, err)
		})
	}
}

func TestRegisterUser(t *testing.T) {
	var mockUser entity.User
	mockRegister := request.UserRegistration{
		Email:    "test@gmail.com",
		Msisdn:   "897772332",
		Name:     "Test Data",
		Password: "PasswordTest",
	}
	succesMockRegister := mockRegister
	emptyEmailMockRegister := mockRegister
	emptyEmailMockRegister.Email = ""

	emptyNameMockRegister := mockRegister
	emptyNameMockRegister.Name = ""

	testCases := []struct {
		name            string
		reqUserRegister request.UserRegistration
		buildStubs      func(usecase *mocks.Usecase)
		checkResponse   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:            "Created",
			reqUserRegister: succesMockRegister,
			buildStubs: func(usecase *mocks.Usecase) {
				usecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("entity.CreateUserParams")).
					Times(1).
					Return(mockUser, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name:            "EmptyEmail",
			reqUserRegister: emptyEmailMockRegister,
			buildStubs: func(usecase *mocks.Usecase) {
				usecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("entity.CreateUserParams")).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:            "EmptyName",
			reqUserRegister: emptyNameMockRegister,
			buildStubs: func(usecase *mocks.Usecase) {
				usecase.On("RegisterUser", mock.Anything, mock.AnythingOfType("entity.CreateUserParams")).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			j, err := json.Marshal(tc.reqUserRegister)
			assert.NoError(t, err)

			mockUCase := new(mocks.Usecase)
			tc.buildStubs(mockUCase)

			e := echo.New()
			req, err := http.NewRequest(echo.POST, "/v2/register-user", strings.NewReader(string(j)))
			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v2/register-user")

			handler := AuthHandler{
				authUsecase: mockUCase,
			}

			err = handler.RegisterUser(c)
			require.NoError(t, err)
		})
	}

}
