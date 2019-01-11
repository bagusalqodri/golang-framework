package main

import (
	"flag"
	"learning-golang/src/system/app"
	DB "learning-golang/src/system/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var port string
var dbhost string
var dbdatabase string
var dbuser string
var dbpass string
var dbport string
var dboptions string

func init() {
	flag.StringVar(&port, "port", "8000", "Assigning the port that the server listen on.")
	flag.Parse()

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	if host := os.Getenv("DB_HOST"); len(host) > 0 {
		dbhost = host
	}

	if database := os.Getenv("DB_DATABASE"); len(database) > 0 {
		dbdatabase = database
	}

	if user := os.Getenv("DB_USER"); len(user) > 0 {
		dbuser = user
	}

	if password := os.Getenv("DB_PASSWORD"); len(password) > 0 {
		dbpass = password
	}

	if port := os.Getenv("DB_PORT"); len(port) > 0 {
		dbport = port
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	db, err := DB.Connect(dbhost, dbport, dbuser, dbpass, dbdatabase, dboptions)

	if db != nil {
		log.Println("Connected Successfully!")
	}

	if err != nil {
		panic(err)
	}

	s := app.NewServer()

	s.Init(port, db)
	s.Start()
}
