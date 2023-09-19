package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/touch-some-grass-bro/vitamart/handlers"
)

// Function to handle routes
func (s *Server) HandleRoutes(mainRouter *chi.Mux) {

	vitamartRouter := chi.NewRouter()
	vitamartRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Vitamart API!"))
	})

	authRouter := chi.NewRouter()

	authRouter.Get(
		"/url",
		handlers.GetAuthURLHandler(
			s.OauthConf,
		),
	)

	authRouter.Get("/callback",
		handlers.CallbackHandler( s.Queries,
			s.OauthConf,
		),
	)

	authRouter.Get("/logout",
		handlers.LogoutHandler(),
	)

  authRouter.Get("/test", handlers.LogoutHandler())

	vitamartRouter.Mount("/auth", authRouter)

	mainRouter.Mount("/api", vitamartRouter)
}
