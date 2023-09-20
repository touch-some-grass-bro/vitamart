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
		handlers.CallbackHandler(s.Queries,
			s.OauthConf,
		),
	)

	authRouter.Get("/logout",
		handlers.LogoutHandler(),
	)

	authRouter.Get("/is-authenticated",
		handlers.IsAuthenticatedHandler(s.Queries),
	)


  // Chat
  vitamartRouter.Get("/chat", handlers.JoinRoomHandler(s.Queries, s.ChatHub))

	// Items
	vitamartRouter.Get("/items", handlers.GetItemsHandler(s.Queries))
	vitamartRouter.Post("/items", handlers.AddItemHandler(s.Queries))

  // Transactions
  vitamartRouter.Get("/buy", handlers.BuyItemHandler(s.Queries))
  vitamartRouter.Get("/setToSold", handlers.SetProductToSoldHandler(s.Queries))

  // Hostel
	vitamartRouter.Post("/hostel", handlers.SetHostel(s.Queries))

	// Mounting
	vitamartRouter.Mount("/auth", authRouter)

	mainRouter.Mount("/api", vitamartRouter)
}
