package main

import (
	"GoLangMiddleware/config"
	"GoLangMiddleware/database"
	user_handler "GoLangMiddleware/handler"
	"GoLangMiddleware/middleware"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"google.golang.org/api/idtoken"
)

var PORT = ":8088"
var (
	cfg                  config.Config
	googleTokenValidator *idtoken.Validator
)

func main() {
	database.Db, database.Err = sql.Open("postgres", ConnectDbPsql(
		database.Host,
		database.User,
		database.Password,
		database.Dbname,
		database.Port))
	if database.Err != nil {
		panic(database.Err)
	}
	defer database.Db.Close()

	database.Err = database.Db.Ping()
	if database.Err != nil {
		panic(database.Err)
	}
	fmt.Println("Successfully Connect to Database")
	_ = cleanenv.ReadConfig(".env", &cfg)

	r := mux.NewRouter()
	userHandler := user_handler.NewUserHandler(database.Db)
	r.HandleFunc("/users", userHandler.UserLoginHandler)
	r.HandleFunc("/users/{id}", userHandler.UserLoginHandler)

	middleware := middleware.NewAuthMiddleware(&cfg, googleTokenValidator)
	r.Use(middleware.GetToken)
	//r.Use(middleware.LoginMiddleware)

	fmt.Println("Now Loading on Port 0.0.0.0" + PORT)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8088",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func nothing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("nothing to show")
}

func ConnectDbPsql(host, user, password, dbname string, port int) string {
	psqlInfo := fmt.Sprintf("host= %s port= %d user= %s "+
		" password= %s dbname= %s sslmode=disable",
		database.Host,
		database.Port,
		database.User,
		database.Password,
		database.Dbname)
	return psqlInfo
}
