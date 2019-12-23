package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID   string `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

func (u *User) TableName() string {
	return "user"
}

type UserDao interface {
	InsertOne(u *User) error
	FindAll() ([]*User, error)
	FindOne(id string) (*User, error)
	FindByTodoID(todoID string) (*User, error)
}

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}

func (d *userDao) InsertOne(u *User) error {
	fmt.Println(d.db.NewRecord(u))
	res := d.db.Create(u)
	if err := res.Error; err != nil {
		return err
	}
	fmt.Println(d.db.NewRecord(u))
	return nil
}

func (d *userDao) FindAll() ([]*User, error) {
	return nil, nil
}

func (d *userDao) FindOne(id string) (*User, error) {
	return nil, nil
}

func (d *userDao) FindByTodoID(todoID string) (*User, error) {
	return nil, nil
}
