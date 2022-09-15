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
	schema.User_username = input.User_username
	schema.User_password = input.User_password

	res, err := s.user.EntityRegister(&schema)
	return res, err
}

func (s *serviceUser) EntityLogin(input *schemas.SchemaUser) (*models.User, schemas.SchemaDatabaseError) {

	var schema schemas.SchemaUser

	schema.User_username = input.User_username
	schema.User_password = input.User_password

	res, err := s.user.EntityLogin(&schema)
	return res, err
}

func (s *serviceUser) EntityResult(input *schemas.SchemaUser) (*models.User, schemas.SchemaDatabaseError) {
	var user schemas.SchemaUser
	user.User_id = input.User_id

	res, err := s.user.EntityResult(&user)
	return res, err
}

func (s *serviceUser) EntityResultMenuByUser(input *schemas.SchemaUser) (*[]models.MenuRole, schemas.SchemaDatabaseError) {
	var user schemas.SchemaUser
	user.User_id = input.User_id

	res, err := s.user.EntityResultMenuByUser(&user)
	return res, err
}

func (s *serviceUser) EntityResults() (*[]models.User, schemas.SchemaDatabaseError) {
	res, err := s.user.EntityResults()
	return res, err
}
