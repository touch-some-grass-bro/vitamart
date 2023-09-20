package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	db "github.com/touch-some-grass-bro/vitamart/db/sqlc"
	"github.com/touch-some-grass-bro/vitamart/models"
	"github.com/touch-some-grass-bro/vitamart/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Server struct {
	Router    *chi.Mux
	Queries   *db.Queries
	DB        *sql.DB
	OauthConf *oauth2.Config
	Logger    zerolog.Logger
}

func New() *Server {
	s := &Server{}
	configFromFile, err := utils.LoadConfig("./configs", "config", "yml")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	models.Config = *configFromFile

	s.PrepareLogger()

	if err := s.PrepareDB(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	s.PrepareOauth2()
	s.PrepareRouter()

	return s
}

func (s *Server) PrepareLogger() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}

	logger := zerolog.New(file).With().Timestamp().Logger()
	log.Logger = logger
	s.Logger = logger
}

func (s *Server) PrepareDB() error {
	tries := 5
	DB, err := sql.Open("pgx", models.Config.Database.URI())
	if err != nil {
		return err
	}

	for tries > 0 {
		log.Info().Msg("Attempting to make a connection to the garrix database...")
		err = DB.Ping()
		if err != nil {
			tries -= 1
			log.Info().Msg(err.Error() + "\nCould not connect. Retrying...")
			time.Sleep(8 * time.Second)
			continue
		}
		s.Queries = db.New(DB)
		s.DB = DB
		log.Info().Msg("Connection to the garrix database established.")
		return nil
	}
	return errors.New("Could not make a connection to the database.")
}

func (s *Server) PrepareRouter() {

	r := chi.NewRouter()

	r.Use(httplog.RequestLogger(s.Logger))
	r.Use(middleware.Heartbeat("/ping"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "PUT", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	//Store Router in Struct
	s.Router = r
}

func (s *Server) PrepareOauth2() {
	s.OauthConf = &oauth2.Config{
		ClientID:     models.Config.Google.ClientID,
		ClientSecret: models.Config.Google.ClientSecret,
		RedirectURL:  models.Config.Google.RedirectURI,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

func (s *Server) RunServer() (err error) {

	apiRouter := chi.NewRouter()
	s.HandleRoutes(apiRouter)

	s.Router.Mount("/", apiRouter)
	log.Info().Msg(fmt.Sprintf("Starting Server at %s:%s", models.Config.API.Host, models.Config.API.Port))
	fmt.Println("Server is running.")
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", models.Config.API.Host, models.Config.API.Port), s.Router)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return
}
