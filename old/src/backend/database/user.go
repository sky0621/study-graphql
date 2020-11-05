package database

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID   string `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

func (u *User) Columns() string {
	tn := u.TableName()
	return fmt.Sprintf("%s.id, %s.name", tn, tn)
}

func (u *User) TableName() string {
	return "user"
}

// Tableを実装するマーカーインタフェース
func (u *User) IsTable() bool {
	return true
}

type UserDao interface {
	InsertOne(ctx context.Context, u *User) error
	FindAll(ctx context.Context) ([]*User, error)
	FindOne(ctx context.Context, id string) (*User, error)
	FindByTodoID(ctx context.Context, todoID string) (*User, error)
}

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}

func (d *userDao) InsertOne(ctx context.Context, u *User) error {
	res := d.db.Create(u)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (d *userDao) FindAll(ctx context.Context) ([]*User, error) {
	var users []*User
	res := d.db.Find(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (d *userDao) FindOne(ctx context.Context, id string) (*User, error) {
	var users []*User
	res := d.db.Where("id = ?", id).Find(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(users) < 1 {
		return nil, nil
	}
	return users[0], nil
}

func (d *userDao) FindByTodoID(ctx context.Context, todoID string) (*User, error) {
	var users []*User
	res := d.db.Table("user").
		Select("user.*").
		Joins("LEFT JOIN todo ON todo.user_id = user.id").
		Where("todo.id = ?", todoID).
		First(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	if users == nil || len(users) == 0 {
		return nil, nil
	}
	return users[0], nil
}