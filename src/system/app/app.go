package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"learning-golang/src/system/router"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/handlers"
)

type Server struct {
	port     string
	database *xorm.Engine
}

func NewServer() Server {
	return Server{}
}

//init all vals
func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Initializing Server...")
	s.port = ":" + port
	s.database = db
}

func (s *Server) Start() {
	log.Println("Starting server on port" + s.port)

	r := router.NewRouter()

	r.Init(s.database)

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(newServer.ListenAndServe())
}
