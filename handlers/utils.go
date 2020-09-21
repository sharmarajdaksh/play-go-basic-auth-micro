package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sharmarajdaksh/basic-auth-microservice/config"
	"github.com/sharmarajdaksh/basic-auth-microservice/db"
	"github.com/sharmarajdaksh/basic-auth-microservice/models"
	"github.com/sharmarajdaksh/basic-auth-microservice/utils"
)

func respondSuccess(w http.ResponseWriter, message string) {
	res, _ := json.Marshal(
		messageResponse{
			Message: message,
		},
	)
	w.Write(res)
}

func respondError(w http.ResponseWriter, errorMessage string, errorCode int) {
	res, _ := json.Marshal(
		messageResponse{
			Message: errorMessage,
		},
	)
	http.Error(w, string(res), errorCode)
}

func respondNotImplemented(w http.ResponseWriter) {
	respondError(w, "Unimplemented REST method", http.StatusNotImplemented)
}

func respondInternalServerError(w http.ResponseWriter) {
	respondError(w, "Internal Server Error", http.StatusInternalServerError)
}

func respondBadRequest(w http.ResponseWriter) {
	respondError(w, "Bad request", http.StatusBadRequest)
}

func respondWeakPassword(w http.ResponseWriter) {
	respondError(
		w,
		"Password is too weak. Passwords must have one lowercase character, one uppercase character, one number, and be atleast 8 characters long",
		http.StatusBadRequest,
	)
}

func respondInvalidAuthenticationAttempt(w http.ResponseWriter) {
	respondError(
		w,
		"Invalid authentication attempt",
		http.StatusUnauthorized,
	)
}

func respondUserNotFound(w http.ResponseWriter) {
	respondError(
		w,
		"User not found",
		http.StatusNotFound,
	)
}

func registerUser(u registerInput) (models.User, error) {
	salt, e := utils.GenerateSalt(config.C.Security.SaltSize)
	if e != nil {
		return models.User{}, e
	}

	hash, e := utils.GenerateSaltedHash(u.Password, salt)
	if e != nil {
		return models.User{}, e
	}

	createdUser := models.User{
		Email:        u.Email,
		Password:     hash,
		PasswordSalt: salt,
	}

	if e := db.DB.Create(&createdUser).Error; e != nil {
		return models.User{}, e
	}

	return createdUser, nil
}
