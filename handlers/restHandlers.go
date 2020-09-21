package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/sharmarajdaksh/basic-auth-microservice/db"
	"github.com/sharmarajdaksh/basic-auth-microservice/models"
	"github.com/sharmarajdaksh/basic-auth-microservice/utils"
)

// Verify an email and password combination
func Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondNotImplemented(w)
		return
	}

	var u registerInput
	if e := json.NewDecoder(r.Body).Decode(&u); e != nil {
		respondBadRequest(w)
		return
	}

	var checkUser models.User
	e := db.DB.Where(&models.User{Email: u.Email}).First(&checkUser)

	responseMessage := "Success"
	if errors.Is(e.Error, gorm.ErrRecordNotFound) {
		respondInvalidAuthenticationAttempt(w)
		return
	}

	isPasswordValid, err := utils.VerifyHash(u.Password, checkUser.PasswordSalt, checkUser.Password)
	if err != nil {
		respondInternalServerError(w)
		return
	}

	if !isPasswordValid {
		respondInvalidAuthenticationAttempt(w)
		return
	}

	respondSuccess(w, responseMessage)
}

// Register an email and password combination
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondNotImplemented(w)
		return
	}

	var u registerInput
	if e := json.NewDecoder(r.Body).Decode(&u); e != nil {
		respondBadRequest(w)
		return
	}

	if !utils.IsPasswordStrong(u.Password) {
		respondWeakPassword(w)
		return
	}

	var checkUser models.User
	e := db.DB.Where(&models.User{Email: u.Email}).First(&checkUser)

	responseMessage := "User registration failed. User with given email already exists"
	if errors.Is(e.Error, gorm.ErrRecordNotFound) {
		var err error
		checkUser, err = registerUser(u)
		if err != nil {
			respondInternalServerError(w)
			return
		}

		responseMessage = "User registration successful"
		w.WriteHeader(http.StatusCreated)
	}

	respondSuccess(w, responseMessage)
}

// Delete will flag a record as deleted in the database
func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondNotImplemented(w)
		return
	}

	var u registerInput
	if e := json.NewDecoder(r.Body).Decode(&u); e != nil {
		respondBadRequest(w)
		return
	}

	var checkUser models.User
	e := db.DB.Where(&models.User{Email: u.Email}).Delete(&checkUser)
	if e.Error != nil {
		respondInternalServerError(w)
		return
	}

	if e.RowsAffected == 0 {
		respondUserNotFound(w)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	respondSuccess(w, "User deleted")
}
