package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zhikariz/go-commerce/internal/entity"
	"github.com/zhikariz/go-commerce/pkg/cache"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByID(id uuid.UUID) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	FindAllUser() ([]entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User) (bool, error)
}

type userRepository struct {
	db           *gorm.DB
	cacheable    cache.Cacheable
	arcCacheable cache.LRUCacheable
}

func NewUserRepository(db *gorm.DB, cacheable cache.Cacheable, arcCacheable cache.LRUCacheable) *userRepository {
	return &userRepository{db: db, cacheable: cacheable, arcCacheable: arcCacheable}
}

func (r *userRepository) FindUserByID(id uuid.UUID) (*entity.User, error) {
	user := &entity.User{}

	// Lakukan query dan gabungkan tabel users dan transactions
	if err := r.db.Where("users.id = ?", id).
		Preload("Transactions").
		Take(user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindUserByEmail(email string) (*entity.User, error) {
	user := new(entity.User)
	if err := r.db.Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindAllUser() ([]entity.User, error) {
	var users []entity.User
	key := "FindAllUsers"

	err := r.arcCacheable.Get(key, &users)
	if err == nil && len(users) > 0 {
		fmt.Println(err)
		fmt.Println(users)
		return users, nil
	}

	err = r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	err = r.arcCacheable.Set(key, users, 2*time.Second)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	// Use map to store fields to be updated.
	fields := make(map[string]interface{})

	// Update fields only if they are not empty.
	if user.Email != "" {
		fields["email"] = user.Email
	}
	if user.Password != "" {
		fields["password"] = user.Password
	}
	if user.Role != "" {
		fields["role"] = user.Role
	}
	if user.Alamat != "" {
		fields["alamat"] = user.Alamat
	}

	// Update the database in one query.
	if err := r.db.Model(user).Where("id = ?", user.ID).Updates(fields).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(user *entity.User) (bool, error) {
	if err := r.db.Delete(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
