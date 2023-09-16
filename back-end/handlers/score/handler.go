package score

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"backend/db"
	"backend/routes/middlewares"
	"backend/utils"

	"github.com/go-playground/validator/v10"
)

type CreateScoreRequest struct {
	HorseID      int `json:"horse_id" validate:"required"`
	JumpHeightID int `json:"jump_height_id" validate:"required"`
	RankID       int `json:"rank_id" validate:"required"`
}

// create a score handler
func (h *ScoreHandler) CreateScore(w http.ResponseWriter, r *http.Request) {

	var scorePayload CreateScoreRequest
	err := json.NewDecoder(r.Body).Decode(&scorePayload)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	// Validate request
	validate := validator.New()
	err = validate.Struct(scorePayload)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	rank, err := h.DB.GetRankById(scorePayload.RankID)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	jump, err := h.DB.GetJumpHeightById(scorePayload.JumpHeightID)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	score := db.Score{
		HorseID:      scorePayload.HorseID,
		JumpHeightID: scorePayload.JumpHeightID,
		UserID:       middlewares.GetUserID(r),
		RankID:       scorePayload.RankID,
		PointsEarned: rank.Point * jump.Multiplier,
	}

	// insert score into database
	id, err := h.DB.InsertScore(&score)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.SuccessResponse(w, map[string]int{"id": id})
}

// create a score dashboard on query-params
func (h *ScoreHandler) GetScore(w http.ResponseWriter, r *http.Request) {
	// get score id from url
	page := r.URL.Query().Get("page")
	pageInt, _ := strconv.Atoi(page)

	pageSize := r.URL.Query().Get("pageSize")
	pageSizeInt, _ := strconv.Atoi(pageSize)

	rank := r.URL.Query().Get("rank")
	rankInt, _ := strconv.Atoi(rank)

	jump := r.URL.Query().Get("jump")
	jumpInt, _ := strconv.Atoi(jump)

	horse := r.URL.Query().Get("horse")
	horseInt, _ := strconv.Atoi(horse)

	userID := middlewares.GetUserID(r)

	// get score from database
	score, err := h.DB.GetAllScores(
		pageInt, pageSizeInt, horseInt, jumpInt, userID, rankInt,
	)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.SuccessResponse(w, score)
}

func (h *ScoreHandler) LeaderBoard(w http.ResponseWriter, r *http.Request) {
	res, err := h.DB.GetUsersByPointsEarned(
		time.Now().Add(-time.Hour*24*7), time.Now(),
	)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.SuccessResponse(w, res)
}

type ScoreHandler struct {
	DB *db.DB
}
