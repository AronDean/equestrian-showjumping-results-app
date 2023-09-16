package horse

import (
	"encoding/json"
	"net/http"
	"strconv"

	"backend/db"
	"backend/routes/middlewares"
	"backend/utils"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type HorseHandler struct {
	DB *db.DB
}

// create horse handler
func (h *HorseHandler) CreateHorse(w http.ResponseWriter, r *http.Request) {
	var horse db.Horse
	err := json.NewDecoder(r.Body).Decode(&horse)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	// Validate request
	validate := validator.New()
	err = validate.Struct(horse)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	horse.OwnerID = middlewares.GetUserID(r)

	// insert horse into database
	id, err := h.DB.InsertHorse(&horse)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.SuccessResponse(w, map[string]int{"id": id})
}

// get horse handler
func (h *HorseHandler) GetHorse(w http.ResponseWriter, r *http.Request) {
	// get horse id from url
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)

	userID := middlewares.GetUserID(r)

	// get horse from database
	horse, err := h.DB.GetHorseById(idInt, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.BadRequestError(w, err)
			return
		}

		utils.InternalServerError(w, err)
		return
	}

	utils.SuccessResponse(w, horse)
}

// update horse handler
func (h *HorseHandler) UpdateHorse(w http.ResponseWriter, r *http.Request) {
	// get horse id from url
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)

	// get name, color, age, weight from request body
	var horse db.Horse
	err := json.NewDecoder(r.Body).Decode(&horse)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	// Validate request
	validate := validator.New()
	err = validate.Struct(horse)
	if err != nil {
		utils.BadRequestError(w, err)
		return
	}

	userID := middlewares.GetUserID(r)

	// update horse in database
	err = h.DB.UpdateHorse(
		idInt, horse.Name, horse.Color,
		horse.Age, userID,
	)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	res := map[string]int{
		"id": idInt,
	}
	utils.SuccessResponse(w, res)
}

// delete horse handler
func (h *HorseHandler) DeleteHorse(w http.ResponseWriter, r *http.Request) {
	// get horse id from url
	id := chi.URLParam(r, "id")
	idInt, _ := strconv.Atoi(id)

	userID := middlewares.GetUserID(r)

	// delete horse from database
	err := h.DB.DeleteHorse(idInt, userID)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	res := map[string]int{
		"id": idInt,
	}
	utils.SuccessResponse(w, res)
}

// get all horses handler
func (h *HorseHandler) GetAllHorses(w http.ResponseWriter, r *http.Request) {
	// get horses from database
	userID := middlewares.GetUserID(r)

	horses, err := h.DB.GetAllHorses(userID)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}

	utils.SuccessResponse(w, horses)
}
