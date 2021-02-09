package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cpartogi/withdrawdeposit/module/withdraw/mocks"
	"github.com/cpartogi/withdrawdeposit/schema/response"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var errorWithdraw = errors.New("error withdraw")

func TestWithdrawHandlerNewWithdrawHandler(t *testing.T) {
	e := echo.New()
	mockWithdraw := new(mocks.Usecase)
	NewWithdrawHandler(e, mockWithdraw)
}

func TestDepositBalance(t *testing.T) {
	type input struct {
		seller_id string
	}

	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockWithdraw *mocks.Usecase,
		)
	}{
		{
			name: "#1 success deposit balance",
			expectedInput: input{
				seller_id: "sdss-23222",
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				balResponse := response.Balance{}
				balResponse.Balance = 20000
				balResponse.SellerId = "sdfsdf-sdfsdf-ssss"

				mockWithdraw.
					On("DepositBalance", mock.Anything, mock.Anything).
					Return(balResponse, nil)
			},
		},
		{
			name: "#2 internal server error deposit balance",
			expectedInput: input{
				seller_id: "sdss-23222",
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				balResponse := response.Balance{}
				mockWithdraw.
					On("DepositBalance", mock.Anything, mock.Anything).
					Return(balResponse, errorWithdraw)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockWithdraw := new(mocks.Usecase)

			payload := testCase.expectedInput.seller_id

			e := echo.New()

			req, err := http.NewRequest(echo.GET, "/v1/deposit/balance/:seller_id", strings.NewReader(string(payload)))
			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/deposit/balance/")

			testCase.configureMock(
				testCase.expectedInput,
				mockWithdraw,
			)

			handler := WithdrawHandler{
				withdrawUsecase: mockWithdraw,
			}

			err = handler.DepositBalance(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}

}

func TestDepositBalanceLog(t *testing.T) {
	type input struct {
		seller_id string
		date_from string
		date_to   string
	}

	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockWithdraw *mocks.Usecase,
		)
	}{
		{
			name: "#1 success deposit log",
			expectedInput: input{
				seller_id: "sdss-23222",
				date_from: "2021-02-01",
				date_to:   "2021-02-05",
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				balResponse := []response.BalanceLog{}
				mockWithdraw.
					On("DepositBalanceLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(balResponse, nil)
			},
		},
		{
			name: "#2 internal server error deposit log",
			expectedInput: input{
				seller_id: "sdss-23222",
				date_from: "2021-02-01",
				date_to:   "2021-02-05",
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				balResponse := []response.BalanceLog{}
				mockWithdraw.
					On("DepositBalanceLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(balResponse, errorWithdraw)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockWithdraw := new(mocks.Usecase)

			seller_id := testCase.expectedInput.seller_id

			e := echo.New()

			req, err := http.NewRequest(echo.GET, "/v1/deposit/log",
				strings.NewReader(string(seller_id)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/deposit/log")

			testCase.configureMock(
				testCase.expectedInput,
				mockWithdraw,
			)

			handler := WithdrawHandler{
				withdrawUsecase: mockWithdraw,
			}

			err = handler.DepositBalanceLog(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}

func TestDepositRegister(t *testing.T) {
	type input struct {
		req map[string]interface{}
	}

	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockWithdraw *mocks.Usecase,
		)
	}{
		{
			name: "#1 success register deposit",
			expectedInput: input{
				req: map[string]interface{}{
					"amount":    1,
					"seller_id": "asdf-sdfsd",
				},
			},
			expectedOutput: output{nil, http.StatusCreated},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				depResponse := response.DepositRegistration{}
				depResponse.Amount = 1000
				depResponse.SellerId = "asfsdf-sdfsdf"
				mockWithdraw.
					On("DepositRegister", mock.Anything, mock.Anything).
					Return(depResponse, nil)
			},
		},
		{
			name: "#2 unprocessable register deposit",
			expectedInput: input{
				req: map[string]interface{}{
					"amount":    "1",
					"seller_id": "asdf-sdfsd",
				},
			},
			expectedOutput: output{nil, http.StatusUnprocessableEntity},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				depResponse := response.DepositRegistration{}

				mockWithdraw.
					On("DepositRegister", mock.Anything, mock.Anything).
					Return(depResponse, nil)
			},
		},
		{
			name: "#3 bad request register deposit",
			expectedInput: input{
				req: map[string]interface{}{
					"seller_id": "asdf-sdfsd",
				},
			},
			expectedOutput: output{nil, http.StatusBadRequest},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				depResponse := response.DepositRegistration{}

				mockWithdraw.
					On("DepositRegister", mock.Anything, mock.Anything).
					Return(depResponse, nil)
			},
		},
		{
			name: "#4 internal server error register deposit",
			expectedInput: input{
				req: map[string]interface{}{
					"amount":    1,
					"seller_id": "asdf-sdfsd",
				},
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				depResponse := response.DepositRegistration{}

				mockWithdraw.
					On("DepositRegister", mock.Anything, mock.Anything).
					Return(depResponse, errorWithdraw)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockWithdraw := new(mocks.Usecase)

			payload, err := json.Marshal(testCase.expectedInput.req)

			assert.NoError(t, err)

			e := echo.New()

			req, err := http.NewRequest(echo.POST, "/v1/deposit/register",
				strings.NewReader(string(payload)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/deposit/register")

			testCase.configureMock(
				testCase.expectedInput,
				mockWithdraw,
			)

			handler := WithdrawHandler{
				withdrawUsecase: mockWithdraw,
			}

			err = handler.DepositRegister(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}

func TestSellerRegister(t *testing.T) {
	type input struct {
		req map[string]interface{}
	}

	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockWithdraw *mocks.Usecase,
		)
	}{
		{
			name: "#1 success register seller",
			expectedInput: input{
				req: map[string]interface{}{
					"seller_account_name":   "a",
					"seller_account_number": "23423423",
					"seller_bank_code":      "bca",
					"seller_email":          "admin@mail.com",
					"seller_name":           "namanya",
				},
			},
			expectedOutput: output{nil, http.StatusCreated},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				selResponse := response.SellerRegistration{}
				selResponse.SellerAccountName = "a"
				selResponse.SellerAccountNumber = "2342342"
				selResponse.SellerBankCode = "bca"
				selResponse.SellerEmail = "admin@mail.com"
				selResponse.SellerName = "namanya"
				mockWithdraw.
					On("SellerRegister", mock.Anything, mock.Anything).
					Return(selResponse, nil)
			},
		},
		{
			name: "#2 bad request",
			expectedInput: input{
				req: map[string]interface{}{
					"seller_account_name":   "a",
					"seller_account_number": "23423423",
					"seller_bank_code":      "bca",
				},
			},
			expectedOutput: output{nil, http.StatusBadRequest},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				selResponse := response.SellerRegistration{}
				mockWithdraw.
					On("SellerRegister", mock.Anything, mock.Anything).
					Return(selResponse, nil)
			},
		},
		{
			name: "#3 unprocessable entity",
			expectedInput: input{
				req: map[string]interface{}{
					"seller_account_name":   "a",
					"seller_account_number": 23423423,
					"seller_bank_code":      "bca",
					"seller_email":          "admin@mail.com",
					"seller_name":           "namanya",
				},
			},
			expectedOutput: output{nil, http.StatusUnprocessableEntity},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				selResponse := response.SellerRegistration{}
				mockWithdraw.
					On("SellerRegister", mock.Anything, mock.Anything).
					Return(selResponse, nil)
			},
		},
		{
			name: "#4 internal server error",
			expectedInput: input{
				req: map[string]interface{}{
					"seller_account_name":   "a",
					"seller_account_number": "23423423",
					"seller_bank_code":      "bca",
					"seller_email":          "admin@mail.com",
					"seller_name":           "namanya",
				},
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockWithdraw *mocks.Usecase,
			) {
				selResponse := response.SellerRegistration{}
				mockWithdraw.
					On("SellerRegister", mock.Anything, mock.Anything).
					Return(selResponse, errorWithdraw)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockWithdraw := new(mocks.Usecase)

			payload, err := json.Marshal(testCase.expectedInput.req)

			assert.NoError(t, err)

			e := echo.New()

			req, err := http.NewRequest(echo.POST, "/v1/seller/register",
				strings.NewReader(string(payload)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/seller/register")

			testCase.configureMock(
				testCase.expectedInput,
				mockWithdraw,
			)

			handler := WithdrawHandler{
				withdrawUsecase: mockWithdraw,
			}

			err = handler.SellerRegister(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}
