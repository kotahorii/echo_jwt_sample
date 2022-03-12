package models

type User struct {
	Id       int64
	Name     string `sql:"size:60"`
	Password string `sql:"size:60"`
}
