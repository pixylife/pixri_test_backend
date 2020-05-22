package main

import (
	"addressbook/config"
	"addressbook/pkg/controller"
	"addressbook/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)
type Property struct {
	Db string
	User string
	Password string
	Url string
	Port string
}

func main(){

	content, _ := ioutil.ReadFile("properties.yml")
	p := Property{}
	err := yaml.Unmarshal([]byte(content), &p)
	db, err := gorm.Open("mysql", p.User+":"+p.Password+"@tcp("+p.Url+":"+p.Port+")/"+p.Db+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	model.InitModels(db)
	env := &config.Env{DB: db}


	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())


	r := e.Group("/")
	controller.BaseController(env)
	controller.ContactController(r, "api")

	e.Logger.Fatal(e.Start(":1235"))
}