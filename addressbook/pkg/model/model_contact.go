
package model

import (
"github.com/jinzhu/gorm"
)

type Contact struct {
	Model
	FirstName string `json:"firstName"`
	LastName   string `json:"lastName"`
	PhoneNo string `json:"phoneNo"`
	Address string `json:"address"`
	Email string `json:"email"`

}

func dbpreloadContact(db *gorm.DB) *gorm.DB {
	return db.Preload("Contact")
}
func AddContact(db *gorm.DB, contact *Contact) {
	db.Create(&contact)
}
func UpdateContact(db *gorm.DB, contact *Contact) *Contact {
	db.Save(&contact)
	return contact
}
func DeleteContact(db *gorm.DB, contact *Contact) *Contact {
	db.Delete(&contact)
	return contact
}
func FindContact(db *gorm.DB, id int) *Contact {
	contact := Contact{}
	dbpreloadContact(db).First(&contact, id)
	return &contact
}
func FindAllContact(db *gorm.DB) []*Contact {
	contact := []*Contact{}
	dbpreloadContact(db).Find(&contact)
	return contact
}
