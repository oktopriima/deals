package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/usecase/authentication/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthenticationUsecase mocks the AuthenticationUsecase interface
type MockAuthenticationUsecase struct {
	mock.Mock
}

func (m *MockAuthenticationUsecase) LoginUsecase(req dto.AuthenticationRequest, ctx context.Context) (dto.AuthenticationResponse, error) {
	args := m.Called(req, ctx)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(dto.AuthenticationResponse), args.Error(1)
}

// MockValidator implements echo.Validator for testing
type MockValidator struct {
	ValidateFunc func(i interface{}) error
}

func (m *MockValidator) Validate(i interface{}) error {
	return m.ValidateFunc(i)
}

func TestLoginByEmail_BindError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("{invalid json")))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUC := new(MockAuthenticationUsecase)
	h := NewAuthenticationHandler(mockUC)

	err := h.LoginByEmail(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NoError(t, err)
}

func TestLoginByEmail_ValidateError(t *testing.T) {
	e := echo.New()
	body := dto.AuthenticationRequest{
		Username: "username",
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	e.Validator = &MockValidator{
		ValidateFunc: func(i interface{}) error {
			return errors.New("validation error")
		},
	}

	mockUC := new(MockAuthenticationUsecase)
	h := NewAuthenticationHandler(mockUC)

	err := h.LoginByEmail(c)
	assert.EqualError(t, err, "validation error")
}

func TestLoginByEmail_LoginUsecaseError(t *testing.T) {
	e := echo.New()
	body := dto.AuthenticationRequest{Username: "test@example.com", Password: "password"}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	e.Validator = &MockValidator{
		ValidateFunc: func(i interface{}) error {
			return nil
		},
	}

	mockUC := new(MockAuthenticationUsecase)
	mockUC.On("LoginUsecase", body, mock.Anything).Return(nil, errors.New("login failed"))
	h := NewAuthenticationHandler(mockUC)

	err := h.LoginByEmail(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NoError(t, err)
}

func TestLoginByEmail_Success(t *testing.T) {
	e := echo.New()
	body := dto.AuthenticationRequest{Username: "test@example.com", Password: "password"}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	e.Validator = &MockValidator{
		ValidateFunc: func(i interface{}) error {
			return nil
		},
	}

	expectedOutput := dto.NewAuthenticationResponse("token", "refresh-token", time.Now().Unix())

	mockUC := new(MockAuthenticationUsecase)
	mockUC.On("LoginUsecase", body, mock.Anything).Return(expectedOutput, nil)
	h := NewAuthenticationHandler(mockUC)

	err := h.LoginByEmail(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NoError(t, err)

	var resp map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput.GetObject().AccessToken, resp["data"].(map[string]interface{})["access_token"])
	assert.Equal(t, expectedOutput.GetObject().RefreshToken, resp["data"].(map[string]interface{})["refresh_token"])
	assert.Equal(t, expectedOutput.GetObject().ExpiresIn, int64(resp["data"].(map[string]interface{})["expires_in"].(float64)))
}
