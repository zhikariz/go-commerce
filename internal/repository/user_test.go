package repository_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/zhikariz/go-commerce/internal/entity"
	"github.com/zhikariz/go-commerce/internal/repository"
	mock_cache "github.com/zhikariz/go-commerce/test/mock/pkg/cache"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type UserTestSuite struct {
	suite.Suite
	ctrl      *gomock.Controller
	db        *gorm.DB
	mock      sqlmock.Sqlmock
	repo      repository.UserRepository
	cacheable *mock_cache.MockCacheable
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) BeforeTest(string, string) {
	s.ctrl = gomock.NewController(s.T())
	db, mock, err := sqlmock.New()
	if err != nil {
		s.FailNow("error mocking db : ", err)
	}

	s.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		s.FailNow("error open db : ", err)
	}

	s.mock = mock
	s.cacheable = mock_cache.NewMockCacheable(s.ctrl)
	s.repo = repository.NewUserRepository(s.db, s.cacheable)
}

func (s *UserTestSuite) AfterTest(string, string) {
	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.FailNow("error expectation : ", err)
	}
}

func (s *UserTestSuite) TestFindUserByID() {
	id := uuid.New()
	s.Run("error find user by id", func() {
		s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE users.id = $1 AND "users"."deleted_at" IS NULL LIMIT $2`)).
			WithArgs(id, 1).
			WillReturnError(gorm.ErrRecordNotFound)
		user, err := s.repo.FindUserByID(id)
		s.NotNil(err)
		s.Equal(uuid.Nil, user.ID)
	})
	s.Run("success find user by id", func() {
		s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE users.id = $1 AND "users"."deleted_at" IS NULL LIMIT $2`)).
			WithArgs(id, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
		s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE "transactions"."user_id" = $1 AND "transactions"."deleted_at" IS NULL`)).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(id))
		user, err := s.repo.FindUserByID(id)
		s.Nil(err)
		s.NotNil(user)
	})
}

func (s *UserTestSuite) TestFindAllUser() {
	s.Run("success find all user by database", func() {
		s.cacheable.EXPECT().Get("FindAllUsers").Return("")
		s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`)).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uuid.New()))
		s.cacheable.EXPECT().Set("FindAllUsers", gomock.Any(), 5*time.Minute)
		users, err := s.repo.FindAllUser()
		s.Nil(err)
		s.NotNil(users)
	})

	s.Run("success find all user by cache", func() {
		s.cacheable.EXPECT().Get("FindAllUsers").Return(fmt.Sprintf(`[{"id":"%v"}]`, uuid.New()))
		users, err := s.repo.FindAllUser()
		s.Nil(err)
		s.NotNil(users)
	})
}

func (s *UserTestSuite) TestCreateUser() {
	s.Run("error create user", func() {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users"`)).
			WillReturnError(errors.New("error"))
		s.mock.ExpectRollback()
		user, err := s.repo.CreateUser(&entity.User{})
		s.NotNil(err)
		s.Equal(uuid.Nil, user.ID)
	})
	s.Run("success create user", func() {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users"`)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mock.ExpectCommit()
		user, err := s.repo.CreateUser(&entity.User{})
		s.Nil(err)
		s.NotNil(user)
	})
}
