package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/jordanlaksono/golang-backend-pos.git/entitys"
	"github.com/jordanlaksono/golang-backend-pos.git/helpers"
	"github.com/jordanlaksono/golang-backend-pos.git/pkg"
	"github.com/jordanlaksono/golang-backend-pos.git/schemas"
)

type handlerUser struct {
	user entitys.EntityUser
}

func NewHandlerUser(user entitys.EntityUser) *handlerUser {
	return &handlerUser{user: user}
}

/**
* ==================================
* Handler Ping User Status Teritory
*==================================
 */

func (h *handlerUser) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping User", http.StatusOK, nil)
}

/**
* ======================================
* Handler Register New Account Teritory
*======================================-
 */

func (h *handlerUser) HandlerRegister(ctx *gin.Context) {
	var body schemas.SchemaUser
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorUser(ctx, body, "register")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.user.EntityRegister(&body)

	if error.Type == "error_register_01" {
		helpers.APIResponse(ctx, "Email already taken", error.Code, nil)
		return
	}

	if error.Type == "error_register_02" {
		helpers.APIResponse(ctx, "Register new user account failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Register new user account success", http.StatusOK, nil)
}

/**
* =================================
* Handler Login Auth Account Teritory
*==================================
 */

func (h *handlerUser) HandlerLogin(ctx *gin.Context) {
	var body schemas.SchemaUser
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorUser(ctx, body, "login")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.user.EntityLogin(&body)

	if error.Type == "error_login_01" {
		helpers.APIResponse(ctx, "User account is not never registered", error.Code, nil)
		return
	}

	if error.Type == "error_login_02" {
		helpers.APIResponse(ctx, "Email or Password is wrong", error.Code, nil)
		return
	}
	accessToken, errorJwt := pkg.Sign(&schemas.JWtMetaRequest{
		Data:      gin.H{"user_id": res.User_id, "user_email": res.User_email},
		SecretKey: pkg.GodotEnv("JWT_SECRET_KEY"),
		Options:   schemas.JwtMetaOptions{Audience: "majoo", ExpiredAt: 1},
	})

	expiredAt := time.Now().Add(time.Duration(time.Minute) * (24 * 60) * 1).Local()
	parseUser := h.user.EntityRegister(&body)

	if errorJwt != nil {
		helpers.APIResponse(ctx, "Generate access token failed", http.StatusBadRequest, nil)
		return
	}

	helpers.APIResponse(ctx, "Login successfuly", http.StatusOK, gin.H{"accessToken": accessToken, "expiredAt": expiredAt, "parse": parseUser})
}

/**
* ======================================
*  All Validator User Input For Merchant
*=======================================
 */

func ValidatorUser(ctx *gin.Context, input schemas.SchemaUser, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "register" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "User_first_name",
					Message: "User_first_name is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "lowercase",
					Field:   "User_first_name",
					Message: "User_first_name must be lowercase",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "User_last_name",
					Message: "User_last_name is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "lowercase",
					Field:   "User_last_name",
					Message: "User_last_name must be lowercase",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "User_email",
					Message: "Email is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "email",
					Field:   "User_email",
					Message: "Email format is not valid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "password",
					Field:   "User_password",
					Message: "Password is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "gte",
					Field:   "User_password",
					Message: "Password must be greater than equal 8 character",
				},
			},
		}
	}

	if Type == "login" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "User_email",
					Message: "Email is required on body",
				},
				gpc.ErrorMetaConfig{
					Tag:     "email",
					Field:   "User_email",
					Message: "Email format is not valid",
				},
				gpc.ErrorMetaConfig{
					Tag:     "required",
					Field:   "User_password",
					Message: "User_password is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
