package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceMenu struct {
	menu entitys.EntityMenu
}

func NewServiceMenu(menu entitys.EntityMenu) *serviceMenu {
	return &serviceMenu{menu: menu}
}

func (s *serviceMenu) EntityInsert(input *schemas.SchemaMenu) (*models.Menu, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaMenu
	schema.Menu_nama = input.Menu_nama
	schema.Menu_link = input.Menu_link

	res, err := s.menu.EntityInsert(&schema)
	return res, err
}
