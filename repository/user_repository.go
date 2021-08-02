package repository

import (
	"frm/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetConnection() *gorm.DB {
	return r.db
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindAll() (*model.Users, error) {
	var users model.Users

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	user := &model.User{}

	if err := r.db.Find(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	if err := r.db.Where("email=?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}
