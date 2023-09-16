package user

import (
	"net/http"
	"time"

	"backend/db"
	"backend/routes/middlewares"
	"backend/utils"
)

type UserHandler struct {
	JwtSecret  string
	AccessExp  time.Duration
	RefreshExp time.Duration

	DB *db.DB
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := middlewares.GetUserID(r)

	user, err := h.DB.GetUserById(userID)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.SuccessResponse(w, user)
}
