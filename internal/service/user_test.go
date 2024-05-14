package service_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/zhikariz/go-commerce/internal/entity"
	"github.com/zhikariz/go-commerce/internal/service"
	mock_encrypt "github.com/zhikariz/go-commerce/test/mock/pkg/encrypt"
	mock_token "github.com/zhikariz/go-commerce/test/mock/pkg/token"
	mock_repository "github.com/zhikariz/go-commerce/test/mock/repository"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

type UserTestSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	repo        *mock_repository.MockUserRepository
	token       *mock_token.MockTokenUseCase
	encryptTool *mock_encrypt.MockEncryptTool
	userService service.UserService
}

func (s *UserTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_repository.NewMockUserRepository(s.ctrl)
	s.token = mock_token.NewMockTokenUseCase(s.ctrl)
	s.encryptTool = mock_encrypt.NewMockEncryptTool(s.ctrl)
	s.userService = service.NewUserService(s.repo, s.token, s.encryptTool)
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) TestLogin() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	user := new(entity.User)
	user.Password = string(hashedPassword)
	s.Run("success login", func() {
		s.repo.EXPECT().FindUserByEmail("email").Return(user, nil)
		s.encryptTool.EXPECT().Decrypt(user.Alamat).Return("alamat", nil)
		s.encryptTool.EXPECT().Decrypt(user.NoHp).Return("no_hp", nil)
		s.token.EXPECT().GenerateAccessToken(gomock.Any()).Return("token", nil)
		result, err := s.userService.Login("email", "password")
		s.Nil(err)
		s.NotEmpty(result)
	})
}
