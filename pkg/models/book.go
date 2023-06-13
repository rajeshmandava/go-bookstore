package models

import(
	"github.com/jinzhu/gorm"
	"github.com/rajeshmandava/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.model
	Name string `gorm:""json:"name"`
}
