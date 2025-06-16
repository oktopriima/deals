package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/deals/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/oktopriima/deals/usecase/attendance/dto"
)

// MockAttendanceUsecase mocks the AttendanceUsecase interface
type MockAttendanceUsecase struct {
	mock.Mock
}

func (m *MockAttendanceUsecase) CreateUsecase(req dto.AttendanceRequest, ctx context.Context) (dto.AttendanceResponse, error) {
	args := m.Called(req, ctx)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(dto.AttendanceResponse), args.Error(1)
}

func TestAdminAttendanceHandler_Serve(t *testing.T) {
	e := echo.New()
	validReq := dto.AttendanceRequest{
		UserId:    2,
		Timestamp: time.Now(),
	}
	validReqBody, _ := json.Marshal(validReq)

	t.Run("Bind error returns 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(validReqBody))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUC := new(MockAttendanceUsecase)
		handler := NewAdminAttendanceHandler(mockUC)

		err := handler.Serve(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Validate error returns error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(validReqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUC := new(MockAttendanceUsecase)
		handler := NewAdminAttendanceHandler(mockUC)

		e.Validator = &MockValidator{ValidateFunc: func(i interface{}) error { return errors.New("validation error") }}

		err := handler.Serve(c)
		assert.EqualError(t, err, "validation error")
	})

	t.Run("Usecase error returns 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(validReqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUC := new(MockAttendanceUsecase)
		mockUC.On("CreateUsecase", mock.Anything, mock.Anything).Return(nil, errors.New("usecase error"))

		handler := NewAdminAttendanceHandler(mockUC)
		e.Validator = &MockValidator{ValidateFunc: func(i interface{}) error { return nil }}

		err := handler.Serve(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		var resp map[string]interface{}
		json.NewDecoder(rec.Body).Decode(&resp)
		assert.Equal(t, float64(http.StatusUnprocessableEntity), resp["code"])
		assert.Equal(t, "usecase error", resp["message"])
	})

	t.Run("Success returns 200", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(validReqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedObj := dto.NewAttendanceResponse(&models.Attendance{
			ID:         1,
			UserID:     2,
			Timestamp:  time.Now(),
			DateString: time.Now().Format("2006-01-02"),
			TimeString: time.Now().Format("15:04:05"),
		})
		mockUC := new(MockAttendanceUsecase)
		mockUC.On("CreateUsecase", mock.Anything, mock.Anything).Return(expectedObj, nil)

		handler := NewAdminAttendanceHandler(mockUC)
		e.Validator = &MockValidator{ValidateFunc: func(i interface{}) error { return nil }}

		err := handler.Serve(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var resp map[string]interface{}
		json.NewDecoder(rec.Body).Decode(&resp)
		assert.Equal(t, float64(http.StatusOK), resp["code"])
		assert.Equal(t, "success", resp["message"])
		assert.NotNil(t, resp["data"])
		assert.Equal(t, expectedObj.GetObject().UserId, int64(resp["data"].(map[string]interface{})["user_id"].(float64)))
	})
}
