package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"dadandev.com/wa-engine/internal/database"
	"dadandev.com/wa-engine/internal/dto"
	"dadandev.com/wa-engine/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var body dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.ResponseErrorWithData(w, http.StatusUnprocessableEntity, errors.Error())
		return
	}
	datab := struct {
		Email string
	}{}
	err := database.DB.GetConnection().Get(&datab, "SELECT email FROM users WHERE email=? LIMIT 1", body.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.ErrorResponse(w, http.StatusNotFound, "User tidak ditemukan!")
		} else {
			utils.ErrorResponse(w, http.StatusNotFound, err.Error())
		}
		return // Stop execution if there's an error
	}

	utils.SuccessResponse(w, body.Password)
}
