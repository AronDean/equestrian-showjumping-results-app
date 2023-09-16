package user

import (
	"encoding/json"
	"net/http"
	"time"

	"backend/utils"
	"github.com/go-playground/validator/v10"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenRequest
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

	// Verify refresh token
	id, err := utils.JwtDecode(h.JwtSecret, req.RefreshToken)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	// Generate new access token
	accessToken, err := utils.JwtEncode(
		h.JwtSecret, id, time.Now().Add(h.AccessExp),
	)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	// Return new access token
	res := RefreshTokenResponse{
		AccessToken: accessToken,
	}
	utils.SuccessResponse(w, res)
}
