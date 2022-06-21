package main

import (
	"GoLangMiddleware/database"
	user_handler "GoLangMiddleware/handler"
	"GoLangMiddleware/middleware"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var PORT = ":8088"

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

	r := mux.NewRouter()
	userHandler := user_handler.NewUserHandler(database.Db)
	r.HandleFunc("/users", userHandler.UserLoginHandler)
	r.HandleFunc("/users/{id}", userHandler.UserLoginHandler)
	r.Use(loginMiddleware)

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

func loginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI, r.Body)
		// userhandler := user_handler.NewUserHandler()
		var check = middleware.Auth(w, r)
		if check {
			middleware.Auth(w, r)
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(401)
			w.Write([]byte("ERROR DATA"))
		}
	})
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
