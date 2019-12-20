package model

import "time"

type Mysql struct {
	Id        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MysqlInterface interface {
	TableName() string
}

type MysqlD struct {
	Id        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}