package delivery

import (
	"github.com/anggardagasta/some_product/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"time"
)

type handler struct {
	usersUseCase service.IServiceUsersUseCase
}

func Router(usersUsecase service.IServiceUsersUseCase) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Accept-Encoding", "Cookie", "Origin", "X-Api-Key"},
	})
	router.Use(cors.Handler)

	h := handler{usersUseCase: usersUsecase}

	router.Group(func(router chi.Router) {
		router.Use(middleware.SetHeader("Content-Type", "application/json"))

		router.Route("/v1/users", func(router chi.Router) {
			router.Post("/register", h.RegisterUser)
			router.Post("/auth", h.Auth)

			router.Get("/profile", h.GetProfile)
			router.Put("/profile", h.UpdatePicture)
		})
	})
	return router
}
