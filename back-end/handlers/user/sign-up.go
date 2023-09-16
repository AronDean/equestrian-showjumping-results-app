package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"backend/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SignupRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	// Validate request
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	// Check if user already exists
	user, err := h.DB.GetUserByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		utils.InternalServerError(w, err)
		return
	}
	if user != nil {
		utils.BadRequestError(w, errors.New("user already exists"))
		return
	}

	// Hash password
	hashedPassword, err := utils.EncryptPassword(req.Password)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// Create user
	user, err = h.DB.InsertUser(req.Email, req.Name, hashedPassword)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// Generate JWT token
	accessToken, err := utils.JwtEncode(
		h.JwtSecret, int(user.ID), time.Now().Add(h.AccessExp),
	)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	refreshToken, err := utils.JwtEncode(
		h.JwtSecret, int(user.ID), time.Now().Add(h.RefreshExp),
	)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// Return token
	res := SignResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *user,
	}
	utils.SuccessResponse(w, res)
}
