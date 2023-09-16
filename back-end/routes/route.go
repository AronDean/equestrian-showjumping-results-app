package routes

import (
	"os"
	"strconv"
	"time"

	"backend/db"
	"backend/handlers/horse"
	"backend/handlers/rank"
	"backend/handlers/score"
	"backend/handlers/user"
	"backend/routes/middlewares"
	"github.com/go-chi/chi"
)

func Routes(db *db.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middlewares.CorsMiddleware)

	router.Mount("/api/v1/user", AuthRoutes(db))
	router.Mount("/api/v1/horse", HorseRoutes(db))
	router.Mount("/api/v1/score", ScoreRoutes(db))
	router.Mount("/api/v1/other", OtherRoutes(db))

	return router
}

func AuthRoutes(db *db.DB) *chi.Mux {
	accessExp, _ := strconv.Atoi(os.Getenv("ACCESS_EXP"))
	refreshExp, _ := strconv.Atoi(os.Getenv("REFRESH_EXP"))
	userHandler := user.UserHandler{
		JwtSecret:  os.Getenv("JWT_SECRET"),
		AccessExp:  time.Duration(accessExp) * time.Minute,
		RefreshExp: time.Duration(refreshExp) * time.Hour,
		DB:         db,
	}

	router := chi.NewRouter()
	router.Post("/sign-in", userHandler.Signin)
	router.Post("/sign-up", userHandler.Signup)
	router.Post("/refresh", userHandler.RefreshToken)

	router.With(middlewares.AuthMiddleware).Get("/", userHandler.GetUser)

	return router
}

func HorseRoutes(db *db.DB) *chi.Mux {
	horseHandler := horse.HorseHandler{
		DB: db,
	}
	router := chi.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	router.Get("/", horseHandler.GetAllHorses)
	router.Post("/", horseHandler.CreateHorse)
	router.Get("/{id}", horseHandler.GetHorse)
	router.Put("/{id}", horseHandler.UpdateHorse)
	router.Delete("/{id}", horseHandler.DeleteHorse)

	return router
}

func ScoreRoutes(db *db.DB) *chi.Mux {
	scoreHandler := score.ScoreHandler{DB: db}
	router := chi.NewRouter()

	router.Get("/leader-board/", scoreHandler.LeaderBoard)

	router.Group(func(router chi.Router) {
		router.Use(middlewares.AuthMiddleware)

		router.Get("/", scoreHandler.GetScore)
		router.Post("/", scoreHandler.CreateScore)
	})

	return router
}

func OtherRoutes(db *db.DB) *chi.Mux {
	rankHandler := rank.RankHandler{DB: db}
	router := chi.NewRouter()

	router.Use(middlewares.AuthMiddleware)

	router.Get("/ranks/", rankHandler.GetRanks)
	router.Get("/jump-heights/", rankHandler.GetJumpHeights)

	return router
}
