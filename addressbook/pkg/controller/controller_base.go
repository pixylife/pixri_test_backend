package controller

import (
	"addressbook/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func BaseController(env *config.Env) {
	db = env.DB
}

