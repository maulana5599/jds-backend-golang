package test

import (
	"jdsapp/entity"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestPayloadPositif(t *testing.T) {
	payload := payloadDataPositif()

	validate := validator.New()
	errValidate := validate.Struct(&payload)

	assert.Nil(t, errValidate)
}

func TestPayloadNegatif(t *testing.T) {
	payload := payloadDataNegatif()

	validate := validator.New()
	errValidate := validate.Struct(&payload)

	assert.Nil(t, errValidate)
}

// func TestLogin(t *testing.T) {

// 	payload := payloadDataPositif()
// 	requestPayload := new(entity.RequestAuth)
// 	requestPayload.Nik = payload.Nik
// 	requestPayload.Password = payload.Password

// 	check, _ := repository.Authentication(requestPayload.Nik, requestPayload.Password)

// 	assert.True(t, check)
// }

func payloadDataPositif() entity.RequestAuth {
	var userPayload entity.RequestAuth
	userPayload.Nik = "3272020505990021"
	userPayload.Password = "maulana186"

	return userPayload
}

func payloadDataNegatif() entity.RequestAuth {
	var userPayload entity.RequestAuth
	userPayload.Nik = "1"
	userPayload.Password = "maulana186"

	return userPayload
}
