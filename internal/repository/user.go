package repository

import (
	"encoding/json"
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
	db        *gorm.DB
	cacheable cache.Cacheable
}

func NewUserRepository(db *gorm.DB, cacheable cache.Cacheable) *userRepository {
	return &userRepository{db: db, cacheable: cacheable}
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
	users := make([]entity.User, 0)

	key := "FindAllUsers"

	data, _ := r.cacheable.Get(key)
	if data == "" {
		if err := r.db.Find(&users).Error; err != nil {
			return users, err
		}
		marshalledUsers, _ := json.Marshal(users)
		err := r.cacheable.Set(key, marshalledUsers, 5*time.Minute)
		if err != nil {
			return users, err
		}
	} else {
		// Data ditemukan di Redis, unmarshal data ke users
		err := json.Unmarshal([]byte(data), &users)
		if err != nil {
			return users, err
		}
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
