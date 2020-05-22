package controller

import (
	"addressbook/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	)

func CreateContact(c echo.Context) error {

	contact := model.Contact{}
	if error := c.Bind(&contact); error != nil {
		return error
	}
	model.AddContact(db, &contact)
	return c.JSON(http.StatusOK, contact)
}
func UpdateContact(c echo.Context) error {

	contact := model.Contact{}
	if error := c.Bind(&contact); error != nil {
		return error
	}
	model.UpdateContact(db, &contact)
	return c.JSON(http.StatusOK, contact)
}
func DeleteContact(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	contact := model.FindContact(db, id)
	model.DeleteContact(db, contact)
	return c.JSON(http.StatusOK, contact)
}
func FindContact(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	contact := model.FindContact(db, id)
	return c.JSON(http.StatusOK, contact)
}
func FindAllContact(c echo.Context) error {

	contact := model.FindAllContact(db)
	return c.JSON(http.StatusOK, contact)
}
func ContactController(g *echo.Group, contextRoot string) {

	g.POST(contextRoot+"/contacts", CreateContact)
	g.PUT(contextRoot+"/contacts/:id", UpdateContact)
	g.DELETE(contextRoot+"/contacts/:id", DeleteContact)
	g.GET(contextRoot+"/contacts/:id", FindContact)
	g.GET(contextRoot+"/contacts", FindAllContact)
}
