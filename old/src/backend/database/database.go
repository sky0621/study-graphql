package database

type Table interface {
	IsTable() bool
	TableName() string
}
