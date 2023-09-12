package models

type IAuth interface {
	IsExist() bool
	Create() string
	Delete() bool
}
