package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	ID   string `gorm:"column:id;primary_key"`
	Text string `gorm:"column:text"`
	Done bool   `gorm:"column:done"`
}

func (u *Todo) TableName() string {
	return "todo"
}

type TodoDao interface {
	InsertOne(u *Todo) error
	FindAll() ([]*Todo, error)
	FindByUserID(userID string) ([]*Todo, error)
	FindOne(id string) (*Todo, error)
}

type todoDao struct {
	db *gorm.DB
}

func NewTodoDao(db *gorm.DB) TodoDao {
	return &todoDao{db: db}
}

func (d *todoDao) InsertOne(u *Todo) error {
	fmt.Println(d.db.NewRecord(u))
	res := d.db.Create(u)
	if err := res.Error; err != nil {
		return err
	}
	fmt.Println(d.db.NewRecord(u))
	return nil
}
func (d *todoDao) FindAll() ([]*Todo, error) {
	return nil, nil
}

func (d *todoDao) FindOne(id string) (*Todo, error) {
	return nil, nil
}

func (d *todoDao) FindByUserID(userID string) ([]*Todo, error) {
	return nil, nil
}
