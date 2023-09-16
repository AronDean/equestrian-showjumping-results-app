package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"backend/db"
	"backend/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SigninRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SignResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	User         db.User `json:"user"`
}

func (h *UserHandler) Signin(w http.ResponseWriter, r *http.Request) {
	var req SigninRequest
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
	if err != nil {
		utils.BadRequestError(w, errors.New("invalid email or password"))
		return
	}

	// Compare password
	err = utils.ComparePassword(req.Password, user.HashedPassword)
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
