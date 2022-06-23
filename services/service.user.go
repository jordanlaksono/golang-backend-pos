package services

import (
	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type serviceUser struct {
	user entitys.EntityUser
}

func NewServiceUser(user entitys.EntityUser) *serviceUser {
	return &serviceUser{user: user}
}

func (s *serviceUser) EntityRegister(input *schemas.SchemaUser) (*models.User, schemas.SchemaDatabaseError) {
	var schema schemas.SchemaUser
	schema.User_first_name = input.User_first_name
	schema.User_last_name = input.User_last_name
	schema.User_email = input.User_email
	schema.User_password = input.User_password

	res, err := s.user.EntityRegister(&schema)
	return res, err
}

func (s *serviceUser) EntityLogin(input *schemas.SchemaUser) (*models.User, schemas.SchemaDatabaseError) {

	var schema schemas.SchemaUser

	schema.User_email = input.User_email
	schema.User_password = input.User_password

	res, err := s.user.EntityLogin(&schema)
	return res, err
}
