package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceMenuRole struct {
	menurole entitys.EntityMenuRole
}

func NewServiceMenuRole(menurole entitys.EntityMenuRole) *serviceMenuRole {
	return &serviceMenuRole{menurole: menurole}
}

func (s *serviceMenuRole) EntityInsert(input *schemas.SchemaMenuRole) (*models.MenuRole, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaMenuRole
	schema.Menu_id = input.Menu_id
	schema.Role_id = input.Role_id

	res, err := s.menurole.EntityInsert(&schema)
	return res, err
}

func (s *serviceMenuRole) EntityResults() (*[]models.MenuRole, schemas.SchemaDatabaseError) {
	res, err := s.menurole.EntityResults()
	return res, err
}
