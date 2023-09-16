package rank

import (
	"net/http"

	"backend/db"
	"backend/utils"
)

type RankHandler struct {
	DB *db.DB
}

func (h *RankHandler) GetRanks(w http.ResponseWriter, r *http.Request) {
	ranks, err := h.DB.GetAllRanks()
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.SuccessResponse(w, ranks)
}

func (h *RankHandler) GetJumpHeights(w http.ResponseWriter, r *http.Request) {
	ranks, err := h.DB.GetAllJumpHeights()
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.SuccessResponse(w, ranks)
}
