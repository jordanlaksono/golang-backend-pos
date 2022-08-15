package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceRole struct {
	role entitys.EntityRole
}

func NewServiceRole(role entitys.EntityRole) *serviceRole {
	return &serviceRole{role: role}
}

func (s *serviceRole) EntityInsert(input *schemas.SchemaRole) (*models.Role, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaRole
	schema.Role_name = input.Role_name

	res, err := s.role.EntityInsert(&schema)
	return res, err
}
