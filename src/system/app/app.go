package app

import (
	"log"
	"net/http"

	"learning-golang/src/system/router"

	"github.com/go-xorm/xorm"
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
	log.Println("Starting Server on port " + s.port)

	r := router.NewRouter()

	r.Init()

	http.ListenAndServe(s.port, r.Router)
}
