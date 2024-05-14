package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"github.com/zhikariz/go-commerce/internal/http/handler"
	mock_service "github.com/zhikariz/go-commerce/test/mock/service"
	"go.uber.org/mock/gomock"
)

type UserTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	userService *mock_service.MockUserService
	userHandler handler.UserHandler
}

func (s *UserTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.userService = mock_service.NewMockUserService(s.ctrl)
	s.userHandler = handler.NewUserHandler(s.userService)
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) TestLogin() {
	s.Run("error ketika di service", func() {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email": "test@example.com", "password": "password"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.userService.EXPECT().Login("test@example.com", "password").Return("", errors.New("error"))
		_ = s.userHandler.Login(c)
		s.Equal(http.StatusInternalServerError, rec.Code)
	})

	s.Run("success login", func() {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email": "test@example.com", "password": "password"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		s.userService.EXPECT().Login("test@example.com", "password").Return("token", nil)
		err := s.userHandler.Login(c)
		s.Equal(http.StatusOK, rec.Code)
		s.Nil(err)
	})
}
